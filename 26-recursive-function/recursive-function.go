package main

import "fmt"

func factorialloop(value int) int {
	result := 1
	for i := value; i > 0; i-- {
		result *= i //result = result * i
	}
	return result
}

func main() {
	loop := factorialloop(5)
	fmt.Println(loop)
	fmt.Println(5 * 4 * 3 * 2 * 1)
}
