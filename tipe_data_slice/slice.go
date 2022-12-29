package main

import "fmt"

// slice untuk membagi isi data array
// kegunaan slice sangat membantu apabila data yang cukup banyak
func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}
	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //fungsi len pada slice
	fmt.Println(cap(slice1)) //fungsi cap pada slice

	// months[5] = "diubah"
	// fmt.Println(slice1)

	// slice1[0] = "mei di ubah"
	// fmt.Println(months)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "ata")
	fmt.Println(slice3)

	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "ataa"
	newSlice[1] = "muchamad"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//hati-hati dalam membuat array dan slice bedanya jika array ada nilainya [5] atau [...]
	//kalau slice [] kosong
	var iniArray = [5]int{1, 2, 3, 4, 5}
	var iniSlice = []int{1, 2, 3, 4, 5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
