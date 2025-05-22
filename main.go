package main

import (
	"context"
	"fmt"
	"os"

	"skillsrock-test/app/routes"
	"skillsrock-test/config"
	"skillsrock-test/config/database"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.InitConfig()

	database.Connect()
	defer database.DB.Close(context.Background())

	app := fiber.New()
	routes.TodoRoutes(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
