package main

import (
	_ "jkl-finance/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {

	// Title global (opsional)
	web.BConfig.WebConfig.Session.SessionOn = true

	// Set folder views
	web.SetViewsPath("views")

	// Run server
	web.Run()
}
