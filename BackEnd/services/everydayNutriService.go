package service

import (
	// "BackEnd/entity"
	model "BackEnd/models"
)

var EverydayNutriService = newEverydayNutriService()

func newEverydayNutriService() *everydayNutriService {
	return &everydayNutriService{m: model.NewEverydayNutriModel()}
}

type everydayNutriService struct {
	m model.EverydayNutriModel
}
