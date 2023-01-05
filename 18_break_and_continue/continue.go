package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // ketika mendapatkan nilai genap akan di continue
		}
		fmt.Println("Perulangan ke", i)
	}
}
