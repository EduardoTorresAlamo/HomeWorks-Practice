/*
 * File: payroll_file_io.go
 * Author: Eduardo R Torres
 * Course: COTI 4039-KJ1
 * Date: 10/03/2024
 * Purpose: This program evaluates a group of payroll employees to display their
 *          Info and Calculate their Salary.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

// Panics if an error occurs.
func checkForError(err error) {
	if err != nil {
		panic(err)
	}
}

// Verifies if the employee type is either 'H' or 'S'.
// Returns true if valid, false otherwise.
func verifyEmployeeType(employeeType string) bool {
	return employeeType == "H" || employeeType == "S"
}

// Parses an input line and returns the corresponding employee (either employeeHour or employeeSales).
func parseEmpleadoLine(line string) pago {
	fields := strings.Split(line, "|")

	if len(fields) < 6 {
		panic("Invalid input format")
	}

	// Verifying the employee type before proceeding
	if !verifyEmployeeType(fields[0]) {
		panic("Unknown employee type: " + fields[0])
	}

	id := fields[1]
	name := fields[2]

	deptVal, err := strconv.Atoi(fields[3])
	checkForError(err)
	dept := department(deptVal)

	// Check if it is an hourly worker (H) or a sales worker (S)
	if fields[0] == "H" {
		payRate, err := strconv.ParseFloat(fields[4], 64)
		checkForError(err)
		hoursWorked, err := strconv.Atoi(fields[5])
		checkForError(err)
		return employeeHour{
			employee:    employee{id: id, name: name, department: dept},
			hoursWorked: hoursWorked,
			payRate:     payRate,
		}
	} else if fields[0] == "S" {
		commissionRate, err := strconv.ParseFloat(fields[4], 64)
		checkForError(err)
		salesAmnt, err := strconv.ParseFloat(fields[5], 64)
		checkForError(err)
		return employeeSales{
			employee:       employee{id: id, name: name, department: dept},
			commissionRate: commissionRate,
			salesAmnt:      salesAmnt,
		}
	} else {
		panic("Unknown employee type")
	}
}

// Returns the evaluation line for the given employee in the requested format.
func newEvaluationLine(emp pago) string {
	switch e := emp.(type) {
	case employeeHour:
		return fmt.Sprintf("H|%s|%s|%d|%.2f\n", e.id, e.name, e.department, e.weeklySalary())
	case employeeSales:
		return fmt.Sprintf("S|%s|%s|%d|%.2f\n", e.id, e.name, e.department, e.weeklySalary())
	default:
		return "Unknown employee type\n"
	}
}

// Starts the execution of the program.
func main() {

	inFile, err := os.Open("employees.txt")
	checkForError(err)
	defer inFile.Close()

	outFile, err := os.Create("payroll.txt")
	checkForError(err)
	defer outFile.Close()

	scanner := bufio.NewScanner(inFile)
	for scanner.Scan() {
		empleado := parseEmpleadoLine(scanner.Text())
		outFile.WriteString(newEvaluationLine(empleado))
	}

}
