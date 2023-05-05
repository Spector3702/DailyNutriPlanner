package service

import (
	"BackEnd/entity"
	model "BackEnd/models"
	"strings"
)

var RecNutriService = newRecNutriService()

func newRecNutriService() *recNutriService {
	return &recNutriService{m: model.NewRecNutriModel()}
}

type recNutriService struct {
	m model.RecNutriModel
}

func (s *recNutriService) CreateUser(gender, age string) (*entity.RecommendedNutrition, error) {
	gender = strings.TrimSpace(gender)
	age = strings.TrimSpace(age)

	user := &entity.RecommendedNutrition{
		Gender: gender,
		Age:    age,
	}

	err := s.m.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
