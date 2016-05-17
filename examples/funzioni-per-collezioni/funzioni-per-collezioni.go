// Spesso abbiamo bisogno che i nostri programmi
// effettuino operazioni su collezioni di dati,
// come selezionare tutti gli elementi che
// soddisfino un determinato condizionale, oppure mappare
// gli elementi su una nuova collezione con un criterio
// specifico e una funzione ad hoc.

// In alcuni linguaggi è idiomatico l'utilizzo di
// strutture di dati e algoritmi [generici](http://en.wikipedia.org/wiki/Generic_programming).
// Go non supporta algoritmi/strutture generiche: in
// genere si forniscono funzioni per collezioni se sono
// specificamente necessari per il tuo programma e per i
// tuoi tipi.

// A seguire un po' di esempi di funzioni che agiscono su
// slice di `string`. Puoi usare questi esempi per
// costruire le tue proprie funzioni. Nota che in alcuni
// casi può essere più chiaro scrivere il corpo della
// funzione direttamente nel tuo codice stesso, piuttosto
// che creare e chiamare una funzione helper.

package main

import "strings"
import "fmt"

// Restituisce il primo indice del valore in `vs`
// corrispondente a `t`, o -1 se nessuna corrispondenza
// viene trovata.
func Indice(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

// Restituisce `true` se la stringa `t` è nello slice `vs`.
func Incluso(vs []string, t string) bool {
    return Indice(vs, t) >= 0
}

// Restituisce `true` se una delle stringhe nello slice
// soddisfa la funzione condizionale `f`.
func AlmenoUno(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

// Restituisce `true` se tutte le `string` nello slice
// soddisfano la funzione condizionale `f`.
func Tutti(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

// Restituisce un nuovo slice che contiene tutte le stringhe
// nello slice `vs` che soddisfano la funzione
// condizionale `f`.
func Filtra(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

// Restituisce un nuovo slice che contiene i risultati
// dell'applicazione della funzione `f` su ogni stringa
// dello slice `vs`.
func Applica(vs []string,
    f func(string) string) []string {
    vsa := make([]string, len(vs))
    for i, v := range vs {
        vsa[i] = f(v)
    }
    return vsa
}

func main() {

    // Di seguito proviamo ad utilizzare le nostre varie
    // funzioni di collezione.
    var strs = []string{"pesca", "mela", "pera", "prugna"}

    fmt.Println(Indice(strs, "pera"))

    fmt.Println(Incluso(strs, "uva"))

    fmt.Println(AlmenoUno(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Tutti(strs, func(v string) bool {
        return strings.HasPrefix(v, "p")
    }))

    fmt.Println(Filtra(strs, func(v string) bool {
        return strings.Contains(v, "r")
    }))

    // Gli esempi sovrastanti usavano tutti funzioni
    // anonime, ma puoi anche usare funzioni pre-esistenti
    // (non anonime) con il tipo corretto.
    fmt.Println(Applica(strs, strings.ToUpper))

}
