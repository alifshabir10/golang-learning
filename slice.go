package main

import "fmt"



func main() {
	// initialisasi slice

	var fruits =[]string{"apple", "grape", "banana", "lemon"}
	var fruitsA =[]string{"apple", "grape", "banana", "lemon"}  // slice
	var fruitsB=[4]string{"apple", "grape", "banana", "lemon"} // array
	var fruitsC =[...]string{"apple", "grape", "banana", "lemon"} // array

	fmt.Println(fruits[0], fruitsA, fruitsB, fruitsC)

	// hubungan slice dan operasi slice
	var newFruits = fruits[0:2]

	fmt.Println(newFruits)

	// tipe data reference
	var aFruits = fruits[0:3]
	var bFruits = fruits[1:4]
	
	var aaFruits = aFruits[1:2]
	var baFruits = bFruits[0:1]

	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)
	
	// mengubah grape menjadi "pinnaple"
	baFruits[0] = "pinnaple"

	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(baFruits)

	// fungsi Len()
	fmt.Println(len(fruits))

	// fungsi cap()
	fmt.Println(cap(fruits))


	fmt.Println(len(aFruits))
	fmt.Println(cap(aFruits))
	fmt.Println(len(aaFruits))
	fmt.Println(cap(aaFruits))

	// fungsi append()

	var cFruits = append(fruits, "papaya")
	fmt.Println(cFruits)


	// fungsi copy()

	dst := make([]string, 3)
	src := []string{"watermelon","pinapple", "apple","orange"}
	n := copy(dst, src)
	
	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
}