package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
	"fmt"
	"time"
)

type RecordModel interface {
	CreateRecord(record *entity.Record) (err error)
	GetRecordsByDate(date time.Time) ([]*entity.Record, error)
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
func (r *recordModel) GetRecordsByDate(date time.Time) ([]*entity.Record, error) {
	var records []*entity.Record
	result := loaders.DB.Debug().Where("date >= ? AND date < ?", date.Format("2006-01-02"), date.Add(24*time.Hour).Format("2006-01-02")).Find(&records)
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}
