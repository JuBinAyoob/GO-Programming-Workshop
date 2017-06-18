package main

import "fmt"

func main(){
	var count int
	fmt.Println(" Enter...")
	fmt.Scanf("%d",&count)
	if count > 0 {
		fmt.Println(count," is positive... just simply ",count < 0)
 	}else if count < 0{
		fmt.Println(count," is negative")
	} else {
		fmt.Println(count," is zero....")
	}
}
