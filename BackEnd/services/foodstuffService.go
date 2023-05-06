package service

import (
	// "BackEnd/entity"
	model "BackEnd/models"
)

var FoodstuffService = newFoodstuffService()

func newFoodstuffService() *foodstuffService {
	return &foodstuffService{m: model.NewFoodstuffModel()}
}

type foodstuffService struct {
	m model.FoodstuffModel
}
