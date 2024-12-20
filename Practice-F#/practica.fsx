type Department = 
    | Finance 
    | HumanResources 
    | InformationTechnology 
    | Marketing 
    | Sales

type Employee = 
    | HourlyEmployee of id: string * name: string * department: Department * hourlyPayRate: float * hoursWorked: int
    | Salesperson of id: string * name: string * department: Department * commissionRate: float * salesAmount: float

let calculateWeeklySalary employee =
    match employee with
    | HourlyEmployee(_, _, _, hourlyPayRate, hoursWorked) ->
        if hoursWorked > 40 then
            let overtimeHours = hoursWorked - 40
            (40.0 * hourlyPayRate) + (float overtimeHours * hourlyPayRate * 1.5)
        else
            float hoursWorked * hourlyPayRate
    | Salesperson(_, _, _, commissionRate, salesAmount) ->
        commissionRate * salesAmount

let employees = [
    HourlyEmployee("1111", "John Doe", InformationTechnology, 15.0, 40)
    HourlyEmployee("2222", "Alice Smith", Finance, 20.0, 45)
    Salesperson("3333", "Bob Johnson", Sales, 0.10, 5000.0)
    Salesperson("4444", "Carol White", Marketing, 0.12, 3000.0)
    HourlyEmployee("5555", "David Brown", HumanResources, 18.0, 38)
]

let displayEmployeeData employee =
    match employee with
    | HourlyEmployee(id, name, department, hourlyPayRate, hoursWorked) ->
        printfn "ID: %s, Name: %s, Department: %A, Hourly Pay Rate: %.2f, Hours Worked: %d, Weekly Salary: %.2f" 
            id name department hourlyPayRate hoursWorked (calculateWeeklySalary employee)
    | Salesperson(id, name, department, commissionRate, salesAmount) ->
        printfn "ID: %s, Name: %s, Department: %A, Commission Rate: %.2f, Sales Amount: %.2f, Weekly Salary: %.2f" 
            id name department commissionRate salesAmount (calculateWeeklySalary employee)

employees |> List.iter displayEmployeeData