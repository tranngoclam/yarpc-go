// Code generated by thriftrw-plugin-yarpc
// @generated

package exampleserviceclient

import (
	"context"

	"go.uber.org/yarpc"
	"go.uber.org/yarpc/api/transport"
	"go.uber.org/yarpc/encoding/thrift"
	"go.uber.org/yarpc/transport/x/cherami/example/thrift/example"
)

// Interface is a client for the ExampleService service.
type Interface interface {
	Award(
		ctx context.Context,
		Token *string,
		opts ...yarpc.CallOption,
	) (yarpc.Ack, error)
}

// New builds a new client for the ExampleService service.
//
// 	client := exampleserviceclient.New(dispatcher.ClientConfig("exampleservice"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{c: thrift.New(thrift.Config{
		Service:      "ExampleService",
		ClientConfig: c,
	}, opts...)}
}

func init() {
	yarpc.RegisterClientBuilder(func(c transport.ClientConfig) Interface {
		return New(c)
	})
}

type client struct{ c thrift.Client }

func (c client) Award(
	ctx context.Context,
	_Token *string,
	opts ...yarpc.CallOption,
) (yarpc.Ack, error) {
	args := example.ExampleService_Award_Helper.Args(_Token)
	return c.c.CallOneway(ctx, args, opts...)
}