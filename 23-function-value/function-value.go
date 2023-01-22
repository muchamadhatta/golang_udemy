package main

import "fmt"

func getGoodbye(name string) string {
	return "Good Bye " + name
}

func main() {
	goodbye := getGoodbye
	fmt.Println(goodbye("ata"))
}
