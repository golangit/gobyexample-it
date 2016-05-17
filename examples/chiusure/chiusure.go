// Go supporta le [_funzioni anonime_](https://it.wikipedia.org/wiki/Funzione_anonima),
// che possono formare delle <a href="https://it.wikipedia.org/wiki/Chiusura_(informatica)"><em>chiusure</em></a>.
// Le funzioni anonime sono utili quando vuoi definire una
// funzione senza darle un nome.

package main

import "fmt"

// Questa funzione `intSeq` restituisce un'altra funzione, che
// definiamo anonimamente dentro il corpo della funzione
// `intSeq`. La funzione restituita _racchiude_ la
// variabile `i` per formare una chiusura.
func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}

func main() {

    // Facciamo una chiamata ad `intSeq`, assegnando il
    // risultato (una funzione) a `nextInt`. Il valore di
    // questa funzione racchiude in sé stessa il valore di
    // `i`, il quale verrà aggiornato la prossima volta che
    // utilizziamo `nextInt`.
    nextInt := intSeq()

    // Osserviamo l'effetto della chiusura facendo una
    // chiamata a `nextInt` un po' di volte.
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // Per confermare che lo stato è unico a quella
    // funzione particolare, creiamo e testiamone una
    // nuova.
    newInts := intSeq()
    fmt.Println(newInts())
}
