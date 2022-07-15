package main

import (
	"fmt"
)

func checkFloat(userFloat float64) bool {
	return userFloat == float64(int(userFloat))
}

func main() {
	userInput := float64(0)
	fmt.Printf("\nEnter a floating point number: ")
	_, err := fmt.Scan(&userInput)
	if err != nil || checkFloat(userInput) == true {
		main()
	} else {
		truncInput := int(userInput)
		fmt.Printf("\n%v truncates to %v\n\n", userInput, truncInput)
	}

}
