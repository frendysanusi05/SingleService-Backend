package databases

import (
	"golang.org/x/crypto/bcrypt"

	"single-service/models"
)

var userSeeder = []models.User {
	models.User {
		Username: "admin",
		Name: "Administrator",
		Password: HashPassword("password"),
	},
}

var barangSeeder = []models.Barang {
	models.Barang {
		ID: "1",
		Nama: "Buku Tulis",
		Harga: 10000,
		Stok: 5,
		Kode: "BARANG1",
		PerusahaanID: "1",
	},
	models.Barang {
		ID: "2",
		Nama: "Buku Gambar",
		Harga: 15000,
		Stok: 1,
		Kode: "BARANG2",
		PerusahaanID: "2",
	},
	models.Barang {
		ID: "3",
		Nama: "Buku Mewarnai",
		Harga: 20000,
		Stok: 6,
		Kode: "BARANG3",
		PerusahaanID: "2",
	},
	models.Barang {
		ID: "4",
		Nama: "Buku Kotak Kecil",
		Harga: 25000,
		Stok: 10,
		Kode: "BARANG4",
		PerusahaanID: "1",
	},
	models.Barang {
		ID: "5",
		Nama: "Buku Kotak Besar",
		Harga: 30000,
		Stok: 3,
		Kode: "BARANG5",
		PerusahaanID: "1",
	},
}

var perusahaanSeeder = []models.Perusahaan {
	models.Perusahaan {
		ID: "1",
		Nama: "Bina Ria",
		Alamat: "Jl. K. F. Tandean",
		NoTelp: "085300001111",
		Kode: "BRI",
	},
	models.Perusahaan {
		ID: "2",
		Nama: "Primadona",
		Alamat: "Jl. K. F. Tandean",
		NoTelp: "083511112222",
		Kode: "PRI",
	},
}

/****** GETTER ********/
func GetUserSeeder() []models.User {
	return userSeeder
}

func GetBarangSeeder() []models.Barang {
	return barangSeeder
}

func GetPerusahaanSeeder() []models.Perusahaan {
	return perusahaanSeeder
}

/******** ADDITIONAL FUNCTION *********/
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hashedPassword)
}
