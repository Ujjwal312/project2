package controller

import (

	"net/http"

	"example.com/m/models"
	"example.com/m/services"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.Orderservice
}

func New(orderService services.Orderservice) OrderController {
	return OrderController{
		orderService: orderService,
	}
}

func (uc *OrderController) Createorder(ctx *gin.Context) {




	var order models.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	res,err1:=http.Get("http://localhost:9090/v1/getproduct/"+order.Product)
	


	if res.StatusCode != http.StatusOK{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err1.Error()})
		return
	}	

	
	
	res,err2:=http.Get("http://localhost:9090/v1/user/get/"+order.UserId)

	if res.StatusCode != http.StatusOK{
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err2.Error()})
		return
	}
	
	err := uc.orderService.CreateOrder(&order)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *OrderController) RegisterorderRoutes(rg *gin.RouterGroup) {
	orderroute := rg.Group("/order")
	orderroute.POST("/createorder", uc.Createorder)
    
}
