package db

import (
	"errors"
	"fmt"

	"github.com/jpdhaliwal22/Todo-Golang/model"
	"gorm.io/gorm"
)

// CreateUser implements Database.
func (m *mysqlDb) CreateUser(user model.User) (int64, error) {
	result := m.db.Create(&user)

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil

}

func (m *mysqlDb) GetUser(filters map[string]interface{}) (model.User, error) {
	var user model.User

	fmt.Println("using filters", filters)
	result := m.db.Where("user_name=?", filters["user_name"]).Find(&user)
	if result.Error != nil {
		fmt.Println("NOT FOUND ERROR")
		return model.User{}, result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		fmt.Println("NOT FOUND")
		return model.User{}, errors.New("record not found")
	}

	fmt.Println("res")
	return user, nil
}
