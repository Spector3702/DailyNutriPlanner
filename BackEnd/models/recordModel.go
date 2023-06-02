package model

import (
	"BackEnd/entity"
	"BackEnd/loaders"
	"fmt"
	"time"
)

type RecordModel interface {
	CreateRecord(record *entity.Record) (err error)
<<<<<<< HEAD
	GetRecordsByDate(date time.Time) ([]*entity.Record, error)
=======
	GetRecordsByDateAndEmail(date time.Time, email string) ([]*entity.Record, error)
>>>>>>> 07a620b99b868d8ac500ddce602cdcdad6fa3c1e
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
<<<<<<< HEAD
func (r *recordModel) GetRecordsByDate(date time.Time) ([]*entity.Record, error) {
	var records []*entity.Record
	result := loaders.DB.Debug().Where("date >= ? AND date < ?", date.Format("2006-01-02"), date.Add(24*time.Hour).Format("2006-01-02")).Find(&records)
=======
func (r *recordModel) GetRecordsByDateAndEmail(date time.Time, email string) ([]*entity.Record, error) {
	var records []*entity.Record
	result := loaders.DB.Debug().Where("date >= ? AND date < ? AND email = ?", date.Format("2006-01-02"), date.Add(24*time.Hour).Format("2006-01-02"), email).Find(&records)
>>>>>>> 07a620b99b868d8ac500ddce602cdcdad6fa3c1e
	if result.Error != nil {
		return nil, result.Error
	}
	return records, nil
}
