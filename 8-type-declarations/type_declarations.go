package main

import "fmt"

// tipe deklarasi adalah membuat tipe data baru dari tipe data yang sudah ada
func main() {
	type NoKTP string
	type Married bool

	var ktp_kamu NoKTP = "1234567"
	var marriedStatus Married = true
	fmt.Println(ktp_kamu)
	fmt.Println(marriedStatus)
}
