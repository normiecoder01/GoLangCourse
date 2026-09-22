package main
import "fmt"

type Contact struct{
	ID int
	Name string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID = 1

func init(){
	contactList = make([]Contact , 0)
	contactIndexByName = make(map[string]int)

}

func addContact(name, email, phone string) {
	_,exists:= contactIndexByName[name]
	if(exists){
		println("Contact by the name",name , "already exists.")
		return
	}
	newContact:= Contact{
		Name  : name,
		Email : email,
		Phone : phone,
		ID    : nextID,
	}
	contactList = append(contactList, newContact)
	contactIndexByName[newContact.Name] = (len(contactList)-1)
	nextID++
	fmt.Println(name, "added to the contact  successfully.")

}

func findContactByName(name string) *Contact {
	index ,exists:= contactIndexByName[name]
	if(exists){
		foundContact := contactList[index];
		return &foundContact; 
	} else{
		return nil
	}
}

func nameSakeFunc(parameter int) float64 {
	floatConvertedInt:= float64(parameter)
	return floatConvertedInt
}

var funcName = func(parameter int) float64 {
	return float64(parameter)
}


func listContacts(){
	println("-----All Contacts----")
	for _ , contact := range contactList{
		println("ID:",contact.ID,"Name:",contact.Name,"Phone:",contact.Phone,"Email:",contact.Email)
	}
	println("---------------------")
}
func main(){
	addContact("Anupal Umale","anupalumale@gmail.com","7028843375")
	addContact("Ashutosh Paliwal","ashutoshpaliwal@gmail.com","7721442935")
	addContact("Pallavi Umale","pallaviumale@gmail.com","8788204107")
	addContact("Anil Umale","anilumale@gmail.com","9404534307")
	addContact("Anupal Umale","anupalumale@gmail.com","7028843375")

	contact:= findContactByName("Ashutosh Paliwal")
	if(contact==nil){
		fmt.Println("contact by the name Ashutosh was not found")
	} else{
		println("ID:",contact.ID,"Name:",contact.Name,"Phone:",contact.Phone,"Email:",contact.Email)
	}
	contact= findContactByName("Avinash")
		if(contact==nil){
		fmt.Println("contact by the name Ashutosh was not found")
	} else{
		println("ID:",contact.ID,"Name:",contact.Name,"Phone:",contact.Phone,"Email:",contact.Email)
	}

	listContacts()

	returnValue:= nameSakeFunc(2)
	returnedValue := funcName(5)

	fmt.Println(returnValue)
	fmt.Println(returnedValue)


}