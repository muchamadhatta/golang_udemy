package main

import "fmt"

// di golang nama variable harus menyimpan tipe data yang sama
func main() {
	var name string // nama variable string dengan tipe data string

	name = "Muchamad Hatta"
	fmt.Println(name)

	// jika ingin mengubah tipe data
	name = "muchamad hatta"
	fmt.Println(name)

	// membuat tipe data secara langsung
	var friendName = "boy" //tipe data string
	fmt.Println(friendName)

	var age = 26
	fmt.Println(age)

	// di golang kata var tidaklah wajib, tapi bisa di ganti :=
	country := "indonesia"
	fmt.Println(country)
	// mengganti isi variable country
	country = "malaysia"
	fmt.Println(country)

	// menggabungkan dua variable di golang
	var (
		firstName = "muchamad"
		lastName  = "hatta"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	//menggabungkan hasil dua variable
	fmt.Println(firstName, lastName)

}
