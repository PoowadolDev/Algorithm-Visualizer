package controller

import (
	"fmt"

	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
	"github.com/PoowadolDev/Algorithm-Visualizer/usecases"
	"github.com/gofiber/fiber/v2"
)

type sort struct {
	sortAlgorithm usecases.AlgorithmUseCase
}

func NewSortController(useCaseAlgorithm usecases.AlgorithmUseCase) *sort {
	return &sort{sortAlgorithm: useCaseAlgorithm}
}

func (h *sort) SolveSortProblem(c *fiber.Ctx) error {
	var request entities.SortData
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	response, err := h.sortAlgorithm.SortProblem(request)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}

func (h *sort) GenerateData(c *fiber.Ctx) error {
	var size = c.QueryInt("size")
	fmt.Println(size)
	if size == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	response, err := h.sortAlgorithm.GenerateData(size)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(response)
}
