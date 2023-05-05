package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
	"fmt"
)

type RecNutriModel interface {
	GetAll() (*[]entity.RecommendedNutrition, error)
	Create(user *entity.RecommendedNutrition) (err error)
}

type recNutriModel struct {
}

func NewRecNutriModel() RecNutriModel {
	return &recNutriModel{}
}

func (r *recNutriModel) Create(user *entity.RecommendedNutrition) (err error) {
	result := loaders.DB.Debug().Create(user)
	err = result.Error
	if result.Error != nil {
		fmt.Println("Create failed")
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
	}
	return
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
	users := &[]entity.RecommendedNutrition{}
	err := loaders.DB.Find(users).Error
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (r *recNutriModel) UpdateUser(user *entity.RecommendedNutrition) (err error) {
	result := loaders.DB.Save(&user)
	err = result.Error
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
		err = fmt.Errorf("RowsAffected Number fault")
	}
	return
}

func (r *recNutriModel) Delete(id int64) error {
	res := loaders.DB.Delete(&entity.RecommendedNutrition{}, "id = ?", id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}
