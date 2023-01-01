package main

import "fmt"

func main() {
	var name = "ata Muchamad hatta"
	var length = len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}
}
