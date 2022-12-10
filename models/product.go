package models

import (
	"errors"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title      string `gorm:"not null" json:"title" form:"title" valid:"required~Your title is required"`
	Price      uint   `gorm:"not null" json:"price" form:"price" valid:"required~Your price is required"`
	Stock      uint   `gorm:"not null" json:"stock" form:"stock" valid:"required~Your stock is required"`
	CategoryID uint   `gorm:"not null" json:"category_id" form:"category_id" valid:"required~Your category_id is required"`
}

// validasi field field di database
func (p *Product) BeforeCreate() (err error) {
	_, errCreate := govalidator.ValidateStruct(p)
	if errCreate != nil {
		err = errCreate
		return
	}
	if p.Price > 50000000 {
		err = errors.New("Price Tidak Boleh Melebehi angka 50.000.000")
		return err

	} else if p.Stock < 5 {
		err = errors.New("Stock Tidak Boleh Kurang dari angka 5")
		return err
	}
	err = nil
	return
}

func (p *Product) BeforeUpdate() (err error) {
	govalidator.ValidateStruct(p)
	if p.Price > 50000000 {
		err = errors.New("Price Tidak Boleh Melebehi angka 50.000.000")
		return err

	} else if p.Stock < 5 {
		err = errors.New("Stock Tidak Boleh Kurang dari angka 5")
		return err
	}
	err = nil
	return
}
