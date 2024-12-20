(*
 * File: scalar_mult.fsx
 * Author: Eduardo R Torres Alamo
 * Course: COTI 4039-VI1
 * Date: 10/23/2024
 * Purpose: This program determines the scalar multiplication of an integer and a 
 *          list of integers using different methods.
 * Note: The List module is automatically imported.
 *)

 // Function to perform scalar multiplication using regular recursion
let rec scalarMult num lst = 
    match lst with
    | [] -> []
    | head::tail -> (head * num) :: scalarMult num tail


// Function to perform scalar multiplication using tail-recursive helper function
let scalarMultIter num lst = 
    let rec scalarMultIterAux num lst acc = 
        match lst with
        | [] -> acc
        | head::tail -> scalarMultIterAux num tail (acc @ [head * num])
    scalarMultIterAux num lst []

// Function to perform scalar multiplication using List.map
let scalarMultHigher scalar lst =
    lst |> List.map (fun x -> x * scalar)



// Main
let scalar = 3
let numbers = [2; 7; 4]

let resultRec = scalarMult scalar numbers

let resultMultIter = scalarMultIter scalar numbers

let result = scalarMultHigher scalar numbers


// Print the result
printfn "Computing the scalar multtiplication of %d by %A" scalar numbers
printfn "\tUsing regular recursion: %A" resultRec
printfn "\tUsing tail-recursive helper function: %A" resultMultIter
printfn "\tUsing List.map: %A" result