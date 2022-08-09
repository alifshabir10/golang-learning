package main

import "fmt"

func news() {
	// menggunakan var, tanpa tipe data, menggunakan perantara "="
	var firstName string = "john"

	var lastName string 

	// tanpa var, tanpa tipe data, menggunakan perantara ":="
	lastName = "wick"

	// nilai dapat di isi secara bersamaan
	var first, second, third string

	fmt.Println("halo %s %s!\n", firstName, lastName, first, second, third)

	//  contoh lain
	var fourth, fifth, sixth string = "empat", "lima", "enam"
	first, second, third = "satu", "dua", "tiga"

	fmt.Println("halo %s %s!\n", fourth, fifth,sixth)


	//  contoh type inference

	one, isFriday, twoPointTwo, say := 1, true, 2.2, "hello"
	fmt.Println(one, isFriday,twoPointTwo, say)

	//  contoh underscore
	_= "belajar golang"
	_= "golang itu mudah"
	name, _ := "john", "wick"


	fmt.Println(name, )

	// deklarasi variabel "new"
	nama := new(string)

	fmt.Println(nama)
	fmt.Println(*nama)
 
	// deklarasi variabel "make"
	type VSlice []int
	b := make([]int, 0, 5)
	// fmt.printSlice("b", b)
	fmt.Print(VSlice(b))
}