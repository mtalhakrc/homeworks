package main

import "fmt"

type wordstruct struct {
	word      string
	similarty int
}

func main() {
	word := "talha"
	comparestrings := []string{"asdasdads", "alha"}
	var result int
	var index int
	var final = make([]wordstruct, len(comparestrings))
	for m := 0; m < len(comparestrings); m++ {
		fmt.Printf("-----NEW GRİD-----\n")
		compare := comparestrings[m]
		result = 0
		var grid = make([][]int, len(compare)+1)
		for i := 0; i < len(compare)+1; i++ {
			grid[i] = make([]int, len(word)+1)
		}
		for i := 0; i < len(compare)+1; i++ {
			for j := 0; j < len(word)+1; j++ {
				if i == 0 || j == 0 {
					grid[i][j] = 0
				} else if compare[i-1] == word[j-1] {
					grid[i][j] = grid[i-1][j-1] + 1
					if grid[i][j] > result {
						result = grid[i][j]
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
		fmt.Println(result)
		tmp := wordstruct{
			word:      compare,
			similarty: result,
		}
		final[index] = tmp
		index++
	}

	var ordered = make([]wordstruct, len(final))
	for i := 0; i < len(final); i++ {
		var biggest = final[i].similarty
		var biggestindex int
		for j := 0; j < len(final); j++ {
			if final[j].similarty > biggest {
				biggest = final[j].similarty
				biggestindex = j
			}
		}
		ordered[i] = wordstruct{
			word:      final[biggestindex].word,
			similarty: final[biggestindex].similarty,
		}
	}
	fmt.Println(ordered)
}
