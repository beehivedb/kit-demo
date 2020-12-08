//Package config system.
//Author majingru.
//Date 20-12-04 10:29:37
//Copyright (C) 2019-2020 cetccloud.com All RIGHTS RESERVED.
//Version 1.0
package config

import (
	"flag"

	"github.com/pelletier/go-toml"
)

var (
	file string
	conf *toml.Tree
)

func init() {
	flag.StringVar(&file, "conf", "conf.toml", "config file.")
}

//GetString -
func GetString(key string) string {
	if conf == nil {
		c, err := toml.Load(file)
		if err != nil {
			return ""
		}
		conf = c
	}

	return conf.Get(key).(string)
}
