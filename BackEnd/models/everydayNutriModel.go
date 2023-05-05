package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
)

type EverydayNutriModel interface {
	GetAll() (*[]entity.EverydayNutrition, error)
	//GetByGenderAndAge(gender, age string) (*entity.RecommendedNutrition, error)
}

type everydayNutriModel struct {
}

func NewEverydayNutriModel() EverydayNutriModel {
	return &everydayNutriModel{}
}

// get all data from recommended_nutrition table
func (e *everydayNutriModel) GetAll() (*[]entity.EverydayNutrition, error) {
	EveNutris := &[]entity.EverydayNutrition{}
	err := loaders.DB.Find(EveNutris).Error
	if err != nil {
		return nil, err
	} else {
		return EveNutris, nil
	}
}
