package repository

import (
	"fmt"

	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
)

type SortAlgorithmRepo interface {
	Selection(dataList entities.SortData) ([]entities.SolveData, error)
	Merge(dataList entities.SortData) ([]entities.SolveData, error)
	Quick(dataList entities.SortData) ([]entities.SolveData, error)
	Bubble(dataList entities.SortData) ([]entities.SolveData, error)
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
		fmt.Printf("Switch : %d with %d \n", temp, dataArray[i])
		result := []int{temp, dataArray[i]}
		dataArray[min] = dataArray[i]
		dataArray[i] = temp

		resultSolve = addStepData(resultSolve, result, fmt.Sprint(i+1))
	}

	resultSolve = addStepData(resultSolve, dataArray, "Result")

	return resultSolve, nil
}

func (r *AlgorithmRepo) Merge(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Merge")
	stepResult := []entities.SolveData{}

	result := mergeSort(data.DataList, stepResult)
	stepResult = addStepData(stepResult, result, "Result")
	return stepResult, nil
}

func mergeSort(array []int, stepResult []entities.SolveData) []int {
	length := len(array)
	if length <= 1 {
		return array
	}

	mid := length / 2
	left := mergeSort(array[:mid], stepResult)
	right := mergeSort(array[mid:], stepResult)

	// fmt.Printf("Result of Merge Right: %d \n", right)

	return merge(left, right, stepResult)
}

func merge(left, right []int, stepResult []entities.SolveData) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			fmt.Printf("Switch: %d with %d \n", left[i], right[j])
			result := []int{left[i], right[j]}
			stepResult = addStepData(stepResult, result, "Left")
			merged = append(merged, left[i])
			i++
		} else {
			fmt.Printf("Switch: %d with %d \n", right[j], left[i])
			result := []int{right[j], left[i]}
			stepResult = addStepData(stepResult, result, "Right")
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged
}

func (r *AlgorithmRepo) Quick(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Quick")
	return []entities.SolveData{}, nil
}

func (r *AlgorithmRepo) Bubble(data entities.SortData) ([]entities.SolveData, error) {
	resultSolve := []entities.SolveData{}
	dataArray := data.DataList
	for i := 0; i < len(dataArray)-1; i++ {
		for j := i + 1; j < len(dataArray); j++ {
			if dataArray[i] > dataArray[j] {
				result := []int{dataArray[i], dataArray[j]}
				resultSolve = addStepData(resultSolve, result, fmt.Sprint(i+1))
				dataArray[i], dataArray[j] = dataArray[j], dataArray[i]
			}
		}

		fmt.Printf("Step: %d Result: %d \n", i+1, dataArray)
	}

	resultSolve = addStepData(resultSolve, dataArray, "Result")

	return resultSolve, nil
}

func addStepData(array []entities.SolveData, result []int, state string) []entities.SolveData {
	array = append(array, entities.SolveData{
		Step:     state,
		DataList: result,
	})
	return array
}
