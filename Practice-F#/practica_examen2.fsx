let sum lst =
    let rec aux lst acc =
        match lst with
        | [] -> acc
        | hd::tl -> aux tl (acc + hd)

    aux lst 0

let rec sumRec lst =
    match lst with
    |[] -> 0 
    |hd::tl -> hd + sumRec tl

let rec insert x lst = 
    match lst with
    | [] -> [x]
    | hd::tl when x <= hd -> x::lst
    | hd::tl -> hd::(insert x tl)

// elevate a list to a power given a number	
let rec powList lst n =
    match lst with
    | [] -> []
    | hd::tl -> (hd ** n) :: powList tl n

let  powListTailRec lst n =
    let rec aux lst acc =
        match lst with
        | [] -> List.rev acc
        | hd::tl -> aux tl ((hd ** n) :: acc)
    aux lst []

let powSeq = seq {0..200} 
            |> Seq.map (fun x -> float x ** 2.0)

printfn "%A" (Seq.toList powSeq)

let sumOfSquares lst =
    lst
    |> List.map (fun x -> x * x)
    |> List.sum

printfn "%d" (sumOfSquares [1; 2; 3; 4; 5])

let sumOfSquares lst =
    lst
    |> List.map (fun x -> x * x)
    |> List.sum

printfn "%d" (sumOfSquares [1; 2; 3; 4; 5])

let sum' lst =
    let rec sumOfSquaresRec' lst =
        match lst with
        | [] -> 0
        | hd::tl -> (hd * hd) + sumOfSquaresRec' tl

    printfn "%d" (sumOfSquaresRec' [1; 2; 3; 4; 5])

    let sumOfSquaresTailRec'' lst =
        let rec aux lst acc =
            match lst with
            | [] -> acc
            | hd::tl -> aux tl (acc + (hd * hd))
        aux lst 0

    printfn "%d" (sumOfSquaresTailRec'' [1; 2; 3; 4; 5])

    let rec fibonacci n =
        match n with
        | 0 -> 0
        | 1 -> 1
        | _ -> fibonacci (n - 1) + fibonacci (n - 2)

    printfn "%d" (fibonacci 10)

let fibonacciTailRec n =
    let rec aux n acc1 acc2 =
        match n with
        | 0 -> acc1
        | _ -> aux (n - 1) acc2 (acc1 + acc2)
    aux n 0 1

