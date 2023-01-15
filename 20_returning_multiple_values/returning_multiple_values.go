package main

import "fmt"

func getFullName() (string, string) {
	return "Muchamad", "Hatta"
}

func main() {
	firstName, lastname := getFullName()
	fmt.Println(firstName, lastname)
}
