package main

import "fmt"

// konstan adalah isi tipe data yang nilainya tidak bisa di ubah, berbeda seperti var yang dapat di ubah
func main() {
	const firstName string = "Muchamad"
	const lastName = "Hatta"
	const value = 1000
	// jika konstan tidak di gunakan tidak akan error berbeda dengan var yang akan error ketika tidak digunakan
	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)
	//membuat gabungan konstan
	const (
		tempatLahir string = "Tegal"
		Hobi               = "badminton"
		umur               = 26
	)
	fmt.Println(tempatLahir)
	fmt.Println(Hobi)
	fmt.Println(umur)
}
