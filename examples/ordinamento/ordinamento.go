// Il package `sort` di Go permette di ordinare tipi
// sia built-in, sia user-defined. Daremo per prima cosa
// un'occhiata al sorting per i built-in.

package main

import "fmt"
import "sort"

func main() {

    // I metodi di sort sono specifici per il tipo
    // built-in; qui vediamo un esempio per le string.
    // Nota che il sorting non richiede né l'uso dei
    // pointer, né restituisce un nuovo slice, ma modifica
    // lo slice che viene passato ordinando
    // direttamente elementi gli elementi al suo
    // interno.
    strs := []string{"c", "a", "b"}
    sort.Strings(strs)
    fmt.Println("Strings:", strs)

    // Un esempio del sorting di `int`.
    ints := []int{7, 2, 4}
    sort.Ints(ints)
    fmt.Println("Ints:   ", ints)

    // Possiamo anche usare `sort` per determinare
    // se uno slice è già stato ordinato.
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ", s)
}
