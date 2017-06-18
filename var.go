package main

import "fmt"

func main(){
	var num int
	num=20
	fmt.Println("num=== ",num,"  stored......")
	fmt.Printf("\n\nnum=== %d     stored......\n",num)

	var name string
	name="JuBeee...."

	fmt.Printf("\n %s \n",name)

	var aFloat float64
	var aBool bool
	aBool=true
	aFloat=3.22

	fmt.Println(aBool)
	bInt := 22
	fmt.Printf("\n %f %d %d\n",aFloat,bInt,num+bInt)
}
