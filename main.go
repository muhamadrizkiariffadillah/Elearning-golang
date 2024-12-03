package main

import (
	"e-learning/handler"
	"e-learning/users"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=root password=123qweasdzxc dbname=e-learning port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// User domain
	userRepository := users.Repositories(db)
	userService := users.Services(userRepository)
	userHandler := handler.UserHandler(userService)

	app := fiber.New()

	v1 := app.Group("/api/v1")

	v1.Post("/user", userHandler.SignUpUser)
	v1.Post("/user/login", userHandler.LoginUser)

	log.Fatal(app.Listen(":8181"))
}
