package main

import "fmt"

func main() {
	var ujian = 90
	var absensi = 79
	//cara rumit
	// var lulusNilaiAkhir bool = nilaiAkhir > 80
	// var lulusAbsensi bool = absensi > 80
	// var lulus bool = lulusNilaiAkhir && lulusAbsensi
	// fmt.Println(lulus)
	//cara metode boolean yang simpel
	fmt.Println(ujian >= 80 && absensi >= 80)
}
