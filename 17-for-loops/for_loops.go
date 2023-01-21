package main

import "fmt"

func main() {
	var counter = 1 //kondisi awal

	for counter <= 20 {
		fmt.Println("Perulangan ke", counter) //kondisi perulangan
		counter++                             //kondisi berenti
	}
}
