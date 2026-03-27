package main
import "fmt"
type Person interface {
	getFullName() string
}

type Employee struct{
	ID	int
	FirstName string
	LastName string
	Department string
	Company string
}

type Employer struct {
	ID	int
	FirstName string
	LastName string
	Department string
	Company string
}

func (emp Employee) getFullName() string{
	return fmt.Sprintf("%s %s",emp.FirstName,emp.LastName)
}

func (emp Employer) getFullName() string{
	return fmt.Sprintf("%s %s",emp.FirstName,emp.LastName)
}

func main(){
	anupal:= Employee {
	ID : 1,
	FirstName : "Anupal",
	LastName : "Umale",
	Department : "ITH/IF",
	Company : "Mercedes Benz",
	}

	anil:= Employer {
	ID : 2,
	FirstName : "Anil",
	LastName : "Umale",
	Department : "Internal",
	Company : "Mercedes Benz",
	}

	fmt.Println(anupal.getFullName())
	fmt.Println(anil.getFullName())

}