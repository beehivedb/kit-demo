//Package transport http.
//Author majingru.
//Date 20-11-27 19:18:16
//Copyright (C) 2019-2020 cetccloud.com All RIGHTS RESERVED.
//Version 1.0
package transport

import (
	"cetc-gzw/endpoint"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

//DecodeRequest decode request.
func DecodeRequest(cxt context.Context, request *http.Request) (interface{}, error) {
	user := request.URL.Query().Get("user")
	pass := request.URL.Query().Get("pass")

	if user == "" || pass == "" {
		return nil, errors.New("invalid param")
	}

	return endpoint.Request{User: user, Pass: pass}, nil
}

//EncodeResponse encode response
func EncodeResponse(cxt context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
