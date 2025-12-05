package controllers

import (
	"strconv"

	"jkl-finance/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type ApprovalController struct {
	web.Controller
}

// Approve pengajuan
func (c *ApprovalController) Approve() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	catatan := c.GetString("catatan")
	approverId := 0
	if sess := c.GetSession("user_id"); sess != nil {
		approverId = sess.(int)
	}

	o := orm.NewOrm()
	p := models.Pengajuan{Id: id}
	if err := o.Read(&p); err == nil {
		p.Status = "approved"
		o.Update(&p)

		ap := models.Approval{
			PengajuanId: id,
			ApproverId:  approverId,
			Status:      "approved",
			Catatan:     catatan,
		}
		o.Insert(&ap)
	}

	c.Redirect("/pengajuan/list", 302)
}

// Reject pengajuan
func (c *ApprovalController) Reject() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)

	catatan := c.GetString("catatan")
	approverId := 0
	if sess := c.GetSession("user_id"); sess != nil {
		approverId = sess.(int)
	}

	o := orm.NewOrm()
	p := models.Pengajuan{Id: id}
	if err := o.Read(&p); err == nil {
		p.Status = "rejected"
		o.Update(&p)

		ap := models.Approval{
			PengajuanId: id,
			ApproverId:  approverId,
			Status:      "rejected",
			Catatan:     catatan,
		}
		o.Insert(&ap)
	}

	c.Redirect("/pengajuan/list", 302)
}
