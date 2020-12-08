//Package the entry of system.
//Author majingru.
//Date 20-12-02 10:16:16
//Copyright (C) 2019-2020 cetccloud.com All RIGHTS RESERVED.
//Version 1.0
package main

import (
	"cetc-gzw/endpoint"
	"cetc-gzw/service"
	"cetc-gzw/transport"
	"net/http"
	"os"
	"os/signal"

	httpTransport "github.com/go-kit/kit/transport/http"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)

	s := &service.Home{}

	index := endpoint.IndexEndPoint(s)

	homeServer := httpTransport.NewServer(index, transport.DecodeRequest, transport.EncodeResponse)

	http.Handle("/home", homeServer)

	go func() {
		http.ListenAndServe(":8000", nil)
	}()
	<-c
}
