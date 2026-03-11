package main
import "fmt"

func main(){

if(false){
	fmt.Println("1==1 works")
}else{
	fmt.Println("1 doesn't work")
}

userAccess := map[string]bool {
	"Anupal": false,
	"Umale": true,
}

hasAccess, ok := userAccess["Umale"];

type User struct{
	Name string
	Age int
}
usersMap := map[string]User{
	"Anupal" : {"Anupal" , 29},
	"Sakshi" : {"Sakshi" , 25},
}
fmt.Println(usersMap)
fmt.Printf("%#v\n" , usersMap)
if(ok && hasAccess){
	fmt.Println("Access granted")
}else{
	fmt.Println("Access denied")
}







}