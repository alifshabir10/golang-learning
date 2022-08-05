package main

import "fmt"

func main(){
	// cara 1
	for i:=0; i<5; i++{
		fmt.Println("angka", i)
	}

	// cara 2
	var i = 0

	for i < 5 {
		fmt.Println("angak", i)
		i++
	}

	// for tanpa argumen
	// for {
	// 	fmt.Println("angka break", i)
	// 	i++
	// 	if i == 5{
	// 		break
	// 	}
	// }

	//  break & continue
	for i := 0; i <=10; i++{
		if i % 2 == 1  {
			continue
		}

		if i > 8 {
			break
		}
		fmt.Println("angka break & continue", i)
	}

	// perulangan bersarang
	for i := 0; i < 5; i++{
		for j := i; j < 5; j++{
		
			fmt.Print(j, "")
		}
		fmt.Println()
	}

	// pemanfaat label dalam perulangan 
	outerLoop:
	for i := 0; i < 5; i++{
		for j := 0; j < 5; j++{
		
			if i == 3{
				break outerLoop
			}
			fmt.Print("matriks [", i, "][", j , "]","\n")
		}
	}
}