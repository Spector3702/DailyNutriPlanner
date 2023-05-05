package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
)

type FoodstuffModel interface {
	GetByName(name string) (*entity.Foodstuff, error)
	GetBySno(Sno string) (*entity.Foodstuff, error)
}

type foodstuffModel struct {
}

func NewFoodstuffModel() FoodstuffModel {
	return &foodstuffModel{}
}

// get the first tuple based on given gender and age
func (f *foodstuffModel) GetByName(name string) (*entity.Foodstuff, error) {
	foodstuff := &entity.Foodstuff{}
	err := loaders.DB.Where("name = ?", name).First(foodstuff).Error
	if err != nil {
		return nil, err
	}
	return foodstuff, nil
}

func (f *foodstuffModel) GetBySno(Sno string) (*entity.Foodstuff, error) {
	foodstuff := &entity.Foodstuff{}
	err := loaders.DB.Where("Sno = ?", Sno).First(foodstuff).Error
	if err != nil {
		return nil, err
	}
	return foodstuff, nil
}
