package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type (
	Config struct {
		OrderService     OrderService
		InventoryService InventoryService
		UserService      UserService
	}

	OrderService struct {
		Addr string
	}

	InventoryService struct {
		Addr string
	}

	UserService struct {
		Addr string
	}
)

func New() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error: %v", err.Error())
	}

	return &Config{
		OrderService: OrderService{
			Addr: os.Getenv("ORDER_SERVICE"),
		},
		InventoryService: InventoryService{
			Addr: os.Getenv("INVENTORY_SERVICE"),
		},
		UserService: UserService{
			Addr: os.Getenv("USER_SERVICE"),
		},
	}
}
