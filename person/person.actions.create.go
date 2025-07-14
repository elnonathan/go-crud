package person

import (
	"errors"
	"gorm.io/gorm"
)

func CreatePerson(person *Person, database *gorm.DB) error {
	if result := database.Create(person); result.Error != nil {
		return errors.New(result.Error.Error())
	}
	return nil
}
