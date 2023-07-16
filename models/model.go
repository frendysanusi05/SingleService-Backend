package models

import (
	"github.com/jinzhu/gorm"
)

/********* USER ***********/
type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Name 	 string `gorm:"size:255;not null;" json:"name"`
	Password string `gorm:"size:255;not null;" json:"password"`
}

/********* BARANG ***********/
type Barang struct {
	gorm.Model
	Nama			string		`gorm:"size:255;not null" json:"nama"`
	Harga			int			`gorm:"size:255;not null" json:"harga"`
	Stok			int			`gorm:"size:255;not null" json:"stok"`
	Kode			string		`gorm:"size:255;not null" json:"kode"`
	IDPerusahaan	Perusahaan	`gorm:"references:id;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"perusahaan_id"`
}

/********* PERUSAHAAN ***********/
type Perusahaan struct {
	gorm.Model
	Nama	string	`gorm:"size:255;not null" json:"nama"`
	Alamat	string	`gorm:"size:255;not null" json:"alamat"`
	NoTelp	string	`gorm:"size:255;not null" json:"no_telp"`
	Kode	string	`gorm:"size:255;not null" json:"kode"`
}

/********* MODEL ***********/
type Model struct {
	Model interface{}
}

/********* VARIABLES ***********/
var u User

/********* GETTER ***********/
// func GetModel() []Model {
// 	return []Model {
// 		{Model: User{}},
// 		{Model: Barang{}},
// 		{Model: Perusahaan{}},
// 	}
// }

func GetUser() (u User) {
	return u
}
