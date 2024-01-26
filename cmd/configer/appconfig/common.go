package appconfig

import "github.com/thirteenths/test-bmstu23/cmd/configer/appconfig/vars"

func basePath(appName string) string {
	return "/" + vars.Project + "/" + appName + "/"
}
