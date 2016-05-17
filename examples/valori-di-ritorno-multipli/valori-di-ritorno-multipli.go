// Go supporta _valori restituiti multipli_, similmente
// a python. Questa funzionalità è usata spesso nel Go
// idiomatico, per esempio per restituire sia il valore
// sia l'eventuale errore nell'esecuzione di una funzione.

package main

import "fmt"

// L'indicazione `(int, int)` in questa funzione ci dice
// che la funzione restituisce due `int`.
func vals() (int, int) {
    return 3, 7
}

func main() {

    // Di seguito utilizzando il _multiple assignment_
    // creiamo due diverse variabili dai valori restituiti
    // della funzione `vals()`.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // Se vuoi soltanto avere una parte dei valori
    // restituiti, usa il blank identifier `_`.
    _, c := vals()
    fmt.Println(c)
}

// todo: named return parameters
// todo: naked returns
