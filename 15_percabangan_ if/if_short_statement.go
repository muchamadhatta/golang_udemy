package main

import "fmt"

// pengecekan kondisi nama terlalu panjang panjang hanya ada di golang
func main() {

	var name = "ataaaaaa"
	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
