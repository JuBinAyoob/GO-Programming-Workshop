package main

import (
	"fmt"
	"time"
)
func sayHello(id int){
	fmt.Println(id,"Hello")
	time.Sleep(time.Second)
	fmt.Println(id,"World")
}
func main(){
	go sayHello(1)
	go sayHello(2)
	go sayHello(3)
	
	var x int
	fmt.Scanf("%d",&x)
}
