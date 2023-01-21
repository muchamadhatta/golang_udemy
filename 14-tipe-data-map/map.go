package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "ata",
		"address": "pamulang",
	}

	//nambah data
	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	// map dengan fungsi make
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Go-Lang"
	book["author"] = "Ata"
	book["ups"] = "Salah"

	fmt.Println(book)

	//menghapus map
	delete(book, "ups")
	fmt.Println(book)

	fmt.Println(book)
	fmt.Println(len(book))
}
