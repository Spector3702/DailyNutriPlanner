package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
	"fmt"
)

type RecordModel interface {
	CreateRecord(record *entity.Record) (err error)
}

type recordModel struct {
}

func NewRecordModel() RecordModel {
	return &recordModel{}
}
func (r *recordModel) CreateRecord(record *entity.Record) (err error) {
	result := loaders.DB.Debug().Create(record)
	err = result.Error

	if err != nil {
		return err
	}
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
	}

	return
}
