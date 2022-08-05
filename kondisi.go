package main

import "fmt"

func main(){
	var point = 10

	if point == 10{
		fmt.Println("lulus dengan nilai sempurna")

	} else if point > 5{
		fmt.Println("lulus")
	} else if point  == 4 {
		fmt.Println("hampir tidak lulus")
	} else {
		fmt.Printf("tidak lulus. nilai anda %d\n", point)
	}

	// percent
	var point2 = 8840.0

	if percent := point2 / 100; percent >=100{
		fmt.Printf("%.1f%s perfect!\n",percent , "&")
		} else if percent >=70 {
			fmt.Printf("%.1f%s good\n",percent , "&")
			}else {
				fmt.Printf("%.1f%s not bad\n",percent , "&")
				

	}

	// switch -case

	switch point {
	case 8:
		fmt.Println("perfect")
		
	case 7:
		fmt.Println("awesome")
		
		
	default:
		fmt.Println("not bad")

	}

	// pemanfaatan case and ,
	switch point {
	case 8:
		fmt.Println("perfect")
		
	case 7,6,5,4:
		fmt.Println("awesome")
		
	default:
		fmt.Println("not bad")

	}

	// pemanfaatan case and default
	switch point {
	case 8:
		fmt.Println("perfect")
		
	case 7,6,5,4:
		fmt.Println("awesome")
		
	default:
		{fmt.Println("not bad")
		fmt.Println("you can be better")}

	}

	// switch dengan if || else
	switch {
	case point == 8:
		fmt.Println("perfect")
		
	case (point < 8) && (point >3):
		fmt.Println("awesome 1")
		
	default:
		{fmt.Println("not bad")
		fmt.Println("you can to learn more")}

	}

		// switch dengan fallthrough
	switch {
	case point == 8:
		fmt.Println("perfect")
		
	case (point < 8) && (point >3):
		fmt.Println("awesome")
		fallthrough
		
	case point < 3 :
		fmt.Println("you need to learn more")
	default:
		{fmt.Println("not bad")
		fmt.Println("you can to learn more")}

	}

	// kondisi bersarang
if point > 7{
	switch point {
	case 10: 
	fmt.Println("Perfect !")

	default:
		fmt.Println("Nice !")
	}
} else {
	if point == 5 {
		fmt.Println("not bad")
		} else if point == 3 {
		fmt.Println("keep trying")

	}else {
		fmt.Println("you can do it")
		if point == 0{
			fmt.Println("try harder !")
		}
	}
}
	
}