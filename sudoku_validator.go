package main

import "fmt"

func main() {
	var testTrue = [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	fmt.Println(validateSolution(testTrue))
}

func validateSolution(m [][]int) bool {
	gridSide := 9

	//Check lines
	for _, line := range m {
		var digitsLineCount = make(map[int]uint, len(line))
		for _, digit := range line {
			digitsLineCount[digit]++
		}
		if checkMap(digitsLineCount) == false {
			return false
		}
	}

	//Check columns
	for i := 0; i < gridSide; i++ {
		var digitsColumnsCount = make(map[int]uint, 9)
		for j := 0; j < gridSide; j++ {
			digitsColumnsCount[m[j][i]]++
		}
		if checkMap(digitsColumnsCount) == false {
			return false
		}
	}

	//Check squares
	squares := 3
	for j := 0; j < squares; j++ {
		for i := 0; i < squares; i++ {
			digitsSquareCount := make(map[int]uint, gridSide)
			for y := 0; y < squares; y++ {
				for x := 0; x < squares; x++ {
					digitsSquareCount[m[x+3*i][y+3*j]]++
				}
			}
			if checkMap(digitsSquareCount) == false {
				return false
			}
		}
	}

	return true
}

func checkMap(digits map[int]uint) bool {
	for _, counts := range digits {
		if counts > 1 {
			return false
		}
	}
	return true
}
