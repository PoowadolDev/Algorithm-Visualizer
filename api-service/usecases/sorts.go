package usecases

import (
	"github.com/PoowadolDev/Algorithm-Visualizer/entities"
)

type SortAlgorithm interface {
	SelectionSort(dataList entities.SortData) error
	MergeSort(dataList entities.SortData) error
	QuickSort(dataList entities.SortData) error
}
