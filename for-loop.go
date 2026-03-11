package main
import "fmt"
func main(){
//C -style loop
fmt.Println("--------C Style Loop--------")
for i:= 1; i<=5 ; i++{
	fmt.Println(i)
}

//while style loop
fmt.Println("--------While Style Loop--------")

j:=0
for j<=5{
	fmt.Println(j)
	j++
}

//infinite loop
fmt.Println("--------Infinite Loop--------")

k:=5
for{
fmt.Println(k)
k--
if(k<= 0){
	break;
}
}
fmt.Println("--------Using for with range--------")

items :=  [3]string{"Go","Python","Java"} 

for index,value:= range items {
	fmt.Println( index, value)
}












}
