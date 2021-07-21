package repository

import (
	"gorm.io/gorm"
)

type LOData struct {
	gorm.Model
	Name        string
	Description string
	Code        string
	Applicable  string
}

func Setup(db *gorm.DB) {
	err := db.AutoMigrate(&LOData{})
	if err != nil {
		panic(err)
	}
}

func Add(l *LOData, db *gorm.DB) int64 {
	x := db.Create(l)
	return x.RowsAffected
}

func Get(code string, db *gorm.DB, data *LOData) {
	db.First(data, "code = ?", code)

}

func GetAll(db *gorm.DB, data *[]LOData) {

	db.Find(data)

}
