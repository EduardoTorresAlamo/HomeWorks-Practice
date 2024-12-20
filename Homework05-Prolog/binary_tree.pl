/*
 * File: binary_tree.pl
 * Author: Antonio F. Huertas 
 * Course: COTI 4039-KJ1
 * Date: 12/12/2024
 * Purpose: This program defines and manipulates a binary search tree.
 */

%% binary_tree(?Tree)
%  Succeeds if the argument is a tree.
binary_tree(nil).
binary_tree(node(_, Left, Right)) :-
    binary_tree(Left),
    binary_tree(Right).

%% member_tree(?Elem, ?Tree)
%  Enumerates all members of the tree in order or determines if an element
%  is a member of the tree.
member_tree(Elem, node(_, Left, _)) :- member_tree(Elem, Left).
member_tree(Elem, node(Elem, _, _)).
member_tree(Elem, node(_, _, Right)) :- member_tree(Elem, Right).

%% memberchk_tree(?Elem, +Tree)
%  Succeeds if the element is a member of the tree, stopping when found.
memberchk_tree(Elem, node(Elem, _, _)) :- !.
memberchk_tree(Elem, node(Root, Left, _)) :-
    Elem @< Root, !,
    memberchk_tree(Elem, Left).    
memberchk_tree(Elem, node(_Root, _, Right)) :-
    % Elem @> _Root
    memberchk_tree(Elem, Right).
    
%% insert(+Elem, +Tree0, ?Tree)
%  Succeeds if an element can be inserted in a tree.
insert(Elem, nil, node(Elem, nil, nil)) :- !.
insert(Elem, node(Elem, Left, Right), node(Elem, Left, Right)) :- !.
insert(Elem, node(Root, Left, Right), node(Root, NewLeft, Right)) :-
    Elem @< Root, !,
    insert(Elem, Left, NewLeft).
insert(Elem, node(Root, Left, Right), node(Root, Left, NewRight)) :-
    % Elem @> Root
    insert(Elem, Right, NewRight).

%% tree_size(+Tree, ?Size)
%  Computes the number of elements in a tree.
tree_size(nil, 0).
tree_size(node(_, Left, Right), Size) :-
    tree_size(Left, SizeLeft),
    tree_size(Right, SizeRight),
    Size is 1 + SizeLeft + SizeRight.

%% tree_sum(+NumTree, ?Sum)
%  Computes the sum of the elements in a tree of numbers.
tree_sum(nil, 0).
tree_sum(node(Root, Left, Right), Sum) :-
    tree_sum(Left, SumLeft),
    tree_sum(Right, SumRight),
    Sum is Root + SumLeft + SumRight.

%% write_tree(+Tree)
%  Displays the elements of the tree in order.
write_tree(nil).
write_tree(node(Root, Left, Right)) :-
    write_tree(Left),
    write(Root), write(" "),
    write_tree(Right).

%% tree_of_nums(?Tree)
%  Declares a binary search tree of numbers.
tree_of_nums(Tree) :-
    insert(30, nil, Tree0),
    insert(10, Tree0, Tree1),
    insert(50, Tree1, Tree2),
    insert(40, Tree2, Tree3),
    insert(20, Tree3, Tree).


%% list_to_tree(+List, ?Tree)
%  Succeeds if Tree is the binary search tree corresponding to List.
list_to_tree(List, Tree) :-
    foldl(insert, List, nil, Tree).

%% tree_height(+Tree, ?Height)
%  Succeeds if Height is the number of levels in Tree.
tree_height(nil, 0).
tree_height(node(_, Left, Right), Height) :-
    tree_height(Left, HeightLeft),
    tree_height(Right, HeightRight),
    Height is 1 + max(HeightLeft, HeightRight).

%% tree_inorder(+Tree, ?List)
%  Succeeds if List is the inorder traversal of Tree.
tree_inorder(nil, []).
tree_inorder(node(Root, Left, Right), List) :-
    tree_inorder(Left, ListLeft),
    tree_inorder(Right, ListRight),
    append(ListLeft, [Root|ListRight], List).
    
%% go
%  Serves as an entry point for the program.
go :-
    tree_of_nums(Tree),
    (
        binary_tree(Tree) ->
            write("The tree is "), writeln(Tree)
            ;
            writeln("Error! This is not a tree"), fail
    ),
    write("The elements of the tree are "), write_tree(Tree), nl,
    
    write("Is 40 a member of the tree? "),
    (
        memberchk_tree(40, Tree) ->
            writeln("Yes")
            ;
            writeln("No")
    ),
    
    node(Root, Left, Right) = Tree,
    nl,write("Its root is "), writeln(Root),
    write("Its left subtree is "), writeln(Left),
    write("Its right subtree is "), writeln(Right),

    tree_size(Tree, Size),
    nl, write("Its size is "), writeln(Size),
    
    tree_sum(Tree, Sum),
    write("The sum of its elements is "), writeln(Sum),

    nl, writeln("These are the members of the tree, one per line: "),
    forall(
        member_tree(Member, Tree),
        writeln(Member)
    ).

