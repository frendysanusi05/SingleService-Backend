package databases

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"single-service/models"
)

var DB *gorm.DB

func GetDB() (DB *gorm.DB) {
	return DB
}

/**** MIGRATIONS ****/
func Migrate() {
	// for _, model := range models.GetModel() {
	// 	err := DB.Debug().AutoMigrate(model.Model)

	// 	if err != nil {
	// 		fmt.Println("Migration failed")
	// 		log.Fatalln("error: ", err)
	// 	}
	// }

	DB.AutoMigrate(&models.User{}, &models.Barang{}, &models.Perusahaan{})

	fmt.Println("Database migrated successfully")
}

func ConnectDatabase() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}	
	
	Dbdriver := os.Getenv("DB_DRIVER")
	DbHost := os.Getenv("DB_HOST")
	DbUser := os.Getenv("DB_USER")
	DbPassword := os.Getenv("DB_PASSWORD")
	DbName := os.Getenv("DB_NAME")
	DbPort := os.Getenv("DB_PORT")

	DBURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DbUser, DbPassword, DbHost, DbPort, DbName)
	
	DB, err = gorm.Open(Dbdriver, DBURL)

	if err != nil {
		fmt.Println("Cannot connect to database ", Dbdriver)
		log.Fatal("Connection error:", err)
	} else {
		fmt.Println("Connected to the database ", Dbdriver)
	}

	Migrate()
}