// Go supporta _multipli valori di ritorno_, similmente
// a python. Questa funzionalità è usata spesso nel Go
// idiomatico, per esempio per ritornare sia il valore
// sia l'eventuale errore nell'esecuzione di una funzione.

package main

import "fmt"

// L'`(int, int)` in questa funzione ci dice che la
// funzione ritorna due `int`.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// Di seguito utilizzando il _multiple assignment_
	// creiamo due diverse variabili dai valori di ritorno
	// della funzione `vals()`.
	// Here we use the 2 different return values from the
	// call with _multiple assignment_.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// Se vuoi soltanto avere una parte dei valori
	// ritornati, usa il blank identifier `_`.
	_, c := vals()
	fmt.Println(c)
}

// todo: named return parameters
// todo: naked returns
