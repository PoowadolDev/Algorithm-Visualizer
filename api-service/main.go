package main

import (
	controller "github.com/PoowadolDev/Algorithm-Visualizer/controller/http"
	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
	"github.com/PoowadolDev/Algorithm-Visualizer/repository"
	"github.com/PoowadolDev/Algorithm-Visualizer/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Set("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS")

		if c.Method() == "OPTIONS" {
			return c.SendStatus(200)
		}

		return c.Next()
	})

	sortRepo := repository.NewAlgorithmRepo(entities.SortData{})
	sortService := usecases.NewSortService(sortRepo)
	sortHandler := controller.NewSortController(sortService)

	app.Post("/sortProblem", sortHandler.SolveSortProblem)
	app.Get("/generateData", sortHandler.GenerateData)

	app.Listen(":4000")

}
