package go_print_slice

import "fmt"

func PrintSlice1D(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Println()
}

func PrintSlice2D(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			fmt.Print(slice[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println()
}
