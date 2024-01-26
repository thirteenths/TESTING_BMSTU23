package main

import (
	"github.com/thirteenths/test-bmstu23/cmd/configer/appconfig"
	"github.com/thirteenths/test-bmstu23/cmd/configer/appconfig/vars"
)

func GetConfig(env vars.Env, appName string) any {
	switch appName {
	case appconfig.APIAppName:
		return appconfig.GetAPIConfig(env)
	}

	return nil
}
