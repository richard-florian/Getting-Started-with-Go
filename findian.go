//Write a program which prompts the user to enter a string.
//The program searches through the entered string for the characters ‘i’, ‘a’, and ‘n’.
//The program should print “Found!” if the entered string starts with the character ‘i’,
//ends with the character ‘n’, and contains the character ‘a’. The program should print
//“Not Found!” otherwise. The program should not be case-sensitive, so it does not
//matter if the characters are upper-case or lower-case.

package main

import (
	"fmt"
	"strings"
)

func main() {
	userInput := ""
	fmt.Printf("\nPlease enter a string: ")
	fmt.Scan(&userInput)
	userString := strings.ToLower(userInput)

	//check if string starts with i, contains a, ends with n.
	switch {
	case strings.HasPrefix(userString, "i") && strings.Contains(userString, "a") && strings.HasSuffix(userString, "n"):
		fmt.Printf("Found!\n\n")
	default:
		fmt.Println("Not Found!")
		fmt.Printf("The string should start with i, contain a, and end with n.\n")
		main()

	}

}
