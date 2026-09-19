package main

import "fmt"
func incrementValue (ptrToInt *int){
	*ptrToInt = *ptrToInt + 1
	fmt.Println(*ptrToInt)
}
func main(){
	x := 10
	var xAdd *int
	xAdd = &x
incrementValue(xAdd)

}