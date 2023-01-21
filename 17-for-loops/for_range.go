package main

import "fmt"

func main() {
	var names = []string{"ata", "muchamad", "hatta"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
