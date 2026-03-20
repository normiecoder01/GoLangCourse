package main

import "fmt"
import "errors"

var ErrDivisorIsLarger = errors.New("The divisor is larger so the answer will come 0")

type ErrDivisorIsSame struct{
	Resource string
	ID int
	Message string
}

func (e ErrDivisorIsSame) Error() string {
	return e.Message
}


func divide(a, b int) (int , error) {
	if b==0{
		return 0 , errors.New("Divide by zero.")
	}
	if(b > a){
		return 0, ErrDivisorIsLarger
	}
	if (a == b){
		return 1,ErrDivisorIsSame{"Method: Divide",500,"Since the numerator and denominator are same the answer would be 1"}
	}
	return a/b, nil
}


func main(){

fmt.Println(divide(4,2))
fmt.Println(divide(2,4))
fmt.Println(divide(2,0))
fmt.Println(divide(4,4))

}