package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kavikkannan/go-jwt/pkg/models"
)

var(
	DB* gorm.DB
)

func Connect(){
	d, err:= gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/auth?parseTime=true")
	if err != nil{
		panic(err)
	}
	DB=d
	d.AutoMigrate(&models.Login{})
}

func GetDB() *gorm.DB{
	return DB
}
