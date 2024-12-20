(*
 * File: payroll.fsx
 * Author: Eduardo R Torres 842-15-8724
 * Course: COTI 4039-KJ1
 * Date: 11/20/2024
 * Purpose: This program reads a group of employees, evaluates
 *          them, and presents them with their salaries based on their position.
 *)

// Se Define el tipo de dato Department que representa los departamentos de la empresa
type Department = 
    | Finance 
    | HumanResources 
    | InformationTechnology 
    | Marketing 
    | Sales

// Se Define el tipo de dato Employee que representa los empleados de la empresa ya sea por hora o por comisión
type Employee = 
    | HourlyEmployee of id: string * name: string * department: Department * hourlyRate: float * hoursWorked: int
    | Salesperson of id: string * name: string * department: Department * commissionRate: float * salesAmount: float

// Se Define calculateWeeklySalary. Calcula salario semanal de un empleado
let calculateWeeklySalary employee =
    match employee with
    | HourlyEmployee(_, _, _, hourlyRate, hoursWorked) ->
        if hoursWorked > 40 then
            let overtimeHours = hoursWorked - 40
            (40.0 * hourlyRate) + (float overtimeHours * hourlyRate * 1.5)
        else
            float hoursWorked * hourlyRate
    | Salesperson(_, _, _, commissionRate, salesAmount) ->
        commissionRate * salesAmount

// Estos son los empleados de la empresa
let employees = [
    HourlyEmployee("1111", "Eduardo Torres", InformationTechnology, 15.0, 40)
    HourlyEmployee("2222", "Elma Quinon", Finance, 20.0, 45)
    Salesperson("3333", "Elca Vallo", Sales, 0.10, 5000.0)
    Salesperson("4444", "Elsa Capuntas", Marketing, 0.12, 3000.0)
    HourlyEmployee("5555", "Vece Rin", HumanResources, 18.0, 38)
]

let displayEmployeeData employee =
    match employee with
    | HourlyEmployee(id, name, department, hourlyRate, hoursWorked) ->
        printfn "ID: %s, Name: %s, Department: %A, Hourly Rate: %.2f, Hours Worked: %d, Weekly Salary: %.2f" 
            id name department hourlyRate hoursWorked (calculateWeeklySalary employee)
    | Salesperson(id, name, department, commissionRate, salesAmount) ->
        printfn "ID: %s, Name: %s, Department: %A, Commission Rate: %.2f, Sales Amount: %.2f, Weekly Salary: %.2f" 
            id name department commissionRate salesAmount (calculateWeeklySalary employee)

employees |> List.iter displayEmployeeData
