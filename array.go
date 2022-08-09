package main

import "fmt"

func array() {
	

	// gaya vertical
	var fruits  [4]string
	// cara horizontal
	 fruits = [4]string{"apple", "grape", "banana", "melon"}
	
	// cara vertical
	 fruits = [4]string{
		"apple",
		 "grape",
		"banana",
		"melon"}
	
		   fmt.Println("Jumlah element \t\t", len(fruits))
		   fmt.Println("Isi semua element \t", fruits)
		   

		//    var array tanpa jumlah element
		   var numbers = [...]int{2, 3, 2, 4,3}

		   fmt.Println("data array \t:",numbers)
		   fmt.Println("jumlah elemen \t:",len(numbers))

		//    array multidimensi
		var number1 = [2][3]int{[3]int{3,2,3}, [3]int{3,4,5}}
		var number2 = [2][3]int{{3,2,3},{3,4,5}}

		fmt.Println("number1", number1)
		fmt.Println("number2", number2)

		// perulangan menggunak for
		for i := 0; i<len(fruits); i++{
			fmt.Printf("element %d : %s\n", i, fruits[i] )
		}	

		// perulangan menggunak for - range
		for i , fruit := range fruits{
			fmt.Printf("element %d : %s\n", i, fruit )
		}	

		// perulangan menggunakan variabel underscore _ di for - range
		for _ , fruit := range fruits{
			fmt.Printf("nama buah : %s\n",  fruit )
		}	
		for i, _  := range fruits{
			fmt.Printf("nama buah : %s\n",  i )
		}	
		for i := range fruits{
			fmt.Printf("nama buah : %s\n",  i )
		}	

		// perulanggan menggunakan make
		var fruitss = make([]string, 2)
		fruits[0]= "apple"
		fruits[1] = "manggo" 	

		fmt.Println(fruitss)
}