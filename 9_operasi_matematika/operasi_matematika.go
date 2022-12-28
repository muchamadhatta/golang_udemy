package main

import "fmt"

func main() {

	//pengjumlahan langsung
	var result = 100 + 100
	fmt.Println(result)

	//pernjumlahan terpisah, begitupun dengan pengurangan dan sebagiannya
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	//augmented assignment
	var i = 10
	i += 10 //artinya i=i+10
	fmt.Println(i)

	//unary operator
	i++            //i= i+1
	fmt.Println(i) //21

	//membuat negatif dan positif
	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}
