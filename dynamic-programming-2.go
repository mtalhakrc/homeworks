package main

import "fmt"

func main() {
	word := "ofetalhadqd"
	compare := "talha"
	//var result int
	//var index int
	var grid = make([][]int, len(compare))
	for i := 0; i < len(compare); i++ {
		grid[i] = make([]int, len(word))
		for j := 0; j < len(word); j++ {
			if compare[i] == word[j] {
				if i == 0 || j == 0 {
					grid[i][j] = grid[i][j] + 1
				} else {
					grid[i][j] = grid[i-1][j-1] + 1
				}
			} else {
				grid[i][j] = 0
			}
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			fmt.Printf("%d  ", grid[i][j])
		}
		fmt.Printf("\n")
	}
}
