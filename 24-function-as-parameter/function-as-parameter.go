package main

import "fmt"

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("hello", filter(name))

}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("ata", spamFilter)
	filter := spamFilter
	sayHelloWithFilter("anjing", filter) //jika kata anjing yang di ketik maka akan di filter jadi ...
}
