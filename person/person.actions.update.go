package person

import (
	"errors"
	"gorm.io/gorm"
)

func UpdatePerson(person *Person, database *gorm.DB) error {
	if person.Id == nil || *person.Id < 1 {
		return errors.New("invalid id")
	}

	if updated := database.Model(&Person{}).Where(*person.Id).Updates(person); updated.Error != nil {
		return errors.New(updated.Error.Error())
	}

	return nil
}
