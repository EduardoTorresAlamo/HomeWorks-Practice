(*
 * File: fibonacci_fns.fsx
 * Author: Eduardo R Torres Alamo
 * Course: COTI 4039-VI1
 * Date: 10/23/2024
 * Purpose: This program computes the nth term of the Fibonacci sequence using different methods.
 * Note: The List module is automatically imported.
 *)

 // Main
 // Function to compute the nth term of the Fibonacci sequence using regular recursion
let rec fibo num =
    if num = 0
    then 0
    elif num = 1
    then 1
    else fibo (num - 1) + fibo (num - 2)


// Function to compute the nth term of the Fibonacci sequence using pattern matching
let rec fiboMatch num = 
    match num with
    | 0 -> 0
    | 1 -> 1
    | _ -> fiboMatch (num - 1) + fiboMatch (num - 2)

// Function to compute the nth term of the Fibonacci sequence using tail-recursive helper function
let fiboIter num =
    let rec fiboIterAux num a b =
        if num = 0
        then a
        else fiboIterAux (num - 1) b (a + b)
    fiboIterAux num 0 1

// Function to compute the nth term of the Fibonacci sequence using List.fold
let fiboHigher num = List.fold (fun (a, b) _ -> (b, a + b)) (0, 1) [1..num] |> fst

// Sequence of the first 200 Fibonacci numbers
let fiboSeq = seq {0..200}
            |> Seq.map (fun num -> fibo num)

// Print
printfn "\nComputing the 7-th term of the Fibonacci sequence:"
printfn "\tUsing if expression: %d" (fibo 7)
printfn "\tUsing pattern matching: %d" (fiboMatch 7)
printfn "\tUsing tail-recursive helper function: %d" (fiboIter 7)
printfn "\tUsing List.fold: %d" (fiboHigher 7)
printfn "The Fibonacci sequence is %A" fiboSeq