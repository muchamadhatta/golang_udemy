package main

import "fmt"

func sayHello(number int) {
	fmt.Println("helo", number)
}

func main() {
	for i := 1; i <= 10; i++ {
		sayHello(i)
	}
}
