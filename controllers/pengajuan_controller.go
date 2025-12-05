package controllers

import (
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type Pengajuan struct {
    ID          int
    Nama        string
    Keterangan  string
    Tanggal     string
    Status      string
}

var dummyPengajuan = []Pengajuan{
    {ID: 1, Nama: "Pengajuan Laptop", Keterangan: "Butuh laptop untuk kerja", Tanggal: "2025-01-10", Status: "Pending"},
    {ID: 2, Nama: "Pengajuan Software", Keterangan: "Lisensi VS Code Pro", Tanggal: "2025-01-15", Status: "Approved"},
}

type PengajuanController struct {
    beego.Controller
}

// ==========================
// HALAMAN LIST
// ==========================
func (c *PengajuanController) List() {
    c.Data["Title"] = "List Pengajuan"
    c.Data["Pengajuan"] = dummyPengajuan
    c.TplName = "pengajuan/index.html"
}

// ==========================
// HALAMAN CREATE (FORM)
// ==========================
func (c *PengajuanController) CreateForm() {
    c.Data["Title"] = "Buat Pengajuan"
    c.TplName = "pengajuan/create.html"
}

// ==========================
// PROSES CREATE
// ==========================
func (c *PengajuanController) Create() {
    nama := c.GetString("nama")
    keterangan := c.GetString("keterangan")

    newID := len(dummyPengajuan) + 1
    dummyPengajuan = append(dummyPengajuan, Pengajuan{
        ID:         newID,
        Nama:       nama,
        Keterangan: keterangan,
        Tanggal:    "2025-01-20",
        Status:     "Pending",
    })

    c.Redirect("/pengajuan", 302)
}

// ==========================
// HALAMAN DETAIL
// ==========================
func (c *PengajuanController) Detail() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

    for _, p := range dummyPengajuan {
        if p.ID == id {
            c.Data["Title"] = "Detail Pengajuan"
            c.Data["Pengajuan"] = p
            c.TplName = "pengajuan/detail.html"
            return
        }
    }

    c.Data["Message"] = "Data tidak ditemukan"
    c.TplName = "errors/404.html"
}

// ==========================
// HALAMAN EDIT
// ==========================
func (c *PengajuanController) EditForm() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

    for _, p := range dummyPengajuan {
        if p.ID == id {
            c.Data["Title"] = "Edit Pengajuan"
            c.Data["Pengajuan"] = p
            c.TplName = "pengajuan/edit.html"
            return
        }
    }

    c.Data["Message"] = "Data tidak ditemukan"
    c.TplName = "errors/404.html"
}

// ==========================
// PROSES UPDATE
// ==========================
func (c *PengajuanController) Update() {
    id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
    nama := c.GetString("nama")
    ket := c.GetString("keterangan")

    for i, p := range dummyPengajuan {
        if p.ID == id {
            dummyPengajuan[i].Nama = nama
            dummyPengajuan[i].Keterangan = ket
        }
    }

    c.Redirect("/pengajuan/"+strconv.Itoa(id), 302)
}
