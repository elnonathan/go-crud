package db

import (
	"go-crud/person"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("db/test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&person.Person{}); err != nil {
		return nil, err
	}

	return db, nil
}
