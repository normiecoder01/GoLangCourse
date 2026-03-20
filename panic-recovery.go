package main

import "fmt"

func onlyPanic() {
	defer fmt.Println("Defering statement  from onlyPanic for testing purpose before panic")
	panic("onlyPanic panicked")
		defer fmt.Println("Defering statement  from onlyPanic for testing purpose after panic")

}

func recoverablePanic(){

		defer func() {
		if err:= recover(); err!= nil{
			fmt.Println("Recovered the error:", err)
		}
	}()

	defer fmt.Println("Defering statement  from recoverablePanic for testing purpose before panic")

	panic("This panic would be recovered. (Hopefully)")

	defer fmt.Println("Defering statement  from recoverablePanic for testing purpose after panic")

}


func main() {

recoverablePanic()

onlyPanic()






}