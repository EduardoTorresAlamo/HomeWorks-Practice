(*
 * File: odd_count.fsx
 * Author: Eduardo R Torres Alamo
 * Course: COTI 4039-VI1
 * Date: 10/23/2024
 * Purpose: This program computes the number of odd elements in a list of integers using 
 *          different methods.
 * Note: The List module is automatically imported.
 *)

 // Computes the number of odd elements in a list of integers using regular recursion
let rec oddCount lst = 
    match lst with
    | [] -> 0
    | head::tail -> if head % 2 <> 0 then 1 + oddCount tail else oddCount tail

// Computes the number of odd elements in a list of integers using tail-recursive helper function
let oddCountIter lst = 
    let rec oddCountIterAux lst acc = 
        match lst with
        | [] -> acc
        | head::tail -> if head % 2 <> 0 then oddCountIterAux tail (acc + 1) else oddCountIterAux tail acc
    oddCountIterAux lst 0

// Computes the number of odd elements in a list of integers using List.filter and List.length
let countOddHigher lst =
    lst |> List.filter (fun x -> x % 2 <> 0) |> List.length

// Main	
let numbers = [5;10;4;3;2;9;1;6;-7;8]

let resultRec = oddCount numbers

let resultIter = oddCountIter numbers

let result = countOddHigher numbers

// Print
printfn "Counting the number of odd in %A" numbers
printfn "\tUsing regular recursion: %d" resultRec
printfn "\tUsing tail-recursive helper function: %d" resultIter
printfn "\tUsing List.filter and List.length: %d" result