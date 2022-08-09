package main

import (
	"fmt"
	"math"
)

// func calculate(d float64)(float64, float64){
// 	// hitung luas
// 	var area = math.Pi * math.Pow(d / 2, 2)
// 	// hitung keliling
// 	var circumference = math.Pi * d

// 	// kembalikan 2 nilai
// 	return area, circumference

// }
 func main(){
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkarang \t\t: %.2f \n", area)
	fmt.Printf("keliling lingkarang \t: %.2f \n", circumference)
 }

//  perendenfined return value
func calculate(d float64)(area float64, circumference float64){

	area= math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi* d

	return
}