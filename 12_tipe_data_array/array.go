package main

import "fmt"

// index array di mulai dari nol
func main() {
	var names [3]string

	names[0] = "ata"
	names[1] = "ataa"
	names[2] = "ataaaa"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//menyimpan 3 array
	var values = [3]int{
		90,
		95,
		80,
	}
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	//menghitung panjang array
	fmt.Println(len(names))
	fmt.Println(len(values))

	//menghitung panjang string di array
	var lagi [10]string //hasilnya akan 10 karena panjang array 10
	fmt.Println(len(lagi))
}
