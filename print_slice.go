package go_print_slice

import "fmt"

func PrintSlice1D(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Println()
}

func PrintSlice1DStr(slice []string) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Println()
}

func PrintSlice2D(slice [][]int) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			if slice[i][j] == 0 {
				fmt.Print("-\t")
			} else {
				fmt.Print(slice[i][j], "\t")
			}
		}
		fmt.Println()
	}
}

func PrintSlice2DStr(slice [][]string) {
	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice[i]); j++ {
			fmt.Print(slice[i][j], "\t")
		}
		fmt.Println()
	}
}
