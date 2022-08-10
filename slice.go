// Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
// The program should be written as a loop. Before entering the loop, the program should create an empty integer slice of size (length) 3.
// During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
// The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
// The slice must grow in size to accommodate any number of integers which the user decides to enter.
// The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	a_slice := make([]int, 0, 3)
	input := ""

	for {
		fmt.Printf("\nAdd an int to the slice (type X to exit)\n")
		fmt.Scan(&input)

		switch {
		case input == "X":
			return
		default:
			new_int, err := strconv.ParseInt(input, 0, 64)
			if err != nil {
				fmt.Printf("try again\n")
			} else {
				a_slice = append(a_slice, int(new_int))
				for i := 0; i < len(a_slice); i++ {
					for j := i; j > 0 && a_slice[j-1] > a_slice[j]; j-- {
						a_slice[j], a_slice[j-1] = a_slice[j-1], a_slice[j]
					}
				}
				fmt.Println(a_slice)
			}
		}
	}

}
