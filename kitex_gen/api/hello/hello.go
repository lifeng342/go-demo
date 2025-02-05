// Code generated by Kitex v0.12.1. DO NOT EDIT.

package hello

import (
	"context"
	"errors"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	api "github.com/lifeng342/go-demo/kitex_gen/api"
)

var errInvalidMessageType = errors.New("invalid message type for service method handler")

var serviceMethods = map[string]kitex.MethodInfo{
	"Echo": kitex.NewMethodInfo(
		echoHandler,
		newHelloEchoArgs,
		newHelloEchoResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
	"Hello": kitex.NewMethodInfo(
		helloHandler,
		newHelloHelloArgs,
		newHelloHelloResult,
		false,
		kitex.WithStreamingMode(kitex.StreamingNone),
	),
}

var (
	helloServiceInfo                = NewServiceInfo()
	helloServiceInfoForClient       = NewServiceInfoForClient()
	helloServiceInfoForStreamClient = NewServiceInfoForStreamClient()
)

// for server
func serviceInfo() *kitex.ServiceInfo {
	return helloServiceInfo
}

// for stream client
func serviceInfoForStreamClient() *kitex.ServiceInfo {
	return helloServiceInfoForStreamClient
}

// for client
func serviceInfoForClient() *kitex.ServiceInfo {
	return helloServiceInfoForClient
}

// NewServiceInfo creates a new ServiceInfo containing all methods
func NewServiceInfo() *kitex.ServiceInfo {
	return newServiceInfo(false, true, true)
}

// NewServiceInfo creates a new ServiceInfo containing non-streaming methods
func NewServiceInfoForClient() *kitex.ServiceInfo {
	return newServiceInfo(false, false, true)
}
func NewServiceInfoForStreamClient() *kitex.ServiceInfo {
	return newServiceInfo(true, true, false)
}

func newServiceInfo(hasStreaming bool, keepStreamingMethods bool, keepNonStreamingMethods bool) *kitex.ServiceInfo {
	serviceName := "Hello"
	handlerType := (*api.Hello)(nil)
	methods := map[string]kitex.MethodInfo{}
	for name, m := range serviceMethods {
		if m.IsStreaming() && !keepStreamingMethods {
			continue
		}
		if !m.IsStreaming() && !keepNonStreamingMethods {
			continue
		}
		methods[name] = m
	}
	extra := map[string]interface{}{
		"PackageName": "api",
	}
	if hasStreaming {
		extra["streaming"] = hasStreaming
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.12.1",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloEchoArgs)
	realResult := result.(*api.HelloEchoResult)
	success, err := handler.(api.Hello).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newHelloEchoArgs() interface{} {
	return api.NewHelloEchoArgs()
}

func newHelloEchoResult() interface{} {
	return api.NewHelloEchoResult()
}

func helloHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*api.HelloHelloArgs)
	realResult := result.(*api.HelloHelloResult)
	success, err := handler.(api.Hello).Hello(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = &success
	return nil
}
func newHelloHelloArgs() interface{} {
	return api.NewHelloHelloArgs()
}

func newHelloHelloResult() interface{} {
	return api.NewHelloHelloResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *api.Request) (r *api.Response, err error) {
	var _args api.HelloEchoArgs
	_args.Req = req
	var _result api.HelloEchoResult
	if err = p.c.Call(ctx, "Echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Hello(ctx context.Context, req string) (r string, err error) {
	var _args api.HelloHelloArgs
	_args.Req = req
	var _result api.HelloHelloResult
	if err = p.c.Call(ctx, "Hello", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
