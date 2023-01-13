package main

import "fmt"

func getFullName() (string, string) {
	return "Muchamad", "Hattaa"
}

func main() {
	firstName, lastname := getFullName()
	fmt.Println(firstName, lastname)
}
