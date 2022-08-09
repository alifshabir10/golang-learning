package main

import "fmt"
func main(){
	var avg = calculate(2,3,4,3,2,3,2,3,2,3,2,)
	var msg = fmt.Sprintf("Rata - rata : %.2f", avg)
	fmt.Println(msg)
}

func calculate(numbers ...int)float64{
	var total int = 0
	for _, number := range numbers{
		total += number
	}

	var avg = float64(total)/ float64(len(numbers))
	return avg

}