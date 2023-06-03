package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
	"fmt"
	"time"
)

type RecordModel interface {
	CreateRecord(record *entity.Record) (err error)
	GetRecordsByDateAndEmail(date time.Time, email string) ([]*entity.Record, error)
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
func (r *recordModel) GetRecordsByDateAndEmail(date time.Time, email string) ([]*entity.Record, error) {
	var records []*entity.Record
	result := loaders.DB.Debug().Where("date >= ? AND date < ? AND email = ?", date.Format("2006-01-02"), date.Add(24*time.Hour).Format("2006-01-02"), email).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}
