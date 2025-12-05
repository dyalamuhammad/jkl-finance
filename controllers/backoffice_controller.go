package controllers

import (
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type BackofficeController struct {
	beego.Controller
}

func (c *BackofficeController) Index() {
	approvedPengajuans := GetPengajuanByStatus("Approved")
	var pengajuans []map[string]interface{}
	for _, p := range approvedPengajuans {
		pengajuans = append(pengajuans, map[string]interface{}{
			"ID":        p.ID,
			"Nama":      p.Nama,
			"NIK":       p.NIK,
			"Kendaraan": p.Kendaraan,
			"DP":        FormatRupiah(p.DP),
			"Tenor":     p.Tenor,
			"Tanggal":   p.Tanggal,
			"Status":    p.Status,
		})
	}

	c.Data["Title"] = "Dokumen Siap Dicetak"
	c.Data["Pengajuans"] = pengajuans
	c.TplName = "backoffice/index.html"
}

func (c *BackofficeController) Doc() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	p := GetPengajuanByID(id)
	if p == nil || p.Status != "Approved" {
		c.Data["Message"] = "Data tidak ditemukan atau belum di-approve"
		c.TplName = "errors/404.html"
		return
	}

	c.Data["Title"] = "Dokumen Siap Dicetak"
	c.Data["Data"] = map[string]interface{}{
		"ID":        p.ID,
		"Nama":      p.Nama,
		"NIK":       p.NIK,
		"Kendaraan": p.Kendaraan,
		"DP":        FormatRupiah(p.DP),
		"Tenor":     p.Tenor,
		"Tanggal":   p.Tanggal,
		"Status":    p.Status,
	}
	c.TplName = "backoffice/doc_ready.html"
}
