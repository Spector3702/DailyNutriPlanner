package model

import (
	"bbs_backend/entity"
	"bbs_backend/loaders"
	"fmt"
)

type UserModel interface {
	GetAll() (*[]entity.User, error)
	Create(user *entity.User) (err error)
}

type userModel struct {
}

func NewUserModel() UserModel {
	return &userModel{}
}

func (r *userModel) Create(user *entity.User) (err error) {
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

func (r *userModel) Get(uid int64) (*entity.User, error) {
	user := &entity.User{}
	err := loaders.DB.First(user, "id = ?", uid).Error
	if err != nil {
		return nil, err
	} else {
		return user, nil
	}
}

func (r *userModel) GetAll() (*[]entity.User, error) {
	users := &[]entity.User{}
	err := loaders.DB.Find(users).Error
	if err != nil {
		return nil, err
	} else {
		return users, nil
	}
}

func (r *userModel) UpdateUser(user *entity.User) (err error) {
	result := loaders.DB.Save(&user)
	err = result.Error
	if result.RowsAffected != 1 {
		fmt.Println("RowsAffected Number fault")
		err = fmt.Errorf("RowsAffected Number fault")
	}
	return
}

func (r *userModel) Delete(id int64) error {
	res := loaders.DB.Delete(&entity.User{}, "id = ?", id)
	if res.RowsAffected < 1 {
		return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", id)
	}

	return nil
}
