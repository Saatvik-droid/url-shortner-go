package model

import (
	"fmt"

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
	dsn := "postgres://tqxfasce:nuSeNzOZMz1EKttgADnRiUNJ7w7KFcma@tiny.db.elephantsql.com/tqxfasce"

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = DB.AutoMigrate(&Url{})
	if err != nil {
		fmt.Println(err)
	}
}
