package service

import (
	"fmt"

	"github.com/jpdhaliwal22/Todo-Golang/model"
)

// CreateUser implements Service.
func (s *svc) CreateUser(user model.User) (int, error) {
	id, err := s.db.CreateUser(user)
	if err != nil {
		return 0, err

	}

	fmt.Println("newlycreated user id", id)

	return int(id), nil
}

func (s *svc) GetUser(filter map[string]interface{}) (model.User, error) {
	user, err := s.db.GetUser(filter)
	if err != nil {
		return model.User{}, err

	}

	return user, nil

}
