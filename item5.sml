(*
Generador de números pseudoAleatorio

Creadores
    Sanabria Solano María Fernanda - 2021005572 
    Obando Arrieta Felipe de Jesús - 2021035489
*)
(*
For loop
Fuente: Fluet, M. (2021). For Loop [Software]. En MLton. http://mlton.org/ForLoops
*)
datatype for = to of int * int
             | downto of int * int
;

infix to downto;

val for =
    fn lo to up =>
       (fn f => let fun loop lo = if lo > up then ()
                                  else (f lo; loop (lo+1))
                in loop lo end)
     | up downto lo =>
       (fn f => let fun loop up = if up < lo then ()
                                  else (f up; loop (up-1))
                in loop up end)
;
(* =================================================== *)

(*
    Determina si x es o no primo
    Parametros:
        -x (int): Numero a determinar.
    Retorna:
        boolean: True si es primo, False si no.
    *)
fun isPrime 0 = false (*No es numero primo*)
    | isPrime 1 = false (*No es numero primo*)
    | isPrime 2 = true (*Primer numero primo*)
    | isPrime x =
        let
        val j = ref 2
        val result = ref true
        in
            while (!j) <> (x div (2+1)) do ( (*Se recorren posibles divisores de x*)
                
                if (x mod !j) = 0 then result := false (*Se encuentra un divisor*)
                else (); 
                j := !j + 1
            );
            !result (*Se devuelve el resultado*)
        end
;

fun isPrime 0 = false (*No es numero primo*)
    | isPrime 1 = false (*No es numero primo*)
    | isPrime 2 = true (*Primer numero primo*)
    | isPrime x =
        let
        val result = ref true
        in
            for (2 to (x div (2+1))) (fn j => (*Se recorren posibles divisores de x*)
                if (x mod j) = 0 then result := false (*Se encuentra un divisor*)
                else ()
            );

            !result (*Se devuelve el resultado*)
        end
;

fun generaAleatorio (semilla,n) = 
    if semilla < 11 then [] else
    if semilla > 257 then [] else
    if n < 500 then [] else
    if n > 5000 then [] else
    let 
        fun getX sem = (* Funcion para optener la x *)
            let
                val repeat = ref true
                val x = ref 0
                val candidato = ref 11

            in
                (*Buscamos al candidato de x*)
                while !repeat do ( (*Se recorren posibles candidatos*)
                    
                    if (isPrime (!candidato)) andalso !candidato >= sem then (
                        x := !candidato; (*Se encontro un candidatos*)
                        repeat := false)
                    else (); 
                    candidato := !candidato + 1
                );
                !x
            end
        val x = ref (getX semilla)
        val lista = ref []
        val m = 4096
        val a = 109
        val b = 853
    in
        for (0 to n-1)(fn j =>
            (lista := !lista @ [!x];
            x := ((a * (!x) + b) mod m))
        );
        !lista
    end;

fun convertidor [] = []
    | convertidor(x::[]) = [x mod 255]
    | convertidor(x::xs) = (x mod 255::convertidor xs)
;    


fun generaAleatorioAux (semilla,n) =
    let
        val list = convertidor(generaAleatorio (semilla,n))
        in
            if list = [] then print("parametros invalidos") else ()
            list



