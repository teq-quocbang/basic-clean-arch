package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/quocbang/learn/config"
	"github.com/quocbang/learn/delivery"
)

func main() {
	s := echo.New()

	// load config
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config, error: %v", err)
	}

	// custom logging

	// custom middleware
	s.Use(middleware.Recover())

	// register routes
	if err := delivery.NewDelivery(s, *config); err != nil {
		log.Fatalf("failed to new delivery, error: %v", err)
	}

	// run sever
	if err := s.Start(":8888"); err != nil {
		log.Fatal(err)
	}
}

// delivery -> usecase -> repository
