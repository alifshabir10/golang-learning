package main

import (
	"fmt"
)

// func
//  main(){
// 	var names =[]string{"John", "wick"}
// 	printMessage("halo", names)

// }

// func printMessage(message string, arr []string){
// 	var nameString= strings.Join(arr, " ")
// 	fmt.Println(message, nameString)
// }

// return value/ nilai balik
// func main(){
// // penggunaan rand.Seed()

// 	rand.Seed(time.Now().Unix())
// 	var randomValue int

// 	randomValue = randomValueRange(2, 10)
// 	fmt.Println("random number:",randomValue)
// 	randomValue = randomValueRange(2, 10)
// 	fmt.Println("random number:",randomValue)
// 	randomValue = randomValueRange(2, 10)
// 	fmt.Println("random number:",randomValue)

// }

// func randomValueRange(min, max int) int{
// 	var value= rand.Int()% (max-min +1) + min
// 	return value

// }

// // declarasik variabel bertipe data sama
// func nameOfFunc(paramA type, paramB type, paramC type) returnType
// func nameOfFunc(paramA, paramB, paramC type)returnType

// func randomValueRange(min int, max int  )int
// func randomValueRange(min, max int)int

// penggunaan return

func fungsi(){
	
	divideNumber(10,2)
	divideNumber(4,0)
	divideNumber(8,-4)

}

func divideNumber(m, n int){
	if n == 0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}
	var res = m/n
	fmt.Printf("%d / %d = %d\n", m, n, res)
}

