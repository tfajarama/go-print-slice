package go_print_slice

import "fmt"

func PrintSlice1D(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Print(slice[i], " ")
	}
	fmt.Println()
}
