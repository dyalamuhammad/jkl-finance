package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	web.Controller
}

func (c *AuthController) Login() {
	if c.Ctx.Input.Method() == "GET" {
		c.TplName = "auth/login.html"
		return
	}

	email := c.GetString("email")
	password := c.GetString("password")

	dummyUsers := []map[string]string{
		{"email": "marketing@test.com", "password": "123", "role": "marketing"},
		{"email": "atasan@test.com", "password": "123", "role": "approver"},
		{"email": "backoffice@test.com", "password": "123", "role": "backoffice"},
	}

	var loggedUser map[string]string
	for _, u := range dummyUsers {
		if u["email"] == email && u["password"] == password {
			loggedUser = u
			break
		}
	}

	if loggedUser == nil {
		c.Data["Error"] = "Email atau password salah"
		c.TplName = "auth/login.html"
		return
	}

	if web.BConfig.WebConfig.Session.SessionOn {
		c.SetSession("email", loggedUser["email"])
		c.SetSession("role", loggedUser["role"])
	}

	switch loggedUser["role"] {
	case "marketing":
		c.Redirect("/pengajuan", 302)
	case "approver":
		c.Redirect("/approval", 302)
	case "backoffice":
		c.Redirect("/backoffice", 302)
	default:
		c.Redirect("/", 302)
	}
}
