package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// User role: konsumen, dealer, marketing, atasan, backoffice
type User struct {
	Id        int       `orm:"auto;pk" json:"id"`
	Nama      string    `orm:"size(150)" json:"nama"`
	Email     string    `orm:"size(150);null" json:"email"`
	Password  string    `orm:"size(255)" json:"password"` // simpan hash
	Role      string    `orm:"size(50)" json:"role"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
}

type Pengajuan struct {
    Id         int
    Nama       string
    JenisKredit string
    Jumlah     int
    Status     string
}

type Dokumen struct {
	Id         int       `orm:"auto;pk" json:"id"`
	PengajuanId int      `json:"pengajuan_id"`
	Jenis      string    `orm:"size(100)" json:"jenis"` // ktp, spk, bukti_bayar, kontrak, po
	FileUrl    string    `orm:"size(500)" json:"file_url"`
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
}

type Approval struct {
	Id         int       `orm:"auto;pk" json:"id"`
	PengajuanId int      `json:"pengajuan_id"`
	ApproverId int       `json:"approver_id"`
	Status     string    `orm:"size(50)" json:"status"` // approved / rejected
	Catatan    string    `orm:"type(text);null" json:"catatan"`
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)" json:"created_at"`
}

func RegisterModels() {
	orm.RegisterModel(new(User), new(Pengajuan), new(Dokumen), new(Approval))
}
