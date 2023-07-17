// Code generated by Kitex v0.5.2. DO NOT EDIT.

package greetingservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	greeting "github.com/invictus555/auto_codes/greeting_service_v1/kitex_gen/greeting"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	SayHello(ctx context.Context, req *greeting.Request, callOptions ...callopt.Option) (r *greeting.Response, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kGreetingServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kGreetingServiceClient struct {
	*kClient
}

func (p *kGreetingServiceClient) SayHello(ctx context.Context, req *greeting.Request, callOptions ...callopt.Option) (r *greeting.Response, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SayHello(ctx, req)
}
