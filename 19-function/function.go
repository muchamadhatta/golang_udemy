package main

import "fmt"

func sayHello() {
	fmt.Println("halo")
}

func main() { // func main untuk mengeksekusi
	sayHello() //panggil fungsi
	sayHello()
	sayHello()

	for i := 0; i < 10; i++ { //bisa juga di panggil secara perulangan
		sayHello()
	}

}
