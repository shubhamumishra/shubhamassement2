package controllers

import (
	"net/http"
	"sellerApp/src/models"
	service "sellerApp/src/service"
	"sellerApp/utils/logging"
	"sellerApp/utils/response"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

/****************************************************API to place order*******************************************************/
// @Tags Order service
// @Summary Place order
// @Description Place an order
// @Accept  json
// @Produce  json
// @Params Place Order body models.Order true "Place Order"
// @Success 200 {object} interface{}
// @Failure 404  "Not Found"
// @Failure 501  "Internal Server Error"
// @Router /api/v1/placeOrder [post]
func PlaceOrder(c *gin.Context) {
	reqModel := &models.Order{}
	if err := c.ShouldBind(&reqModel); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	var service = service.OrderService{}
	placeOrder, err := service.PlaceOrder(reqModel)
	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	logging.Logger.Info("Order placed successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(placeOrder))
}

/***********************************************API to cancel order***************************************************/
// @Tags Order service
// @Summary Cancel an order
// @Description Cancel an order
// @Accept  json
// @Produce  json
// @Param orderId query string true "orderId"
// @Success 200 {object} interface{}
// @Failure 404  "Not Found"
// @Failure 501  "Internal Server Error"
// @Router /api/v1/cancelOrder [post]
func CancelOrder(c *gin.Context) {
	orderIdstr := c.Query("orderId")
	orderId, _ := strconv.Atoi(orderIdstr)
	if err := c.ShouldBindQuery(&orderId); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorMessage(http.StatusBadRequest, err))
		return
	}
	var service = service.OrderService{}
	cancel, err := service.CancelOrder(orderId)

	if err != nil {
		logging.Logger.WithField("error", err).WithError(err).Error(http.StatusInternalServerError, err)
		c.JSON(http.StatusInternalServerError, response.ErrorMessage(http.StatusInternalServerError, err))
		return
	}
	logging.Logger.Info("Order cancelled successfully.")
	c.JSON(http.StatusOK, response.SuccessResponse(cancel))
}
