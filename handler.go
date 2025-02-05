package main

import (
	"context"
	api "github.com/lifeng342/go-demo/kitex_gen/api"
)

// HelloImpl implements the last service interface defined in the IDL.
type HelloImpl struct{}

// Echo implements the HelloImpl interface.
func (s *HelloImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	return &api.Response{Message: req.GetMessage()}, nil
}

// Hello implements the HelloImpl interface.
func (s *HelloImpl) Hello(ctx context.Context, req string) (resp string, err error) {
	return req, nil
}
