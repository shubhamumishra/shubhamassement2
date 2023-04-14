package models

import (
	"sellerApp/utils/database"

	"gorm.io/gorm"
)

// user table migrate
type Order struct {
	gorm.Model
	OrderId      string  `gorm:"type:varchar(255);unique" json:"id,omitempty" validate:"required"`
	Status       string  `json:"status,omitempty" validate:"required"`
	Items        []Items `json:"items,omitempty" validate:"required"`
	Total        float64 `json:"total,omitempty" `
	CurrencyUnit string  `json:"currencyUnit,omitempty"`
}
type Items struct {
	Id          string  `json:"id,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float64 `json:"price,omitempty"`
	Quantity    int64   `json:"quantity,omitempty"`
}

func OrderMigrate() {
	database.DB.Debug().AutoMigrate(&Order{})
	database.DB.Debug().AutoMigrate(&Items{})
}
