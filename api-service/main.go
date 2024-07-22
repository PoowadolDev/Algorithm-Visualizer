package main

import (
	"github.com/PoowadolDev/Algorithm-Visualizer/controller"
	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
	"github.com/PoowadolDev/Algorithm-Visualizer/repository"
	"github.com/PoowadolDev/Algorithm-Visualizer/usecases"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	sortRepo := repository.NewAlgorithmRepo(entities.SortData{})
	sortService := usecases.NewSortService(sortRepo)
	sortHandler := controller.NewHttpHandlerController(sortService)

	app.Get("/sortProblem", sortHandler.SolveSortProblem)

	app.Listen(":3000")

}
