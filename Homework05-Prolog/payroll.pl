#!/usr/bin/perl
use strict;
use warnings;

# Define employee data
my @employees = (
    {
        id => '1111',
        name => 'John Doe',
        department => 'Information Technology',
        type => 'hourly',
        hourly_pay_rate => 15.00,
        hours_worked => 40
    },
    {
        id => '2222',
        name => 'Jane Doe',
        department => 'Sales',
        type => 'salesperson',
        commission_rate => 0.15,
        sales_amount => 5000.00
    },
    {
        id => '3333',
        name => 'Alice Smith',
        department => 'Finance',
        type => 'hourly',
        hourly_pay_rate => 20.00,
        hours_worked => 45
    },
    {
        id => '4444',
        name => 'Bob Johnson',
        department => 'Marketing',
        type => 'salesperson',
        commission_rate => 0.10,
        sales_amount => 3000.00
    },
    {
        id => '5555',
        name => 'Charlie Brown',
        department => 'Human Resources',
        type => 'hourly',
        hourly_pay_rate => 18.00,
        hours_worked => 38
    }
);

# Calculate weekly salary
sub calculate_weekly_salary {
    my ($employee) = @_;
    if ($employee->{type} eq 'hourly') {
        my $hours = $employee->{hours_worked};
        my $rate = $employee->{hourly_pay_rate};
        if ($hours > 40) {
            return (40 * $rate) + (($hours - 40) * $rate * 1.5);
        } else {
            return $hours * $rate;
        }
    } elsif ($employee->{type} eq 'salesperson') {
        return $employee->{commission_rate} * $employee->{sales_amount};
    }
    return 0;
}

# Display employee data and weekly salary
foreach my $employee (@employees) {
    my $salary = calculate_weekly_salary($employee);
    print "ID: $employee->{id}\n";
    print "Name: $employee->{name}\n";
    print "Department: $employee->{department}\n";
    print "Weekly Salary: \$" . sprintf("%.2f", $salary) . "\n";
    print "-----------------------------\n";
}