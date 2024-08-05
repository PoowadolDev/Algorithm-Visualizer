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

	sortRepo := repository.NewAlgorithmRepo(entities.SortData{})
	sortService := usecases.NewSortService(sortRepo)
	sortHandler := controller.NewSortController(sortService)

	app.Get("/sortProblem", sortHandler.SolveSortProblem)
	app.Get("/generateData", sortHandler.GenerateData)

	app.Listen(":4000")

}
