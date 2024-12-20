(* Cierto o falso 
10 preguntas
Fold 
Map
Filter
Sequencial

5 programas

1 Insert un elemneto en una lista ordenada sin desordenar la lista
4->  1235
12345

2 Count busca un elemento en una lista y devuelve el indice

3 Math Pow
Lista devolver pow de todos los elementos

4 Suma de elementos de una lista recursiva y tail recursive

5  Filtrar una lista para que sean soo los elementos pares y sumarlos
        Hacerlo con Pipes y recursion normal

6 Seq de power y usar una funcion de otro ejercicio
*)

// 1 Insert un elemneto en una lista ordenada sin desordenar la lista
let rec insertOrdered x lst =
    match lst with
    | [] -> [x]
    | head :: tail when x <= head -> x :: lst
    | head :: tail -> head :: (insertOrdered x tail)

// Ejemplo de uso
let listaOrdenada = [1; 2; 5; 3]
let nuevaLista = insertOrdered 4 listaOrdenada
printfn "%A" nuevaLista  // Output: [1; 2; 3; 4; 5]



// 2 Count busca un elemento en una lista y devuelve el indice
let rec countElement x lst =
    let rec countElementAux x lst index =
        match lst with
        | [] -> -1
        | head :: tail when head = x -> index
        | head :: tail -> countElementAux x tail (index + 1)
    countElementAux x lst 0

// Ejemplo de uso
let listacount = [1; 2; 3; 4; 5]
let indice = countElement 4 listacount
printfn "%d" indice  // Output: 3


// 3 Math Pow
let powListMP lst =
    lst |> List.map (fun x -> x * x)

// Recursion
let rec powList lst =
    match lst with
    | [] -> []
    | head :: tail -> (head * head) :: (powList tail)

// Tail-recursion
let powListTailRec lst =
    let rec aux lst acc =
        match lst with
        | [] -> List.rev acc
        | head :: tail -> aux tail ((head * head) :: acc)
    aux lst []

// Ejemplo de uso
let lista = [1; 2; 3; 4; 5]
let listaCuadrada = powList lista
printfn "%A" listaCuadrada  // Output: [1; 4; 9; 16; 25]

let listaCuadradaTailRec = powListTailRec lista
printfn "%A" listaCuadradaTailRec  // Output: [1; 4; 9; 16; 25]

// Ejemplo de uso
let listaMP = [1; 2; 3; 4; 5]
let listaCuadradaMP = powList listaMP
printfn "%A" listaCuadradaMP  // Output: [1; 4; 9; 16; 25]



// 4 Suma de elementos de una lista recursiva y tail recursive
let rec sumList lst =
    match lst with
    | [] -> 0
    | head :: tail -> head + sumList tail


// 5 Filtrar una lista para que sean solo los elementos pares y sumarlos
let rec sumEven lst =
    match lst with
    | [] -> 0
    | head :: tail when head % 2 = 0 -> head + sumEven tail
    | head :: tail -> sumEven tail

// tail recursive
let sumEvenTailRec lst =
    let rec aux lst acc =
        match lst with
        | [] -> acc
        | head :: tail when head % 2 = 0 -> aux tail (acc + head)
        | head :: tail -> aux tail acc
    aux lst 0

let sumlistTailRec lst =
    let rec aux lst acc =
        match lst with
        | [] -> acc
        | head :: tail -> aux tail (acc + head)
    aux lst 0
    

