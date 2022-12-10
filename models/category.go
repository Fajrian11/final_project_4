package models

import (
	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	Type                string `gorm:"not null" json:"type" form:"type" valid:"required~Your Type is required"`
	Sold_product_amount uint   `json:"sold_product_amount" form:"sold_product_amount"`
	Product             []Product
}

func (c *Category) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(c)
	if errCreate != nil {
		err = errCreate
		return
	}
	err = nil
	return
}

func (c *Category) BeforeUpdate() (err error) {
	govalidator.ValidateStruct(c)

	err = nil
	return
}
