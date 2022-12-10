package models

import (
	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type TransactionHistory struct {
	gorm.Model
	ProductID  int `gorm:"not null" json:"product_id" form:"product_id" valid:"required~Your product_id is required"`
	Product    *Product
	UserID     int
	User       *User
	Quantity   uint `gorm:"not null" json:"quantity" form:"quantity" valid:"required~Your quantity is required"`
	TotalPrice uint `json:"total_price" form:"total_price"`
}

// validasi field field di database
func (t *TransactionHistory) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(t)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

// func (p *Product) BeforeUpdate() (err error) {
// 	_, errUpdate := govalidator.ValidateStruct(u)

// 	if errUpdate != nil {
// 		err = errUpdate
// 		return
// 	}
// 	if u.Role != "admin" && u.Role != "member" {
// 		err = errors.New("Role Hanya boleh diisi dengan admin atau member")
// 		return err
// 	}
// 	u.Password = helpers.HashPass(u.Password)

// 	err = nil
// 	return
// }
