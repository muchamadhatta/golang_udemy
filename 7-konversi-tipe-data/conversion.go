package main

import "fmt"

func main() {
	//mengkonversi isi variabel nilai integer dari 32byte ke 64byte
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	//mengkonversi dari dari 32byte ke 8byte hasilnya tidak bisa karena mencapai batas mencapai titik bawah
	var nilai8 int8 = int8(nilai32)
	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	//mengkonversi dari string ke byte , hasilnya a
	var name = "ata"
	var e = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}
