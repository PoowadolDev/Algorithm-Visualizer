package repository

import (
	"fmt"

	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
)

type SortAlgorithmRepo interface {
	Selection(dataList entities.SortData) ([]entities.SolveData, error)
	Merge(dataList entities.SortData) ([]entities.SolveData, error)
	Insertion(dataList entities.SortData) ([]entities.SolveData, error)
	Bubble(dataList entities.SortData) ([]entities.SolveData, error)
}

type AlgorithmRepo struct {
	solveData entities.SortData
}

func NewAlgorithmRepo(solveData entities.SortData) SortAlgorithmRepo {
	return &AlgorithmRepo{solveData: solveData}
}

func InitData(startData entities.SortData) ([]int, []entities.SolveData) {
	stepResult := []entities.SolveData{}
	dataArray := startData.DataList
	return dataArray, stepResult
}

func (r *AlgorithmRepo) Selection(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Selection")

	dataArray, stepResult := InitData(data)
	dataArray, stepResult = processSelectionSort(dataArray, stepResult)

	addStepData(&stepResult, dataArray, "Result")

	return stepResult, nil
}

func processSelectionSort(array []int, stepResult []entities.SolveData) ([]int, []entities.SolveData) {
	for i := 0; i < len(array); i++ {
		min := i
		for j := i + 1; j < len(array); j++ {
			if array[min] > array[j] {
				min = j
			}

		}
		temp := array[min]
		fmt.Printf("Switch : %d with %d \n", temp, array[i])
		result := []int{temp, array[i]}
		array[min] = array[i]
		array[i] = temp

		addStepData(&stepResult, result, fmt.Sprint(i+1))
	}

	return array, stepResult
}

func (r *AlgorithmRepo) Merge(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Merge")
	dataArray, stepResult := InitData(data)

	result, stepResult := processMergeSort(dataArray, stepResult)
	addStepData(&stepResult, result, "Result")
	return stepResult, nil
}

func processMergeSort(array []int, stepResult []entities.SolveData) ([]int, []entities.SolveData) {
	length := len(array)
	if length <= 1 {
		return array, stepResult
	}

	mid := length / 2
	left, stepResult := processMergeSort(array[:mid], stepResult)
	right, stepResult := processMergeSort(array[mid:], stepResult)

	// fmt.Printf("Result of Merge Right: %d \n", right)

	return merge(left, right, stepResult)
}

func merge(left, right []int, stepResult []entities.SolveData) ([]int, []entities.SolveData) {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			fmt.Printf("Switch: %d with %d \n", left[i], right[j])
			result := []int{left[i], right[j]}
			addStepData(&stepResult, result, "Left")
			merged = append(merged, left[i])
			i++
		} else {
			fmt.Printf("Switch: %d with %d \n", right[j], left[i])
			result := []int{right[j], left[i]}
			addStepData(&stepResult, result, "Right")
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	return merged, stepResult
}

func (r *AlgorithmRepo) Insertion(data entities.SortData) ([]entities.SolveData, error) {
	dataArray, stepResult := InitData(data)
	dataArray, stepResult = processInsertionSort(dataArray, stepResult)

	addStepData(&stepResult, dataArray, "Result")

	return stepResult, nil
}

func processInsertionSort(array []int, stepResult []entities.SolveData) ([]int, []entities.SolveData) {
	for i, v := range array {
		j := i - 1
		if i >= 1 {
			for j >= 0 && array[j] > v {
				result := []int{array[j+1], array[j]}
				addStepData(&stepResult, result, fmt.Sprint(i+1))
				array[j+1] = array[j]
				j--
			}
			array[j+1] = v
		}
	}

	return array, stepResult
}

func (r *AlgorithmRepo) Bubble(data entities.SortData) ([]entities.SolveData, error) {
	stepResult := []entities.SolveData{}
	dataArray := data.DataList

	dataArray, stepResult = processBubbleSort(dataArray, stepResult)

	addStepData(&stepResult, dataArray, "Result")

	return stepResult, nil
}

func processBubbleSort(array []int, stepResult []entities.SolveData) ([]int, []entities.SolveData) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				result := []int{array[i], array[j]}
				addStepData(&stepResult, result, fmt.Sprint(i+1))
				array[i], array[j] = array[j], array[i]
			}
		}
	}

	return array, stepResult
}

func addStepData(array *[]entities.SolveData, result []int, state string) {
	newStepData := entities.SolveData{
		Step:     state,
		DataList: result,
	}
	*array = append(*array, newStepData)
}
