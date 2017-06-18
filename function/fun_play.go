package main

import( 
	"fmt"
	"errors"
)

func remainder(a,b int) (int, error) {
	if b == 0 {
		return 0,errors.New("Division by zero not possible")
	}
	return a%b,nil
}

func main(){
	
	//fmt.Println(sayHello("JuBee....","Ayoob"))

	rem,err := remainder(5,0)
	if err != nil {
		panic(err)
	}
	
	fmt.Println(rem)
	
}
