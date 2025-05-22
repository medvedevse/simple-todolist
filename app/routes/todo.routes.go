package routes

import (
	"skillsrock-test/app/services"

	"github.com/gofiber/fiber/v2"
)

func TodoRoutes(app fiber.Router) {
	app.Get("/", services.Preview)

	app.Post("/tasks", services.CreateTask)
	app.Get("/tasks", services.GetTasks)
	app.Put("/tasks/:id", services.UpdateTask)
	app.Delete("/tasks/:id", services.DeleteTask)
}
