package main

import (
	"fmt"
)

func main() {
	unsortedArray := []int{25, 10, 12, 11, 51, 4}
	sortedArray := selectionSort(unsortedArray)
	fmt.Println("sorted :", sortedArray)
}
func selectionSort(array []int) []int {
	sortedArray := make([]int, len(array))
	length := len(array)
	for i := 0; i < length; i++ {
		smallest, index := findSmallest(array)
		sortedArray[i] = smallest
		//smallest itemi çıkar.
		if index == 0 {
			//başta
			array = array[1:]
		} else if index == len(array)-1 {
			//sonda
			array = array[:index]
		} else {
			//ortada
			array = append(array[:index], array[index+1:]...)
		}
	}
	return sortedArray
}

func findSmallest(array []int) (int, int) {
	i := 0
	selectedItem := array[i]
	smallestIndex := 0
	for ; i < len(array); i++ {
		if array[i] < selectedItem {
			selectedItem = array[i]
			smallestIndex = i
		}
	}
	return selectedItem, smallestIndex
}
