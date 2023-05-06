package service

import (
	"BackEnd/entity"
	model "BackEnd/models"
	"strconv"
)

var PersnInfoService = newPersnInfoService()

func newPersnInfoService() *persnInfoService {
	return &persnInfoService{m: model.NewPersnInfoModel()}
}

type persnInfoService struct {
	m model.PersnInfoModel
}

// create new tuple by given info
func (s *persnInfoService) CreateByInfo(email, gender, age, weightStr, heightStr, workload string) (*entity.PersonalInfo, error) {
	// Convert weight and height strings to float64
	weight, err := strconv.ParseFloat(weightStr, 64)
	if err != nil {
		return nil, err
	}

	height, err := strconv.ParseFloat(heightStr, 64)
	if err != nil {
		return nil, err
	}

	persnInfo := &entity.PersonalInfo{
		Email:    email,
		Gender:   gender,
		Age:      age,
		Weight:   weight,
		Height:   height,
		WorkLoad: workload,
	}

	err = s.m.CreatePersnInfo(persnInfo)
	if err != nil {
		return nil, err
	}

	return persnInfo, nil
}
