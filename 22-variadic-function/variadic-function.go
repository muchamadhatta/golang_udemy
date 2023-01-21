package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(50, 50, 40) //total dari ini
	fmt.Println(total)

	slice := []int{10, 20, 30}
	total = sumAll(slice...)
	fmt.Println(total)
}
