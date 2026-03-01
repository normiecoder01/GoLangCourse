package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	const(
		Sunday = iota + 1
		Monday = 2
		Tuesday = 37
		Wednesday = 39
		Thursday = 45
		Friday = 89
		Saturday = 32
	)
	fmt.Println("Sunday:", Sunday)
	fmt.Println("Monday:", Monday)
	fmt.Println("Tuesday:", Tuesday)
	fmt.Println("Wednesday:", Wednesday)
	fmt.Println("Thursday:", Thursday)
	fmt.Println("Friday:", Friday)
	fmt.Println("Saturday:", Saturday)

}
