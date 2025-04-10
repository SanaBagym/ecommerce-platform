package main

import (
	conf "order/config"
	"order/internal/handler"
	"order/internal/repository"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := conf.ConnectDB()
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return
	}
	fmt.Println("Успешное подключение к базе данных!")

	orderRepo := repository.NewOrderRepository(db)
	orderHandler := handler.NewOrderHandler(orderRepo)

	r := gin.Default()

	// routes for orders
	r.POST("/orders", orderHandler.CreateOrder)
	r.GET("/orders", orderHandler.ListOrders)
	r.GET("/orders/:id", orderHandler.GetOrderByID)
	r.PATCH("/orders/:id", orderHandler.UpdateOrder)
	r.DELETE("/orders/:id", orderHandler.DeleteOrder)

	r.Run(":8081")
}
