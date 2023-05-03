package service

import (
	"bbs_backend/entity"
	model "bbs_backend/models"
	"strings"
)

var UserService = newUserService()

func newUserService() *userService {
	return &userService{m: model.NewUserModel()}
}

type userService struct {
	m model.UserModel
}

func (s *userService) CreateUser(nickname, email, password string) (*entity.User, error) {
	nickname = strings.TrimSpace(nickname)
	email = strings.TrimSpace(email)

	user := &entity.User{
		Nickname: nickname,
		Email:    email,
		Password: password, // TODO hashing
	}

	err := s.m.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
