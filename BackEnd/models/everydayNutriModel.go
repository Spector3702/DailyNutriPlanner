package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
)

type EverydayNutriModel interface {
	GetAll() (*[]entity.EverydayNutrition, error)
	GetAllFoodByNeedNutri(nutri, need string) (*[]entity.EverydayNutrition, error)
	GetByName(name string) (*[]entity.EverydayNutrition, error)
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

// get all food satisfied the needed nutri
func (e *everydayNutriModel) GetAllFoodByNeedNutri(nutri, need string) (*[]entity.EverydayNutrition, error) {
	eveNutris := &[]entity.EverydayNutrition{}
	err := loaders.DB.Where("analysis_item = ? AND content_per_unit >= ?", nutri, need).Find(eveNutris).Error
	if err != nil {
		return nil, err
	}
	return eveNutris, nil
}

func (e *everydayNutriModel) GetByName(name string) (*[]entity.EverydayNutrition, error) {
	Nutri := &[]entity.EverydayNutrition{}
	err := loaders.DB.Where("name = ?", name).Find(&Nutri).Error
	if err != nil {
		return nil, err
	}
	return Nutri, nil
}
