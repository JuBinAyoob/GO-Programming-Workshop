package main

import( 
	"fmt"
)

type Person struct {
	firstName string
	lastName string
	age int
}
func sayHello(per Person) string {
	helloString := fmt.Sprintf("Hello %s %s %d\n",per.firstName,per.lastName,per.age)
	return helloString
}

func main(){
	
	p1 := Person{firstName : "JuBin",lastName :"Ayoob", age:22}
	fmt.Println(sayHello(p1))
}
