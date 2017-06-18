package main

import "fmt"

func sayHello(name string,lastName string) (string,int) {
	helloString := fmt.Sprintf("Hello %s %s\n",name,lastName)
	return helloString,len(helloString)
}

func main(){
	
	fmt.Println(sayHello("JuBee....","Ayoob"))

	str, length := sayHello("Beegita ","Arubam...")
	fmt.Println(length,str)

	strr, _ := sayHello("JuBin ","Ayoob...")
	fmt.Println(strr)
}
