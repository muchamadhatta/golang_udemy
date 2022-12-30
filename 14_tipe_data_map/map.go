package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "ata",
		"address": "pamulang",
	}
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
}
