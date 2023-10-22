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
	d, err:= gorm.Open("mysql","root:Kkannan@mysql1@/ZOOMANIAN?charset=utf8&parseTime=True&loc=Local")
	if err != nil{
		panic(err)
	}
	DB=d
	d.AutoMigrate(&models.User{})
}

func GetDB() *gorm.DB{
	return DB
}