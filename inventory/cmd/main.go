package main

import (
	conf "inventory/config"
	"inventory/internal/handler"
	"inventory/internal/repository"

	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := conf.Connect()
	if err != nil {
		fmt.Println("Ошибка подключения к базе данных:", err)
		return
	}
	fmt.Println("Успешное подключение к базе данных!")

	productRepo := repository.NewProductRepository(db)
	productHandler := handler.NewProductHandler(productRepo)

	r := gin.Default()

	// Роуты для продуктов
	r.POST("/products", productHandler.CreateProduct)
	r.GET("/products", productHandler.ListProducts)
	r.GET("/products/:id", productHandler.GetProductByID)
	r.PATCH("/products/:id", productHandler.UpdateProduct)
	r.DELETE("/products/:id", productHandler.DeleteProduct)

	r.Run(":8081")
}
