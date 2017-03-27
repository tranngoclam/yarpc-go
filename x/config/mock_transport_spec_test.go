// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package config

import (
	"reflect"

	"go.uber.org/yarpc/api/transport"

	"github.com/golang/mock/gomock"
)

// Builds mockTransportSpec objects.
//
// 	mockSpec := mockTransportSpecBuilder{
// 		Name: "...",
// 		TransportConfig: reflect.TypeOf(myConfig{}),
// 	}.Build()
//
// 	mockSpec.EXPECT().BuildTransport(myConfig{...}).Return(...)
// 	mockSpec.Spec()
type mockTransportSpecBuilder struct {
	Name            string
	TransportConfig reflect.Type

	// Any of the following may be nil to indicate that the transport does not
	// support that functionality.

	InboundConfig        reflect.Type
	UnaryOutboundConfig  reflect.Type
	OnewayOutboundConfig reflect.Type
}

// build the mockTransportSpec.
func (b mockTransportSpecBuilder) Build(ctrl *gomock.Controller) *mockTransportSpec {
	s := TransportSpec{Name: b.Name}
	m := mockTransportSpec{ctrl: ctrl, spec: &s}

	// Build a Spec where the Build* functions point to a dummy function
	// (generated by builderFunc) which calls into the mock controller to get
	// the value to return.

	s.BuildTransport = m.builderFunc("BuildTransport", []reflect.Type{b.TransportConfig}, _typeOfTransport)

	if b.InboundConfig != nil {
		s.BuildInbound = m.builderFunc("BuildInbound", []reflect.Type{b.InboundConfig, _typeOfTransport}, _typeOfInbound)
	}
	if b.UnaryOutboundConfig != nil {
		s.BuildUnaryOutbound = m.builderFunc("BuildUnaryOutbound", []reflect.Type{b.UnaryOutboundConfig, _typeOfTransport}, _typeOfUnaryOutbound)
	}
	if b.OnewayOutboundConfig != nil {
		s.BuildOnewayOutbound = m.builderFunc("BuildOnewayOutbound", []reflect.Type{b.OnewayOutboundConfig, _typeOfTransport}, _typeOfOnewayOutbound)
	}

	return &m
}

// mockTransportSpec sets up a fake TransportSpec. The underlying Spec can be
// obtained using the Spec function.
type mockTransportSpec struct {
	spec *TransportSpec
	ctrl *gomock.Controller
}

// builderFunc builds a `build` function that calls into the controller.
//
// The generated function has the signature,
//
// 	func $name($args[0], $args[1], ..., $args[N]) ($output, error)
//
// This function may be fed as an argument to BuildTransport, BuildInbound,
// etc. and it will be interpreted correctly.
func (m *mockTransportSpec) builderFunc(name string, argTypes []reflect.Type, output reflect.Type) interface{} {
	// We dynamically generate a function with the correct arg type rather
	// than interface{} because we want to verify we're getting the correct
	// type decoded.

	resultTypes := []reflect.Type{output, _typeOfError}
	return reflect.MakeFunc(
		reflect.FuncOf(argTypes, resultTypes, false),
		func(callArgs []reflect.Value) []reflect.Value {
			args := make([]interface{}, len(callArgs))
			for i, a := range callArgs {
				args[i] = a.Interface()
			}

			results := m.ctrl.Call(m, name, args...)
			callResults := make([]reflect.Value, len(results))
			for i, r := range results {
				// Use zero-value where the result is nil because
				// reflect.ValueOf(nil) is an error.
				if r == nil {
					callResults[i] = reflect.Zero(resultTypes[i])
					continue
				}

				callResults[i] = reflect.ValueOf(r).Convert(resultTypes[i])
			}

			return callResults
		},
	).Interface()
}

// The following Build* functions are never called directly. They're just
// there to let gomock verify the signatures and return types.

func (m *mockTransportSpec) BuildTransport(interface{}) (transport.Transport, error) {
	panic("This function should never be called")
}

func (m *mockTransportSpec) BuildInbound(interface{}, transport.Transport) (transport.Inbound, error) {
	panic("This function should never be called")
}

func (m *mockTransportSpec) BuildUnaryOutbound(interface{}, transport.Transport) (transport.UnaryOutbound, error) {
	panic("This function should never be called")
}

func (m *mockTransportSpec) BuildOnewayOutbound(interface{}, transport.Transport) (transport.OnewayOutbound, error) {
	panic("This function should never be called")
}

// EXPECT may be used to define expectations on the TransportSpec.
func (m *mockTransportSpec) EXPECT() *_transportSpecRecorder {
	return &_transportSpecRecorder{m: m, ctrl: m.ctrl}
}

// Spec returns a TransportSpec based on the expectations set on this
// mockTransportSpec.
func (m *mockTransportSpec) Spec() TransportSpec {
	return *m.spec
}

// Provides functions to record TransportSpec expectations.
type _transportSpecRecorder struct {
	m    *mockTransportSpec
	ctrl *gomock.Controller
}

func (r *_transportSpecRecorder) BuildTransport(cfg interface{}) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildTransport", cfg)
}

func (r *_transportSpecRecorder) BuildInbound(cfg interface{}, t transport.Transport) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildInbound", cfg, t)
}

func (r *_transportSpecRecorder) BuildUnaryOutbound(cfg interface{}, t transport.Transport) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildUnaryOutbound", cfg, t)
}

func (r *_transportSpecRecorder) BuildOnewayOutbound(cfg interface{}, t transport.Transport) *gomock.Call {
	return r.ctrl.RecordCall(r.m, "BuildOnewayOutbound", cfg, t)
}