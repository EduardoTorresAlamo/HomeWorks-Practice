/*
 * File: payroll.go
 * Author: Eduardo R Torres
 * Course: COTI 4039-KJ1
 * Date: 10/03/2024
 * Purpose: This program evaluates a group of payroll employees to display their
 *          Info and Calculate their Salary.
 */

package main

import "fmt"

type department int

const (
	Finance department = iota
	HumanResources
	InformationTechnology
	Marketing
	Sales
)

// Returns the string representation of the department.
func (e department) String() string {
	if Finance > e || e > Sales {
		return "unknown department"
	}
	departments := [...]string{"Finance", "Human Resources", "Informartion Technology", "Marketing", "Sales"}
	return departments[e]
}

type employee struct {
	id   string
	name string
	department
}

type pago interface {
	weeklySalary() float64
}

type employeeHour struct {
	employee
	hoursWorked int
	payRate     float64
}

func (e employeeHour) weeklySalary() float64 {
	overTime := e.hoursWorked - 40
	overRate := e.payRate * 1.5
	extraPay := 0.0

	if overTime > 0 {
		extraPay = float64(overTime) * overRate
	}

	return e.payRate*float64(e.hoursWorked) + extraPay
}

type employeeSales struct {
	employee
	salesAmnt      float64
	commissionRate float64
}

func (e employeeSales) weeklySalary() float64 {
	return e.commissionRate * float64(e.salesAmnt)
}

func printData(emp pago) {
	// Cast the pago interface to employeeHour to access employee details
	if empData, ok := emp.(employeeHour); ok {
		fmt.Println()
		fmt.Println("Id:", empData.id)
		fmt.Println("Name:", empData.name)
		fmt.Println("Department:", empData.department)
		fmt.Println("Hours Worked:", empData.hoursWorked)
		fmt.Printf("Pay Rate: %.2f\n", empData.payRate)
		fmt.Printf("Salary: $%.2f\n", empData.weeklySalary())

	} else if empDataS, ok := emp.(employeeSales); ok {
		fmt.Println()
		fmt.Println("Id:", empDataS.id)
		fmt.Println("Name:", empDataS.name)
		fmt.Println("Department:", empDataS.department)
		fmt.Println("Sales Ammount:", empDataS.salesAmnt)
		fmt.Printf("Commission Rate: %.2f\n", empDataS.commissionRate)
		fmt.Printf("Salary: $%.2f\n", empDataS.weeklySalary())

	} else {
		fmt.Printf("Unknown employee type")
		fmt.Println()
	}
}

//  Starts the execution of the program.
func main() {
	pagos := []pago{
		employeeHour{employee: employee{id: "1111", name: "John Doe", department: Marketing}, hoursWorked: 40, payRate: 15.0},
		employeeSales{employee: employee{id: "2222", name: "Johnny Doesn't", department: Sales}, salesAmnt: 5000.00, commissionRate: .15},
		employeeHour{employee: employee{id: "3333", name: "Johnnatan Doever", department: HumanResources}, hoursWorked: 45, payRate: 22.0},
		employeeSales{employee: employee{id: "4444", name: "John Joestar", department: Sales}, salesAmnt: 38, commissionRate: 1.0},
		employeeHour{employee: employee{id: "5555", name: "Johnny Does", department: Finance}, hoursWorked: 38, payRate: 20.0},
		employeeSales{employee: employee{id: "6666", name: "Johnnatan Jostar", department: Sales}, salesAmnt: 38, commissionRate: .10},
	}

	fmt.Println("These are the Employees...")
	for _, pago := range pagos {
		printData(pago)
	}
}
