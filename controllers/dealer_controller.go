package controllers

import (
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type Dealer struct {
	ID      int
	Nama    string
	Alamat  string
	Telepon string
}

var dummyDealer = []Dealer{
	{ID: 1, Nama: "Dealer A", Alamat: "Jakarta", Telepon: "021-123456"},
	{ID: 2, Nama: "Dealer B", Alamat: "Bandung", Telepon: "022-789012"},
}

type DealerController struct {
	beego.Controller
}

func (c *DealerController) Index() {
	c.Data["Title"] = "List Dealer"
	c.Data["Dealer"] = dummyDealer
	c.TplName = "dealer/index.html"
}

func (c *DealerController) New() {
	c.Data["Title"] = "Buat Dealer Baru"
	c.TplName = "dealer/new.html"
}

func (c *DealerController) Create() {
	nama := c.GetString("nama")
	alamat := c.GetString("alamat")
	telepon := c.GetString("telepon")

	newID := len(dummyDealer) + 1
	dummyDealer = append(dummyDealer, Dealer{
		ID:      newID,
		Nama:    nama,
		Alamat:  alamat,
		Telepon: telepon,
	})

	c.Redirect("/dealer", 302)
}

func (c *DealerController) Edit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	for _, d := range dummyDealer {
		if d.ID == id {
			c.Data["Title"] = "Edit Dealer"
			c.Data["Dealer"] = d
			c.TplName = "dealer/edit.html"
			return
		}
	}

	c.Data["Message"] = "Data tidak ditemukan"
	c.TplName = "errors/404.html"
}

func (c *DealerController) Update() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	nama := c.GetString("nama")
	alamat := c.GetString("alamat")
	telepon := c.GetString("telepon")

	for i, d := range dummyDealer {
		if d.ID == id {
			dummyDealer[i].Nama = nama
			dummyDealer[i].Alamat = alamat
			dummyDealer[i].Telepon = telepon
			break
		}
	}

	c.Redirect("/dealer", 302)
}

func (c *DealerController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	for i, d := range dummyDealer {
		if d.ID == id {
			dummyDealer = append(dummyDealer[:i], dummyDealer[i+1:]...)
			break
		}
	}

	c.Redirect("/dealer", 302)
}
