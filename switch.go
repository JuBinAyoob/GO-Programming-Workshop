package main

import "fmt"

func main(){
	var count int
	fmt.Println(" Enter...")
	fmt.Scanf("%d",&count)
	switch count {
		case 1: fmt.Println("count is 1...")
		case 2,5: fmt.Println("count is 2 or 5...")
		default: fmt.Println("count is negative")
	}
}
