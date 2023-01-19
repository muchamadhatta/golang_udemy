package main

import "fmt"

func getFullName() (string, string) {
	return "Muchamaddd", "Hattaaaa"
}

func main() {
	firstName, lastname := getFullName()
	fmt.Println(firstName, lastname)
}
