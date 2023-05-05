package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
)

type PersnInfoModel interface {
	GetByEmail(email string) (*entity.PersonalInfo, error)
}

type persnInfoModel struct {
}

func NewPersnInfoModel() PersnInfoModel {
	return &persnInfoModel{}
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
