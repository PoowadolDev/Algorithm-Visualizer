package repository

import (
	"fmt"

	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
)

type SortAlgorithmRepo interface {
	Selection(dataList entities.SortData) ([]entities.SolveData, error)
	Merge(dataList entities.SortData) ([]entities.SolveData, error)
	Quick(dataList entities.SortData) ([]entities.SolveData, error)
}

type AlgorithmRepo struct {
	solveData entities.SortData
}

func NewAlgorithmRepo(solveData entities.SortData) SortAlgorithmRepo {
	return &AlgorithmRepo{solveData: solveData}
}

func (r *AlgorithmRepo) Selection(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Selection")
	resultSolve := []entities.SolveData{}
	dataArray := data.DataList
	for i := 0; i < len(dataArray); i++ {
		min := i
		for j := i + 1; j < len(dataArray); j++ {
			if dataArray[min] > dataArray[j] {
				min = j
			}

		}
		temp := dataArray[min]
		dataArray[min] = dataArray[i]
		dataArray[i] = temp

		result := make([]int, len(dataArray))
		copy(result, dataArray)
		resultSolve = append(resultSolve, entities.SolveData{
			Step:     i + 1,
			DataList: result,
		})
	}

	return resultSolve, nil
}

func (r *AlgorithmRepo) Merge(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Merge")
	return []entities.SolveData{}, nil
}

func (r *AlgorithmRepo) Quick(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Quick")
	return []entities.SolveData{}, nil
}
