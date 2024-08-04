package usecases

import (
	"errors"

	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
	"github.com/PoowadolDev/Algorithm-Visualizer/repository"
)

type AlgorithmUseCase interface {
	SortProblem(dataList entities.SortData) ([]entities.SolveData, error)
	GenerateData(size int) ([]int, error)
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

	if data.SortType == "Insertion" {
		return s.repo.Insertion(data)
	}

	if data.SortType == "Bubble" {
		return s.repo.Bubble(data)
	}

	return []entities.SolveData{}, errors.New("This is Error")
}

func (s *AlgorithmService) GenerateData(size int) ([]int, error) {
	if size <= 0 {
		return []int{}, errors.New("Size must be greater than 0")
	}
	return s.repo.GenerateData(size)
}
