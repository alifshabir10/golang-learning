package main

import "fmt"

func main(){
	var getMinMax = func(n []int) (int, int){

		var min, max int
		for i, e := range n {
			switch{
			case i == 0:
				max, min = e, e
			case e > max:
				max = e

			case e < min:
				min = e

}
		}
		return min, max

	}

	var number= []int{2,3,2,3,3,4,3,2,3,2,2}
	var min, max = getMinMax(number)

	fmt.Printf("data : %v\nmin %v\nmax : %v\n ",number, min, max)

	// Immediately-Invoked Function Expression (IIFE)
	var newNumbers = func (min int)[]int {
		var r []int
		for _, e := range number{
			if e < min {
				continue
			}

			r = append(r, e)
		}
		return r 

	} (3)

	fmt.Println("original number: ", number)
	fmt.Println("filtered number: ", newNumbers)

	// var max = 3
	var howMany, getNumbers = findMax(number,max)
	var theNumber = getNumbers

	fmt.Println("number \t",number)
	fmt.Printf("find \t: %d\n\n", max)
	fmt.Println("found \t", howMany)
	fmt.Println("value\t", theNumber)

	
	
}
// closure sebagai nilai kembalian

func findMax( numbers []int, max int )(int, func() []int){
	var res []int
	for _,e := range numbers{
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}

