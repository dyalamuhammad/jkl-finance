package controllers

import (
	"strconv"

	"github.com/beego/beego/v2/server/web"
)

type ApprovalController struct {
	web.Controller
}

func (c *ApprovalController) Index() {
	pendingPengajuans := GetPengajuanByStatus("Pending")
	var pengajuans []map[string]interface{}
	for _, p := range pendingPengajuans {
		pengajuans = append(pengajuans, map[string]interface{}{
			"ID":        p.ID,
			"Nama":      p.Nama,
			"Kendaraan": p.Kendaraan,
			"Status":    p.Status,
		})
	}

	c.Data["Title"] = "Approval Pengajuan"
	c.Data["Pengajuans"] = pengajuans
	c.TplName = "approval/approval_list.html"
}

func (c *ApprovalController) Detail() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	p := GetPengajuanByID(id)
	if p == nil {
		c.Data["Message"] = "Data tidak ditemukan"
		c.TplName = "errors/404.html"
		return
	}

	data := map[string]interface{}{
		"ID":        p.ID,
		"Nama":      p.Nama,
		"NIK":       p.NIK,
		"Kendaraan": p.Kendaraan,
		"DP":        FormatRupiah(p.DP),
		"Tenor":     p.Tenor,
		"Status":    p.Status,
	}

	c.Data["Title"] = "Review Pengajuan"
	c.Data["Data"] = data
	c.TplName = "approval/approval_detail.html"
}

func (c *ApprovalController) Approve() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	catatan := c.GetString("catatan")
	UpdatePengajuanStatus(id, "Approved")
	_ = catatan
	c.Redirect("/approval", 302)
}

func (c *ApprovalController) Reject() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	catatan := c.GetString("catatan")
	UpdatePengajuanStatus(id, "Rejected")
	_ = catatan
	c.Redirect("/approval", 302)
}
