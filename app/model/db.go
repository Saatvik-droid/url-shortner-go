package model

import (
	"fmt"

	"github.com/Saatvik-droid/url-shortner/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Url struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Url      string `json:"url" gorm:"not null;unique"`
	Clicked  uint64 `json:"clicked"`
}

var DB *gorm.DB

func SetupDatabase() {
	var err error

	config, err := utils.LoadConfig()
	if err != nil {
		fmt.Println(err)
	}
	DB, err = gorm.Open(postgres.Open(config.DSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&Url{})
	if err != nil {
		fmt.Println(err)
	}
}
