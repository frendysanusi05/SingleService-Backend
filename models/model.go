package models

import (
	"time"
)

/********* USER ***********/
type User struct {
	Username	string `gorm:"size:255;not null;unique" json:"username"`
	Name 	 	string `gorm:"size:255;not null;" json:"name"`
	Password 	string `gorm:"size:255;not null;" json:"password"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

/********* BARANG ***********/
type Barang struct {
	ID		 		string		`gorm:"size:255;not null;unique;primary_key" json:"id"`
	Nama			string		`gorm:"size:255;not null" json:"nama"`
	Harga			int			`gorm:"not null" json:"harga"`
	Stok			int			`gorm:"not null" json:"stok"`
	Kode			string		`gorm:"size:255;not null" json:"kode"`
	PerusahaanID	string		`gorm:"not null" json:"perusahaan_id"`
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

/********* PERUSAHAAN ***********/
type Perusahaan struct {
	ID			string		`gorm:"size:255;not null;unique;primary_key" json:"id"`
	Nama		string		`gorm:"size:255;not null" json:"nama"`
	Alamat		string		`gorm:"size:255;not null" json:"alamat"`
	NoTelp		string		`gorm:"size:255;not null" json:"no_telp"`
	Kode		string		`gorm:"size:255;not null" json:"kode"`
	CreatedAt	time.Time
	UpdatedAt	time.Time
}

/********* MODEL ***********/
type Model struct {
	Model interface{}
}

/********* GETTER ***********/
// func GetModel() []Model {
// 	return []Model {
// 		{Model: User{}},
// 		{Model: Barang{}},
// 		{Model: Perusahaan{}},
// 	}
// }
