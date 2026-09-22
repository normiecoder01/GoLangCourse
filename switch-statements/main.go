package  main
import "fmt"
import "math/rand"

func main(){
days:= [7]string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"};

day:= days[rand.Intn(7)]
fmt.Println("Today is", day)

switch day{
case "Sunday", "Saturday":
	fmt.Println("Weekend Woohoo!!!")
case "Monday" , "Tuesday":
	fmt.Println("I need motivation...")
case "Wednesday" , "Thursday":
	fmt.Println("Holding my breath for the weekend")
case "Friday":
	fmt.Println("Making plans already...")
}

checkType :=func(t interface{}) {
	switch v:= t.(type){
	case int:
		fmt.Printf("Integer: %d\n", v )
	case string:
		fmt.Printf("String: %s\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type %T\n", v)
	}
}
checkType("Hello")
checkType(true)
checkType(32)
checkType(32.22)





}