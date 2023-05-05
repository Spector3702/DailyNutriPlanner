package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
)

type RecNutriModel interface {
	GetAll() (*[]entity.RecommendedNutrition, error)
}

type recNutriModel struct {
}

func NewRecNutriModel() RecNutriModel {
	return &recNutriModel{}
}

func (r *recNutriModel) Get(uid int64) (*entity.RecommendedNutrition, error) {
	user := &entity.RecommendedNutrition{}
	err := loaders.DB.First(user, "id = ?", uid).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (r *recNutriModel) GetAll() (*[]entity.RecommendedNutrition, error) {
	recNutris := &[]entity.RecommendedNutrition{}
	err := loaders.DB.Find(recNutris).Error
	if err != nil {
		return nil, err
	} else {
		return recNutris, nil
	}
}
