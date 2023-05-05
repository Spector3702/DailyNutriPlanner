package service

import (
	// "BackEnd/entity"
	model "BackEnd/models"
)

var RecNutriService = newRecNutriService()

func newRecNutriService() *recNutriService {
	return &recNutriService{m: model.NewRecNutriModel()}
}

type recNutriService struct {
	m model.RecNutriModel
}
