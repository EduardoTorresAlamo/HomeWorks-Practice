(*
 * File: payroll_file_io.fsx
 * Author: Eduardo R Torres 842-15-8724
 * Course: COTI 4039-KJ1
 * Date: 11/20/2024
 * Purpose: This program reads a group of employees from a file, evaluates
 *          them, and creates a new file with their salaries based on their position.
 *)

open System.IO

// Define the Department type
type Department = 
    | Finance
    | HumanResources
    | InformationTechnology
    | Marketing
    | Sales

// Function to parse department code
let parseDepartment (deptCode: int) =
    match deptCode with
    | 1 -> Finance
    | 2 -> HumanResources
    | 3 -> InformationTechnology
    | 4 -> Marketing
    | 5 -> Sales
    | _ -> failwith "Invalid department code"

// Function to convert department to code
let departmentToCode (department: Department) =
    match department with
    | Finance -> 1
    | HumanResources -> 2
    | InformationTechnology -> 3
    | Marketing -> 4
    | Sales -> 5

// Define the Employee type
type Employee =
    | HourlyEmployee of id: string * name: string * department: Department * hourlyRate: float * hoursWorked: int
    | Salesperson of id: string * name: string * department: Department * commissionRate: float * salesAmount: float

// Function to calculate weekly salary for an employee
let calculateWeeklySalary (employee: Employee) =
    match employee with
    | HourlyEmployee(_, _, _, hourlyRate, hoursWorked) ->
        if hoursWorked > 40 then
            let overtimeHours = hoursWorked - 40
            let overtimePay = float overtimeHours * hourlyRate * 1.5
            (40.0 * hourlyRate) + overtimePay
        else
            float hoursWorked * hourlyRate
    | Salesperson(_, _, _, commissionRate, salesAmount) ->
        commissionRate * salesAmount

// Function to parse a line of input into an Employee
let parseEmployee (line: string) =
    let fields = line.Split('|')
    match fields.[0] with
    | "H" ->
        let id = fields.[1]
        let name = fields.[2]
        let department = parseDepartment (int fields.[3])
        let hourlyRate = float fields.[4]
        let hoursWorked = int fields.[5]
        HourlyEmployee(id, name, department, hourlyRate, hoursWorked)
    | "S" ->
        let id = fields.[1]
        let name = fields.[2]
        let department = parseDepartment (int fields.[3])
        let commissionRate = float fields.[4]
        let salesAmount = float fields.[5]
        Salesperson(id, name, department, commissionRate, salesAmount)
    | _ -> failwith "Invalid employee type"
// Function to serialize an Employee into an output line
let serializeEmployee (employee: Employee) =
    match employee with
    | HourlyEmployee(id, name, department, hourlyRate, hoursWorked) ->
        let weeklySalary = calculateWeeklySalary employee
        sprintf "H|%s|%s|%d|%.2f" id name (departmentToCode department) weeklySalary
    | Salesperson(id, name, department, commissionRate, salesAmount) ->
        let weeklySalary = calculateWeeklySalary employee
        sprintf "S|%s|%s|%d|%.2f" id name (departmentToCode department) weeklySalary

// Main logic
let processPayroll inputFile outputFile =
    let inputLines = File.ReadAllLines(inputFile)
    let employees = inputLines |> Array.map parseEmployee
    let outputLines = employees |> Array.map serializeEmployee
    File.WriteAllLines(outputFile, outputLines)

// Process employees.txt and create payroll.txt
let inputFile = "employees.txt"
let outputFile = "payroll.txt"
processPayroll inputFile outputFile

printfn "Payroll file generated: %s" outputFile
