package main

import "fmt"

// pengecekan kondisi panjang hanya ada di golang
func main() {

	var name = "ataaaaaa"
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
