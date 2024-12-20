/*
 * File: my_list_preds.pl
 * Author: Antonio F. Huertas
 * Course: COTI 4039-KJ1
 * Date: 12/05/2024
 * Purpose: This program defines some list predicates.
 */

%% list(?List)
%  Succeeds if the argument is a list.
%  Note: This is a tail-recursive predicate.
list([]).
list([_|Tail]) :- list(Tail).

%% list_length(+List, ?Length)
%  Computes the number of elements in a list.
list_length([], 0).
list_length([_|Tail], Length) :-
    list_length(Tail, LengthTail),
    Length is 1 + LengthTail.

%% list_length_v2(+List, ?Length)
%  Computes the number of elements in a list efficiently.
list_length_v2(List, Length) :- length_aux(List, 0, Length).

%  length_aux(+List, +Accum, ?Length)
%  TODO Complete this predicate.

%% list_sum(+NumList, ?Sum)
%  Computes the sum of the elements in a list of numbers.
list_sum([], 0).
list_sum([Head|Tail], Sum) :-
    list_sum(Tail, SumTail),
    Sum is Head + SumTail.

%% list_sum_v2(+NumList, ?Sum)
%  Computes the sum of the elements in a list of numbers efficiently.
list_sum_v2(NumList, Sum) :- sum_aux(NumList, 0, Sum).

%  sum_aux(+NumList, +Accum, ?Sum)
%  TODO Complete this predicate.

%% member_list(?Elem, +List)
%  Enumerates all members of a list or determines if an element is a member.
%  Note: This is a tail-recursive predicate.
member_list(Elem, [Elem|_]).
member_list(Elem, [_|Tail]) :- member_list(Elem, Tail).

%% memberchk_list(+Elem, +List)
%  Succeeds if an element is a member of the list, stopping when found.
%  Note: This is a tail-recursive predicate.
memberchk_list(Elem, [Elem|_]) :- !.
memberchk_list(Elem, [_|Tail]) :- memberchk_list(Elem, Tail).

%% nth_list_elem(+Index, +List, ?Elem)
%  Computes the element at the nth index of a 0-based list.
%  Note: This is a tail-recursive predicate.
nth_list_elem(0, [Head|_], Head) :- !.
nth_list_elem(Index, [_|Tail], Elem) :-
    Index > 0,
    NewIndex is Index - 1,
    nth_list_elem(NewIndex, Tail, Elem).

%% list_reversed(+List, ?Reversed)
%  Reverses the elements of a list.
list_reversed([], []).
list_reversed([Head|Tail], Reversed) :-
    list_reversed(Tail, ReversedTail),
    append(ReversedTail, [Head], Reversed).

%% list_reversed_v2(+List, ?Reversed)
%  Reverses the elements of a list efficiently.
list_reversed_v2(List, Reversed) :- reverse_aux(List, [], Reversed).

%  reverse_aux(+List, +Accum, ?Reversed)
%  Note: This is a tail-recursive predicate.
%  TODO Complete this predicate.

%% lists_appended(?List1, ?List2, ?Appended)
%  Appends the second list to the first one.
%  Note: This is a tail-recursive modulo cons predicate.
lists_appended([], List2, List2).
lists_appended([Head1|Tail1], List2,[Head1|AppendedTail]) :-
   lists_appended(Tail1, List2, AppendedTail).
   
%% list_sorted(+List, ?Sorted)
%  Sorts a list in ascending order.
list_sorted([], []).
list_sorted([Head|Tail], Sorted) :-
    partition_list(Tail, Head, Lessers, Greaters),
    list_sorted(Lessers, SortedLessers),
    list_sorted(Greaters, SortedGreaters),
    append(SortedLessers, [Head|SortedGreaters], Sorted).
    
%% partition_list(+Elems, +Pivot, ?Lessers, ?Greaters)
%  Partitions a list of elements into those lesser than the pivot and
%  those greater or equal than the pivot.
%  Note: This is a tail-recursive modulo cons predicate.
partition_list([], _, [], []).
partition_list([Elem|Rest], Pivot, [Elem|Lessers], Greaters) :-
    Elem @< Pivot, !,
    partition_list(Rest, Pivot, Lessers, Greaters).
partition_list([Elem|Rest], Pivot, Lessers, [Elem|Greaters]) :-
    % Elem @>= Pivot
    partition_list(Rest, Pivot, Lessers, Greaters).

%% go
%  Serves as an entry point for the program.
go :-
    Numbers = [30,10,50,40,20],
    list(Numbers),
    write("The list is "), writeln(Numbers),
    
    [Head|Tail] = Numbers,
    write("Its head is "), writeln(Head),
    write("Its tail is "), writeln(Tail),
    
    list_length(Numbers, Length),
    nl, write("Its length is "), writeln(Length),
    
    list_length_v2(Numbers, Length2),
    write("Tail recursion: Its length is "), writeln(Length2),
    
    list_sum(Numbers, Sum),
    nl, write("The sum of its elements is "), writeln(Sum),
    
    list_sum_v2(Numbers, Sum2),
    write("Tail recursion: The sum of its elements is "), writeln(Sum2),
    
    nl, write("Is 40 a member of the list? "),
    (
        memberchk_list(40, Numbers) ->
            writeln("Yes")
            ;
            writeln("No")
    ),
    
    writeln("These are the members of the list, one per line: "),
    forall(
        member_list(Member, Numbers),
        writeln(Member)
    ),
    
    nth_list_elem(3, Numbers, Element),
    write("The element at index #3 is "), writeln(Element),

    list_reversed(Numbers, Reversed),
    nl, write("The reversed list is "), writeln(Reversed),
    
    list_reversed_v2(Numbers, Reversed2),
    write("Tail recursion: The reversed list is "), writeln(Reversed2),
    
    list_sorted(Numbers, Sorted),
    write("The sorted list is "), writeln(Sorted),
    
    Numbers2 = [100|[90, 80, 70]],
    list(Numbers2),
    nl, write("The second list is "), writeln(Numbers2),
    
    lists_appended(Numbers, Numbers2, Appended),
    write("The second list appended to the first is "), 
    writeln(Appended).
    
    %% list_maximum(+NumList, ?Max)
    %  Finds the maximum element in a list of numbers.
    list_maximum([Head|Tail], Max) :- list_maximum(Tail, Head, Max).

    %  list_maximum(+NumList, +CurrentMax, ?Max)
    list_maximum([], Max, Max).
    list_maximum([Head|Tail], CurrentMax, Max) :-
        NewMax is max(Head, CurrentMax),
        list_maximum(Tail, NewMax, Max).

    %% scalar_mult(+Scalar, +NumList, ?MultList)
    %  Multiplies each element in a list of numbers by a scalar.
    scalar_mult(_, [], []).
    scalar_mult(Scalar, [Head|Tail], [MultHead|MultTail]) :-
        MultHead is Scalar * Head,
        scalar_mult(Scalar, Tail, MultTail).

    %% insert(+Elem, +SortedList, ?NewList)
    %  Inserts an element into a sorted list.
    insert(Elem, [], [Elem]).
    insert(Elem, [Head|Tail], [Elem, Head|Tail]) :- Elem =< Head, !.
    insert(Elem, [Head|Tail], [Head|NewTail]) :- insert(Elem, Tail, NewTail).

    %% insertion_sort(+List, ?Sorted)
    %  Sorts a list using insertion sort.
    insertion_sort([], []).
    insertion_sort([Head|Tail], Sorted) :-
        insertion_sort(Tail, SortedTail),
        insert(Head, SortedTail, Sorted).