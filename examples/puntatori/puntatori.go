// Go permette l'utilizzo dei <a href="https://it.wikipedia.org/wiki/Puntatore_(programmazione)"><em>puntatori</em></a>,
// che si traduce nell'abilità di passare riferimenti a
// valori all'interno del programma.

package main

import "fmt"

// Dimostreremo come i puntatori funzionino diversamente
// dai valori tramite 2 funzioni: `zeroval` e `zeroptr`.
// `zeroval` ha un parametro di tipo `int`, quindi il
// parametro passato sarà un valore, non un puntatore.
// Quando chiameremo la funzione `zeroval`, il suo
// parametro `ival` verrà copiato da quello della
// funzione chiamante.
func zeroval(ival int) {
    ival = 0
}

// `zeroptr` invece ha un parametro di tipo `*int`, e ciò
// significa che è un puntatore a un `int`. L'istruzione
// `*iptr` nel corpo della funzione permette di
// _dereferenziare_ l'indirizzo di memoria puntato da `iptr` in
// modo da otternere il suo valore. Se si assegna un valore
// ad `*iptr` si va a modificare il valore all'indirizzo di memoria
// puntanto.
func zeroptr(iptr *int) {
    *iptr = 0
}

func main() {
    i := 1
    fmt.Println("iniziale: ", i)

    zeroval(i)
    fmt.Println("zeroval:  ", i)

    // La formula `&i` restituisce l'indirizzo nella memoria
    // di `i`, ovvero un puntatore ad `i`.
    zeroptr(&i)
    fmt.Println("zeroptr:  ", i)

    // Anche i puntatori possono essere stampati.
    fmt.Println("puntatore:", &i)
}
