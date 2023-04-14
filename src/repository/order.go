package repository

import (
	"errors"
	"log"
	"sellerApp/src/models"
	"sellerApp/utils/database"
)

/******************************************************Api to place an order*************************************************/

func (r *MySqlRepositoryRepo) PlaceOrder(obj interface{}) error {

	db := database.DB.Debug().Create(obj)
	if db.Error != nil {
		log.Printf("Error while saving the data %s", db.Error)
		return db.Error
	}
	return nil

}

/********************************************************Api to cancel order****************************************/

func (r *MySqlRepositoryRepo) CancelOrder(orderId int) error {
	data := &models.Order{}
	db := database.DB.Debug().Table("orders").Find(data)
	if data.Status == "CANCELLED" {
		log.Printf("order already cancelled %s", db.Error)
		return errors.New("Order is already cancelled.")
	}
	db1 := database.DB.Debug().Table("orders").Update("status", "CANCELLED")
	if db1.Error != nil {
		log.Printf("Error while cancelling order %s", db1.Error)
		return db1.Error
	}
	return nil
}

/********************************************************Store item in database****************************************/

func (r *MySqlRepositoryRepo) StoreItem(obj interface{}) error {

	db := database.DB.Debug().Create(obj)
	if db.Error != nil {
		log.Printf("Error while inserting items %s", db.Error)
		return db.Error
	}
	return nil
}
