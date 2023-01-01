package main

import "fmt"

//fungsi sederhana dari if

func main() {
	var name = "ata"
	switch name {
	case "ata":
		fmt.Println("Hello ata")
		//didalam percabangan bisa menambahkan lebih dari satu isi
		fmt.Println("Hello ata")
	case "blu":
		fmt.Println("Hello blu")
	default:
		fmt.Println("Hi, Boleh Kenalan?")
	}
}
