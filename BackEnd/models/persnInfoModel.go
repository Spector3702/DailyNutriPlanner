package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
	"fmt"
)

type PersnInfoModel interface {
	CreatePersnInfo(persnInfo *entity.PersonalInfo) (err error)
	GetByEmail(email string) (*entity.PersonalInfo, error)
}

type persnInfoModel struct {
}

func NewPersnInfoModel() PersnInfoModel {
	return &persnInfoModel{}
}

// create persnal_info tuple after registering
func (p *persnInfoModel) CreatePersnInfo(persnInfo *entity.PersonalInfo) (err error) {
	result := loaders.DB.Debug().Create(persnInfo)
	err = result.Error

	if err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
	}

	return
}

// get tuple in personal_info by email while login
func (p *persnInfoModel) GetByEmail(email string) (*entity.PersonalInfo, error) {
	personalInfo := &entity.PersonalInfo{}
	err := loaders.DB.Where("email = ?", email).First(personalInfo).Error
	if err != nil {
		return nil, err
	}
	return personalInfo, nil
}
