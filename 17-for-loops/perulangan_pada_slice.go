package main

import "fmt"

func main() {

	var slice = []string{"ata", "muchamad", "hatta"} //data slice
	for i := 0; i < len(slice); i++ {                //dipanggil secara berulang
		fmt.Println(slice[i])
	}
}
