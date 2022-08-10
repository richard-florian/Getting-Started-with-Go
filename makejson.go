// Write a program which prompts the user to first enter a name, and then enter an address.
// Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
// Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Name: ")
	scanner.Scan()
	namIn := scanner.Text()

	fmt.Printf("Address: ")
	scanner.Scan()
	addIn := scanner.Text()

	namAdd := make(map[string]string)
	namAdd["Name"] = namIn
	namAdd["Address"] = addIn

	bMap, err := json.Marshal(namAdd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(bMap))
}
