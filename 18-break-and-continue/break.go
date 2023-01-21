package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //untuk menghentikan program pada kondisi 5
		}
		fmt.Println("Perulangan ke ", i)
	}
}
