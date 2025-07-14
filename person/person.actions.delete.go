package person

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func DeletePerson(query string, database *gorm.DB) (*int, error) {
	var id int
	var err error

	if id, err = strconv.Atoi(query); err != nil || id < 1 {
		return nil, errors.New("invalid id")
	}

	deleted := database.Delete(&Person{}, id)
	if deleted.Error != nil {
		return nil, deleted.Error
	}

	return &id, nil
}
