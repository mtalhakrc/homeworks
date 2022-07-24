package main

import "fmt"

func main() {
	sortedArray := []int{1, 3, 5, 7, 9}
	fmt.Println(binarySearch(sortedArray, 2))

}
func binarySearch(array []int, item int) bool {
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == item {
			return true
		} else if array[mid] > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return false
}
