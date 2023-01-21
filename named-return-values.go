package main

import "fmt"

func getCompleteName() (firstName string, middleName string, lastName string) {
	firstName = "muchamad"
	middleName = "hatta"
	lastName = "ata"

	return
}

func main() {
	firstName, middleName, lastName := getCompleteName()
	fmt.Println(firstName, middleName, lastName)
}
