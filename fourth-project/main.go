package main

import "fmt"
import "strings"

const(
	division = "Division"
	divisionErr = "Division by 0 is not allowed"
)

type MathError struct {
Operation string
InputA int
InputB int
Message string
}

func (err MathError) Error() string {
var inputs []string
if err.Operation == division{
	inputs = append(inputs,fmt.Sprintf("a=%d",err.InputA))
	inputs = append(inputs,fmt.Sprintf("b=%d",err.InputB))
}
return fmt.Sprintf("Math error in : %s (%s) : %s",err.Operation, strings.Join(inputs,",") , err.Message)
}

func sum(numbers ...int) int {
	total:= 0
	for _ ,number:= range numbers {
		total+= number
	}
	return total
}

func safeDivision (a,b int) (int, error) {
	if b== 0{
		return 0,&MathError{
			Operation: division,
			InputA: a,
			InputB: b,
			Message: divisionErr,
		}
	}

	return a/b, nil
}
func main(){
fmt.Println(sum(2,3,4))
fmt.Println(safeDivision(20,10))
fmt.Println(safeDivision(20,0))








}