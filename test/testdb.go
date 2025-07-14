package test

import (
	"go-crud/person"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func OpenTestDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&person.Person{}); err != nil {
		return nil, err
	}

	return db, nil
}
