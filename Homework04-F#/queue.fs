(*
 * File: queue.fs
 * Author: Eduardo R Torres 841-##-####
 * Course: COTI 4039-KJ1
 * Date: 11/21/2024
 * Purpose: This source file contains the private implementation of the queue
 *          using a two-list data structure.
 *)

 namespace UprbCollections

 exception EmptyCollection of string
 
 module Queue =
 
     type 'a Queue = Queue of 'a list * 'a list
 
     let empty = Queue ([], [])
 
     let isEmpty = function
         | Queue ([], []) -> true
         | _ -> false
 
     let enqueue elem (Queue (front, back)) = Queue (front, elem::back)
 
     let rec dequeue = function
         | Queue ([], []) -> raise (EmptyCollection "empty queue")
         | Queue ([], back) -> dequeue (Queue (List.rev back, []))
         | Queue (x::rest, back) -> (x, Queue (rest, back))
 
     let peek = function
         | Queue ([], []) -> raise (EmptyCollection "empty queue")
         | Queue ([], back) -> List.head (List.rev back)
         | Queue (x::_, _) -> x
 
     let rec contains elem = function
         | Queue (front, back) -> 
             List.contains elem front || List.contains elem back
 
     let rec elements queue =
         seq {
             match queue with
             | Queue (front, back) ->
                 yield! front
                 yield! List.rev back
         }
 