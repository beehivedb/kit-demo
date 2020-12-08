//Package endpoint request and response.
//Author majingru.
//Date 20-11-27 19:17:18
//Copyright (C) 2019-2020 cetccloud.com All RIGHTS RESERVED.
//Version 1.0
package endpoint

import (
	"cetc-gzw/service"
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Request 请求格式
type Request struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

// Response 响应格式
type Response struct {
	Status string `json:"status"`
}

//IndexEndPoint - Send Hello
func IndexEndPoint(e service.Index) endpoint.Endpoint {
	return func(cxt context.Context, request interface{}) (response interface{}, err error) {
		r, ok := request.(Request)
		if !ok {
			return Response{}, nil
		}
		return Response{Status: e.Login(r.User, r.Pass)}, nil
	}
}

//HealthCheckEndPoint - Send Bye
func HealthCheckEndPoint(e service.Index) endpoint.Endpoint {
	return func(cxt context.Context, request interface{}) (response interface{}, err error) {
		return map[string]interface{}{"status": e.HealthCheck()}, nil
	}
}
