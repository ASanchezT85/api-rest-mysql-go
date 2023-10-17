package commons

import (
	"log"

	"github.com/ASanchezT85/api-rest-mysql-go/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/qa_cashship?charset=utf8")
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Println("Migration")
	db.AutoMigrate(&models.User{})
}
