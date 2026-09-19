package main
import "fmt"

func gstSetter(slabPercent float64) func(price float64) float64 {

	return func(price float64) float64 {
		return  price + (price/100)*slabPercent
	}
}

func main(){
	slab18 := gstSetter(18)
	slab20 := gstSetter(20)
	finalPrice := slab18(500);
	finalPrice02 := slab20(500);

	fmt.Println(finalPrice); 
	fmt.Println(finalPrice02); 
}

