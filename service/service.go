//Package service business service.
//Author majingru.
//Date 20-11-27 19:17:52
//Copyright (C) 2019-2020 cetccloud.com All RIGHTS RESERVED.
//Version 1.0
package service

import (
	"cetc-gzw/config"
	"fmt"
)

//Index simple test interface.
type Index interface {
	Login(user, pass string) string
	HealthCheck() bool
}

//Home implement Index Interface.
type Home struct {
}

//Login say hello to user.
func (e *Home) Login(user, pass string) string {
	u := config.GetString("user")
	p := config.GetString("pass")
	if u == user && p == pass {
		return fmt.Sprintf("%s: welcome", user)
	}
	return fmt.Sprintf("%s: login failed", user)
}

//HealthCheck - check status.
func (e *Home) HealthCheck() bool {
	return true
}
