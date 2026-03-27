package main

import "fmt"

type Payable interface{
	CalculatePay() float64
	fmt.Stringer
}

	type SalariedEmployee struct {
		Name string
		Salary float64
	}

func (sal SalariedEmployee) CalculatePay() float64{
	return sal.Salary / 12.0
}

func (sal SalariedEmployee) String() string {
	return fmt.Sprint("Salaried:",sal.Name," Annual Salary:",sal.Salary)
}

type HourlyEmployee struct {
	Name string
	HourlyRate float64
	HoursWorked float64
}

func (he HourlyEmployee) CalculatePay() float64{
	return he.HourlyRate * he.HoursWorked
}

func (he HourlyEmployee) String() string{
	return fmt.Sprint("Hourly Employee:",he.Name," Hourly Rate:",he.HourlyRate," Hours Worked:", he.HoursWorked)
}

type CommisionedEmployee struct{
	Name string
	BaseSalary float64
	CommisionRate float64
	SalesAmount float64
}

func (ce CommisionedEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.SalesAmount * ce.CommisionRate )
}

func (ce CommisionedEmployee) String() string {
	return fmt.Sprint("Commissioned Employee:",ce.Name," Base Salary:",ce.BaseSalary," Commission Rate:",ce.CommisionRate, " Sales Amount:",ce.SalesAmount, )
}

// func(p Payable) CalculatePay() float64{
// 	return p.CalculatePay()
// }

func PrintEmployeeSummary [P fmt.Stringer] (employee P) {
	fmt.Println("Processing:",employee)
}

func ProcessPayroll (employees []Payable) {
	fmt.Println("--------------Processing Payroll--------------")
	totalPayroll:= 0.0
	for _ , employee := range employees{
		PrintEmployeeSummary(employee)
		pay:= employee.CalculatePay()
		fmt.Println("Monthly Pay :",pay)
		totalPayroll+= pay
	}
	fmt.Println("----------------------------------------")
	fmt.Println("Total Monthly Payroll:",totalPayroll)
}


func main() {

s1 := SalariedEmployee{
	Name:   "Amit Sharma",
	Salary: 85000,
}

s2 := SalariedEmployee{
	Name:   "Neha Patil",
	Salary: 92000,
}

h1 := HourlyEmployee{
	Name:        "Priya Deshmukh",
	HourlyRate:  500,
	HoursWorked: 160,
}

h2 := HourlyEmployee{
	Name:        "Rohit Kulkarni",
	HourlyRate:  650,
	HoursWorked: 140,
}

c1 := CommisionedEmployee{
	Name:          "Rahul Verma",
	BaseSalary:    30000,
	CommisionRate: 0.10,
	SalesAmount:   200000,
}

c2 := CommisionedEmployee{
	Name:          "Sneha Joshi",
	BaseSalary:    35000,
	CommisionRate: 0.08,
	SalesAmount:   150000,
}
employees:= [] Payable {s1,s2,c1,c2,h1,h2}
ProcessPayroll(employees)




}