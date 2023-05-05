package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
)

type RecNutriModel interface {
	GetAll() (*[]entity.RecommendedNutrition, error)
	GetByGenderAndAge(gender, age string) (*entity.RecommendedNutrition, error)
}

type recNutriModel struct {
}

func NewRecNutriModel() RecNutriModel {
	return &recNutriModel{}
}

// get all data from recommended_nutrition table
func (r *recNutriModel) GetAll() (*[]entity.RecommendedNutrition, error) {
	recNutris := &[]entity.RecommendedNutrition{}
	err := loaders.DB.Find(recNutris).Error
	if err != nil {
		return nil, err
	} else {
		return recNutris, nil
	}
}

// get the first tuple based on given gender and age
func (r *recNutriModel) GetByGenderAndAge(gender, age string) (*entity.RecommendedNutrition, error) {
	recommendedNutrition := &entity.RecommendedNutrition{}
	err := loaders.DB.Where("gender = ? AND age = ?", gender, age).First(recommendedNutrition).Error
	if err != nil {
		return nil, err
	}
	return recommendedNutrition, nil
}
