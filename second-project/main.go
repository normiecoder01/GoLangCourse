package main
import "fmt"
import "strings"
var productPrices = map[string]float64{
	"TSHIRT" : 399.00,
	"POWERBANK": 2499.00,
	"CAP" : 499.00,
	"MOBILE": 14999.00,
	"HEADPHONES": 2999.00,
	"SHOE CLEANER": 1499.00,
}
func calculateItemPrice(itemName string) (float64, bool) {
	basePrice, ok := productPrices[itemName];
	if(!ok){
		if(strings.HasSuffix(itemName, "_SALE")){
			originalItemName := strings.TrimSuffix(itemName, "_SALE")
			basePrice, ok:= productPrices[originalItemName]
			if(ok){
				salePrice := basePrice*0.9
				fmt.Printf("Item name : %s | Original Price:Rs. %.2f | Sales Price:Rs. %.2f\n",originalItemName,basePrice, salePrice )
				return salePrice, true
			}
		}
		fmt.Println("No product found for the search :", itemName);
		return 0.0 , false			
	}
	return basePrice, true	
}


func main(){

	fmt.Println("--------------Simple Sales Order Processor--------------")
	orderItems:= []string{"POWERBANK_SALE","CAP", "HEADPHONES", "SHOE CLEANER", "MOBILE_SALE", }
	var subtotal float64
	fmt.Println("-------PROCESSING ORDER ITEMS-------")
	for _,item:= range orderItems{
		price , ok := calculateItemPrice(item)
		if(ok){
			subtotal+=price;
		}
		fmt.Printf("Item Name : %-25sPrice:%10.2f\n",item,price)
	}
	fmt.Println("----------------------------------------------------------")
	fmt.Printf("Subtotal  :%42.2f\n",subtotal)










}