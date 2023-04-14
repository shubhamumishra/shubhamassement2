package service

import (
	"log"
	"sellerApp/src/models"
	"sellerApp/src/repository"
	"sellerApp/utils/logging"
)

type OrderService struct {
}

/********************************************************API to place order********************************************/

func (c *OrderService) PlaceOrder(req *models.Order) (*models.Order, error) {
	err := repository.Repo.PlaceOrder(req)
	if err != nil {
		logging.Logger.WithError(err).WithField("error", err).Error("Error while placing the order", err)
		return nil, err
	}
	for i := range req.Items {
		err := repository.Repo.StoreItem(&req.Items[i])
		if err != nil {
			log.Print("Error while inserting item")
			return nil, err
		}
	}
	return req, nil
}

/***************************************************API to cancel order**********************************************/

func (c *OrderService) CancelOrder(orderId int) (string, error) {

	err := repository.Repo.CancelOrder(orderId)
	if err != nil {
		logging.Logger.Error("Error while cancelling the order", err)
		return "nil", err
	}
	str := "Order Cancelled successfully."
	return str, nil
}
