package main

import "fmt"

func getFullName() (string, string, int) {
	return "Muchamad", "Hatta", 1
}

func main() {
	firstName, lastName, nomor := getFullName()
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(nomor)
}
