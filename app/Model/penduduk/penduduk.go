package penduduk

import (
	"main.core/go-validator-request/config/database"
)

var DB = database.ConnectDatabase()

type Penduduk struct {
	Nik             string `gorm:"primaryKey" json:"nik"`
	Nama            string `gorm:"type:varchar(300)" json:"nama"`
	Tgl_lahir       string `gorm:"type:date" json:"tgl_lahir"`
	Waktu_kunjungan string `gorm:"type:timestamp" json:"waktu_kunjungan"`
}

type Pend struct {
	Nik  string `gorm:"primaryKey" json:"nik"`
	Nama string `gorm:"type:varchar(300)" json:"nama"`
}
