package main

import "fmt"
import "errors"

var ErrDivisorIsLarger = errors.New("The divisor is larger so the answer will come 0")
func divide(a, b int) (int , error) {
	if b==0{
		return 0 , errors.New("Divide by zero.")
	}
	if(b > a){
		return 0, ErrDivisorIsLarger
	}
	return a/b, nil
}


func main(){

fmt.Println(divide(4,2))
fmt.Println(divide(2,4))
fmt.Println(divide(2,0))

}