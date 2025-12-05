package controllers

import (
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type Pengajuan struct {
	ID         int
	Nama       string
	NIK        string
	Pasangan   string
	Kendaraan  string
	DP         int
	Tenor      int
	Keterangan string
	Tanggal    string
	Status     string
}

var dummyPengajuan = []Pengajuan{
	{
		ID: 1, Nama: "Budi Santoso", NIK: "3201234567890123", Pasangan: "Siti Nurhaliza",
		Kendaraan: "Honda Civic", DP: 50000000, Tenor: 36,
		Keterangan: "Butuh laptop untuk kerja", Tanggal: "2025-01-10", Status: "Pending",
	},
	{
		ID: 2, Nama: "Ahmad Fauzi", NIK: "3201234567890124", Pasangan: "Dewi Sartika",
		Kendaraan: "Toyota Avanza", DP: 30000000, Tenor: 48,
		Keterangan: "Lisensi VS Code Pro", Tanggal: "2025-01-15", Status: "Approved",
	},
}

type PengajuanController struct {
	beego.Controller
}

func (c *PengajuanController) Index() {
	c.Data["Title"] = "List Pengajuan"
	c.Data["Pengajuan"] = dummyPengajuan
	c.TplName = "pengajuan/index.html"
}

func (c *PengajuanController) New() {
	c.Data["Title"] = "Buat Pengajuan"
	c.TplName = "pengajuan/new_pengajuan.html"
}

func (c *PengajuanController) Create() {
	nama := c.GetString("nama")
	nik := c.GetString("nik")
	pasangan := c.GetString("pasangan")
	kendaraan := c.GetString("kendaraan")
	dp, _ := strconv.Atoi(c.GetString("dp"))
	tenor, _ := strconv.Atoi(c.GetString("tenor"))
	keterangan := c.GetString("keterangan")

	newID := len(dummyPengajuan) + 1
	dummyPengajuan = append(dummyPengajuan, Pengajuan{
		ID:         newID,
		Nama:       nama,
		NIK:        nik,
		Pasangan:   pasangan,
		Kendaraan:  kendaraan,
		DP:         dp,
		Tenor:      tenor,
		Keterangan: keterangan,
		Tanggal:    "2025-01-20",
		Status:     "Pending",
	})

	c.Redirect("/pengajuan", 302)
}

func (c *PengajuanController) Edit() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	for _, p := range dummyPengajuan {
		if p.ID == id {
			c.Data["Title"] = "Edit Pengajuan"
			c.Data["Data"] = p
			c.TplName = "pengajuan/edit_pengajuan.html"
			return
		}
	}

	c.Data["Message"] = "Data tidak ditemukan"
	c.TplName = "errors/404.html"
}

func (c *PengajuanController) Update() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	nama := c.GetString("nama")
	nik := c.GetString("nik")
	pasangan := c.GetString("pasangan")
	kendaraan := c.GetString("kendaraan")
	dp, _ := strconv.Atoi(c.GetString("dp"))
	tenor, _ := strconv.Atoi(c.GetString("tenor"))
	keterangan := c.GetString("keterangan")

	for i, p := range dummyPengajuan {
		if p.ID == id {
			dummyPengajuan[i].Nama = nama
			dummyPengajuan[i].NIK = nik
			dummyPengajuan[i].Pasangan = pasangan
			dummyPengajuan[i].Kendaraan = kendaraan
			dummyPengajuan[i].DP = dp
			dummyPengajuan[i].Tenor = tenor
			dummyPengajuan[i].Keterangan = keterangan
			break
		}
	}

	c.Redirect("/pengajuan", 302)
}

func (c *PengajuanController) Delete() {
	id, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	for i, p := range dummyPengajuan {
		if p.ID == id {
			dummyPengajuan = append(dummyPengajuan[:i], dummyPengajuan[i+1:]...)
			break
		}
	}

	c.Redirect("/pengajuan", 302)
}

func UpdatePengajuanStatus(id int, status string) {
	for i := range dummyPengajuan {
		if dummyPengajuan[i].ID == id {
			dummyPengajuan[i].Status = status
			break
		}
	}
}

func GetPengajuanByID(id int) *Pengajuan {
	for i := range dummyPengajuan {
		if dummyPengajuan[i].ID == id {
			return &dummyPengajuan[i]
		}
	}
	return nil
}

func GetPengajuanByStatus(status string) []Pengajuan {
	var result []Pengajuan
	for _, p := range dummyPengajuan {
		if p.Status == status {
			result = append(result, p)
		}
	}
	return result
}

func FormatRupiah(amount int) string {
	return "Rp " + formatNumber(amount)
}

func formatNumber(n int) string {
	str := strconv.Itoa(n)
	var result string
	for i, char := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += "."
		}
		result += string(char)
	}
	return result
}
