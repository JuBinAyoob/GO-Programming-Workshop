package main

import "fmt"

func printAll(arr []int){
	for i :=0; i<len(arr);i++ {
		fmt.Printf("%d =>  %d\n",i,arr[i])
	}
	j:=0
	for j<len(arr){
		fmt.Printf("%d =>%d\n",j,arr[j])
		j++
	}
	fmt.Println("Range")
	for index,value := range arr{
		fmt.Printf("%d => %d\n",index,value)	
	}

}
func main(){
	var arr [5] int
	fmt.Println(arr)

	arr[1] = 1
	arr[3] = 3
	fmt.Println(arr)

	//barr := [5] int {5,4,3,2}
	barr := [...] int {5,4,3,2}
	fmt.Println(barr)

	aSlice := []int{1,2,3}
	fmt.Println(aSlice)
	aSlice = append(aSlice,4)
	fmt.Println(aSlice)

	bSlice := make([]int,5)
	bSlice = append(bSlice,5)
	bSlice = append(bSlice,4)
	bSlice[1]=1;
	fmt.Println("BSlice",bSlice)

	printAll(aSlice)

	fmt.Println("-----------------------------------------------")
	barrrr := [...]int{5,4,2,2,200,103,299}
	fmt.Println(barrrr[2:5])
}
