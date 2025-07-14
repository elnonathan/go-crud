package person

import (
	"errors"
	"gorm.io/gorm"
	"strconv"
)

func GetUser(query string, database *gorm.DB) (*Person, error) {
	var id int
	var err error
	var found Person

	if id, err = strconv.Atoi(query); err != nil || id < 1 {
		return nil, errors.New("invalid id")
	}

	result := database.First(&found, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &found, nil
}

func ListUser(email string, limitStr string, offsetStr string, database *gorm.DB) (*[]Person, error) {
	limit, limitErr := strconv.Atoi(limitStr)
	if limitErr != nil {
		return nil, limitErr
	}

	offset, offsetErr := strconv.Atoi(offsetStr)
	if offsetErr != nil {
		return nil, offsetErr
	}

	if limit < 1 || offset < 0 {
		return nil, errors.New("limit or offset off limits")
	}

	var list []Person
	query := database.Limit(limit).Offset(offset)

	if email != "" {
		query.Where("email LIKE ?", "%"+email+"%")
	}

	if err := query.Find(&list).Error; err != nil {
		return nil, err
	}

	return &list, nil
}
