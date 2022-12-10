package models

import (
	"errors"
	"final_project_4/helpers"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Full_name string `gorm:"not null" json:"full_name" form:"full_name" valid:"required~Your full_name is required"`
	Email     string `gorm:"not null" json:"email" form:"email" valid:"required~Your email is required,email~Invalid email format"`
	Password  string `gorm:"not null" json:"password" form:"password" valid:"required~Your password is required,minstringlength(6)~Password minimal harus 6 karakter"`
	Role      string `gorm:"not null" json:"role" form:"role" valid:"required~Your role is required"`
	Balance   uint   `json:"balance" form:"balance"`
}

type LoginInput struct {
	Email    string `json:"email" form:"email" binding:"required"`
	Password string `json:"password" form:"password" binding:"required"`
}

// validasi field field di database
func (u *User) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if errCreate != nil {
		err = errCreate
		return
	}
	if u.Role != "admin" && u.Role != "customer" {
		err = errors.New("Role Hanya boleh diisi dengan admin atau customer")
		return err
	}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}

func (u *User) BeforeUpdate() (err error) {
	govalidator.ValidateStruct(u)

	if u.Balance < 0 {
		err = errors.New("Balance Tidak Boleh Kurang dari angka 0")
		return err
	} else if u.Balance > 100000000 {
		err = errors.New("Balance Tidak Boleh Melebehi angka 100.000.000")
		return err
	}
	u.Password = helpers.HashPass(u.Password)

	err = nil
	return
}
