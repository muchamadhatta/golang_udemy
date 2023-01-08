package main

import "fmt"

func sayHello(name string) string {
	if name == "" {
		return "halo bro"
	} else if name == "ata" {
		return "halo ata"
	} else {
		return "halo " + name
	}
}

func main() {
	result := sayHello("boy")
	fmt.Println(result)
	fmt.Println(sayHello(""))
	fmt.Println(sayHello("ata"))
}
