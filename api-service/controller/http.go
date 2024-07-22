package controller

import (
	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
	"github.com/PoowadolDev/Algorithm-Visualizer/usecases"
	"github.com/gofiber/fiber/v2"
)

type httpHandler struct {
	sortAlgorithm usecases.AlgorithmUseCase
}

func NewHttpHandlerController(useCase usecases.AlgorithmUseCase) *httpHandler {
	return &httpHandler{sortAlgorithm: useCase}
}

func (h *httpHandler) SolveSortProblem(c *fiber.Ctx) error {
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
