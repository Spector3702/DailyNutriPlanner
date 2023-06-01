package service

import (
	// "BackEnd/entity"
	model "BackEnd/models"
)

var RecordService = newRecordService()

func newRecordService() *recordService {
	return &recordService{m: model.NewRecordModel()}
}

type recordService struct {
	m model.RecordModel
}
