// Code generated by thriftrw-plugin-yarpc
// @generated

package barserver

import (
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	fooserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends/fooserver"
)

// Interface is the server-side interface for the Bar service.
type Interface interface {
	fooserver.Interface
}

// New prepares an implementation of the Bar service for
// registration.
//
// 	handler := BarHandler{}
// 	dispatcher.Register(barserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {

	service := thrift.Service{
		Name:    "Bar",
		Methods: []thrift.Method{},
	}

	procedures := make([]transport.Procedure, 0, 0)

	procedures = append(
		procedures,
		fooserver.New(
			impl,
			append(
				opts,
				thrift.Named("Bar"),
			)...,
		)...,
	)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }