% arithmethic.pl

% gcd using cases
gcd(0, B, B) :- B > 0.
gcd(A, 0, A) :- A > 0.
gcd(A, B, GCD) :-
    A >= B,
    A1 is A mod B,
    gcd(A1, B, GCD).
gcd(A, B, GCD) :-
    A < B,
    gcd(B, A, GCD).

% gcd using if-then-else
gcd_if(A, B, GCD) :-
    ( A =:= 0 -> GCD = B
    ; B =:= 0 -> GCD = A
    ; A >= B -> A1 is A mod B, gcd_if(A1, B, GCD)
    ; gcd_if(B, A, GCD)
    ).

% fibo using regular recursion
fibo(0, 0).
fibo(1, 1).
fibo(N, Fibo) :-
    N > 1,
    N1 is N - 1,
    N2 is N - 2,
    fibo(N1, Fibo1),
    fibo(N2, Fibo2),
    Fibo is Fibo1 + Fibo2.

% fibo using tail-recursive helper predicate
fibo_tail(N, Fibo) :-
    fibo_helper(N, 0, 1, Fibo).

fibo_helper(0, A, _, A).
fibo_helper(N, A, B, Fibo) :-
    N > 0,
    N1 is N - 1,
    Sum is A + B,
    fibo_helper(N1, B, Sum, Fibo).