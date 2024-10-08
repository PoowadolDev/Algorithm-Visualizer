package repository

import (
	"fmt"
	"math/rand"

	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
)

type SortAlgorithmRepo interface {
	Selection(dataList entities.SortData) ([]entities.SolveData, error)
	Merge(dataList entities.SortData) ([]entities.SolveData, error)
	Insertion(dataList entities.SortData) ([]entities.SolveData, error)
	Bubble(dataList entities.SortData) ([]entities.SolveData, error)
	GenerateData(size int) ([]int, error)
}

type AlgorithmRepo struct {
	solveData entities.SortData
}

func NewAlgorithmRepo(solveData entities.SortData) SortAlgorithmRepo {
	return &AlgorithmRepo{solveData: solveData}
}

func (r *AlgorithmRepo) GenerateData(size int) ([]int, error) {
	// Create an array of the specified size
	randomArray := make([]int, size)

	// Populate the array with random integers
	for i := 0; i < size; i++ {
		randomArray[i] = rand.Intn(100) // Random integer between 0 and 99
	}

	return randomArray, nil
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

	// addStepData(&stepResult, dataArray, dataArray, "Result")

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
		dataSwitch := []int{temp, array[i]}
		array[min] = array[i]
		array[i] = temp

		currentArray := make([]int, len(array))
		copy(currentArray, array)
		addStepData(&stepResult, dataSwitch, currentArray, fmt.Sprint(i+1))
	}

	return array, stepResult
}

func (r *AlgorithmRepo) Merge(data entities.SortData) ([]entities.SolveData, error) {
	fmt.Println("Merge")
	dataArray, stepResult := InitData(data)

	result, stepResult := processMergeSort(dataArray, stepResult, dataArray)
	fmt.Println("Result of Merge: ", result)
	// addStepData(&stepResult, result, result, "Result")
	return stepResult, nil
}

func processMergeSort(array []int, stepResult []entities.SolveData, fullArray []int) ([]int, []entities.SolveData) {
	length := len(array)
	if length <= 1 {
		return array, stepResult
	}

	mid := length / 2
	left, stepResult := processMergeSort(array[:mid], stepResult, fullArray)
	right, stepResult := processMergeSort(array[mid:], stepResult, fullArray)

	// fmt.Printf("Result of Merge Right: %d \n", right)

	return merge(left, right, stepResult, fullArray)
}

func merge(left, right []int, stepResult []entities.SolveData, fullArray []int) ([]int, []entities.SolveData) {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			fmt.Printf("Switch: %d with %d \n", left[i], right[j])
			merged = append(merged, left[i])
			i++
		} else {
			fmt.Printf("Switch: %d with %d \n", right[j], left[i])
			merged = append(merged, right[j])
			j++
		}
	}

	merged = append(merged, left[i:]...)
	merged = append(merged, right[j:]...)

	fullMergedArray := make([]int, len(fullArray))
	copy(fullMergedArray, fullArray)

	position := len(fullArray) - len(merged)
	for k := 0; k < len(merged); k++ {
		fullMergedArray[position+k] = merged[k]
	}

	addStepData(&stepResult, merged, fullMergedArray, "Merge")

	return merged, stepResult
}

func (r *AlgorithmRepo) Insertion(data entities.SortData) ([]entities.SolveData, error) {
	dataArray, stepResult := InitData(data)
	dataArray, stepResult = processInsertionSort(dataArray, stepResult)

	addStepData(&stepResult, dataArray, dataArray, "Result")

	return stepResult, nil
}

func processInsertionSort(array []int, stepResult []entities.SolveData) ([]int, []entities.SolveData) {
	for i, v := range array {
		j := i - 1
		if i >= 1 {
			for j >= 0 && array[j] > v {
				result := []int{array[j+1], array[j]}
				array[j+1] = array[j]
				currentArray := make([]int, len(array))
				copy(currentArray, array)
				addStepData(&stepResult, result, currentArray, fmt.Sprint(i+1))
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

	addStepData(&stepResult, dataArray, dataArray, "Result")

	return stepResult, nil
}

func processBubbleSort(array []int, stepResult []entities.SolveData) ([]int, []entities.SolveData) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i] > array[j] {
				result := []int{array[i], array[j]}
				array[i], array[j] = array[j], array[i]
				currentArray := make([]int, len(array))
				copy(currentArray, array)
				addStepData(&stepResult, result, currentArray, fmt.Sprint(i+1))
			}
		}
	}

	return array, stepResult
}

func addStepData(array *[]entities.SolveData, dataSwitch []int, result []int, state string) {
	newStepData := entities.SolveData{
		Step:       state,
		DataSwitch: dataSwitch,
		DataList:   result,
	}
	*array = append(*array, newStepData)
}
