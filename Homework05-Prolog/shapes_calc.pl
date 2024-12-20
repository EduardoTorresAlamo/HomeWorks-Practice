/*
 * File: shapes_calc.pl
 * Author: Antonio F. Huertas 
 * Course: COTI 4039-KJ1
 * Date: 12/10/2024
 * Purpose: This program displays the data, area and perimeter of a list of shapes.
 */

%% Declares the available shape colors.
color(white).
color(red).
color(green).
color(blue).
color(yellow).
color(black).

%% shape_area(+Shape, ?Area)
%  Computes the area of a shape.
shape_area(circle(_, _, radius(Rad)), Area) :-
    Area is pi * Rad ^ 2.
shape_area(rectangle(_, _, width(Wdt), height(Hgt)), Area) :-
    Area is Wdt * Hgt.

%% shape_perimeter(+Shape, ?Perim)
%  Computes the perimeter of a shape.
shape_perimeter(circle(_, _, radius(Rad)), Perim) :-
    Perim is 2 * pi * Rad.
shape_perimeter(rectangle(_, _, width(Wdt), height(Hgt)), Perim) :-
    Perim is 2 * (Wdt + Hgt).
    
%% write_dimensions(+Shape)
%  Displays the dimensions of a shape.
write_dimensions(circle(_, _, radius(Rad))) :-
    format("Radius: ~1f~n", [Rad]).
write_dimensions(rectangle(_, _, width(Wdt), height(Hgt))) :-
    format("Width: ~1f, Height: ~1f~n", [Wdt, Hgt]).

%% write_data(+Shape)
%  Displays the data of a shape, including its area and perimeter.
write_data(Shape) :-
    Shape =.. [Type, Color, Loc|_],        % gets the first three components
    color(Color),                          % checks that the color exits
    
    format("~nThe shape is a ~a~n", [Type]),
    format("Color: ~a~n", [Color]),
    format("Location: ~w~n", [Loc]),
    write_dimensions(Shape),
    
    shape_area(Shape, Area),
    shape_perimeter(Shape, Perim),
    format("Area: ~3f, Perimeter: ~3f~n", [Area, Perim]).

%% go
%  Serves as an entry point for the program.
go :-
    Shapes = [
        circle(red, point(2, 1), radius(5)),
        rectangle(yellow, point(8, 3), width(4), height(9)),
        circle(blue, point(4, 5), radius(7))
    ],
    format("These are the shapes...~n"),
    forall(
        member(Shape, Shapes),
        write_data(Shape)
    ).