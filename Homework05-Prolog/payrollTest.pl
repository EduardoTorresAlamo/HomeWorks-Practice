/*
 * File: payroll.pl
 * Author: Angel A. Cardona Cotto
 * Course: COTI 4039-VI1
 * Date: 12/17/2024
 * Purpose: This program displays the data of different employees.
 */

%% Declares the available departments.
department(finance).
department(human_Resources).
department(information_Technology).
department(marketing).
department(sales).

%% weekly_salary(+Employee, ?Salary)
%  Computes the weekly salary of an employee.
weekly_salary(hourlyEmp(_, _, _, hourlyPay(PayRate), hoursWorked(Hours)), Salary) :-
    (
    Hours > 40 ->  
    	Overtime is Hours - 40,
    	Salary is PayRate * 40.0 + PayRate * Overtime * 1.5
    	;   
    	Salary is PayRate * Hours
    ).
weekly_salary(salesPerson(_, _, _, commission(Com), sale(Sale)), Salary) :-
    Salary is Com * Sale.

%% write_inf(+Employee)
%  Displays the information of an employee.
write_inf(hourlyEmp(_, _, _, hourlyPay(PayRate), hoursWorked(Hours))) :-
    format("Hourly Pay Rate: ~2f~nHours Worked: ~d~n", [PayRate, Hours]).
write_inf(salesPerson(_, _, _, commission(Com), sale(Sale))) :-
    format("Commission: ~2f~nSales Amout: ~2f~n", [Com, Sale]).

%% write_data(+Employee)
%  Displays the data of an employee, including its weekly salary.
write_data(Employee) :-
    Employee =.. [Type, Id, Name, Dep|_],        % gets the first three components
    department(Dep),                             % checks that the department exits
    
    format("~nThe employee is an ~a~n", [Type]),
    format("Id: ~a~n", [Id]),
    format("Name: ~a~n", [Name]),
    format("Department: ~a~n", [Dep]),
    write_inf(Employee),
    
    weekly_salary(Employee, Salary),
    format("Weekly Salary: ~2f~n", [Salary]).

%% go
%  Serves as an entry point for the program.
go :-
    Employees = [
        hourlyEmp(1111, 'John Doe', finance, hourlyPay(10.50), hoursWorked(35)),
        salesPerson(2222, 'Jane Doe', sales, commission(0.15), sale(10500.50)),
        hourlyEmp(3333, 'Joe Mama', marketing, hourlyPay(7.50), hoursWorked(45)),
        salesPerson(4444, 'mark Johnson', sales, commission(0.25), sale(100000.01)),
        hourlyEmp(5555, 'Paul Blart', human_Resources, hourlyPay(1), hoursWorked(100))
    ],
    format("These are the Employees...~n"),
    forall(
   		member(Emp, Employees),
        write_data(Emp)
    ).