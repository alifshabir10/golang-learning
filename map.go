package main

import "fmt"

func maps(){

	// how to use map
	var chicken map[string]int
	chicken = map[string]int{
		"january": 50,
		"february": 40,
		"maret": 46,
		"april": 54,
	}

	chicken["january"]= 50
	chicken["february"]= 40

	fmt.Println("january", chicken["january"])
	fmt.Println("mei", chicken["mei"])

	// inisialisasi map
	// var data map[string]int
	// data["one"]=1

	// data = map[string]int{}
	// data["one"]=1

	// cara horizontal
	// var chicken1 = map[string]int{"january":40 ,"february": 50}

	// cara vertical

	var chicken2 = map[string]int{
		"january": 40,
		"february": 50,
	}

	// integrasi for - range
	// for key, val := range chicken2 {

	// 	fmt.Println(key,"  \t:", val)
	// }
	


	// menghapus map

	fmt.Println(len(chicken2))
	fmt.Println(chicken2)
	
	delete(chicken2, "january")
	
	fmt.Println(len(chicken2))
	fmt.Println(chicken2)

	// mendeteksi index key

	var value, isExist = chicken2["mei"]

	if isExist{
		fmt.Println(value)

	} else {
		fmt.Println("item is not exist")
	}

	// kombinasi slice & map
	var chickens = []map[string]string{
		map[string]string{"name": "chicken red", "gender": "male"},
		map[string]string{"name": "chicken blue", "gender": "male"},
		map[string]string{"name": "chicken yellow", "gender": "female"},
	}
	for _, chicken  := range chickens{
		fmt.Println(chicken["gender"], chicken["name"])
	}
	
}