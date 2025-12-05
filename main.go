package main

import (
	_ "jkl-finance/routers"

	"github.com/beego/beego/v2/server/web"
)

func main() {
	web.BConfig.WebConfig.Session.SessionOn = true
	web.BConfig.WebConfig.Session.SessionName = "jklfinance_session"
	web.BConfig.WebConfig.Session.SessionGCMaxLifetime = 3600
	web.BConfig.WebConfig.Session.SessionAutoSetCookie = true
	web.BConfig.WebConfig.Session.SessionDomain = ""

	web.SetViewsPath("views")
	web.Run()
}
