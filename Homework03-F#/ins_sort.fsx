(*
 * File: ins_sort.fsx
 * Author: Eduardo R Torres Alamo
 * Course: COTI 4039-VI1
 * Date: 10/23/2024
 * Purpose: This program implements a recursive insertion sort of a list.
 * Note: The List module is automatically imported.
 *)

 // Functions
 // Inserts a element into a sorted list
 let rec insert elem lst = 
    match lst with
    | [] -> [elem]
    | head::tail -> if elem <= head then elem::lst else head::(insert elem tail)

// Sorts a list using insertion sort
let rec insertionSort lst = 
    match lst with
    | [] -> []
    | head::tail -> insert head (insertionSort tail)
