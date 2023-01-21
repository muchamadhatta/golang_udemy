package main

import "fmt"

func getHello(name string) string {
	return "Hello " + name
}
func main() {
	result := getHello("ata")
	fmt.Println(result)
}
