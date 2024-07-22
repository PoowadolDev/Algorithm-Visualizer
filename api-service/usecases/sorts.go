package usecases

import (
	"errors"

	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
	"github.com/PoowadolDev/Algorithm-Visualizer/repository"
)

type AlgorithmUseCase interface {
	SortProblem(dataList entities.SortData) ([]entities.SolveData, error)
}

type AlgorithmService struct {
	repo repository.SortAlgorithmRepo
}

func NewSortService(sortRepo repository.SortAlgorithmRepo) AlgorithmUseCase {
	return &AlgorithmService{repo: sortRepo}
}

func (s *AlgorithmService) SortProblem(data entities.SortData) ([]entities.SolveData, error) {

	if data.SortType == "Selection" {
		return s.repo.Selection(data)
	}

	if data.SortType == "Merge" {
		return s.repo.Merge(data)
	}

	if data.SortType == "Quick" {
		return s.repo.Quick(data)
	}

	return []entities.SolveData{}, errors.New("This is Error")
}
