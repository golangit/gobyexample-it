// Talvolta è necessario ordinare una collezione seguendo
// un ordinamento che è differente dal suo ordinamento
// naturale. Ad esempio, supponiamo di dover ordinare
// un insieme di stringhe in base alla loro lunghezza
// (invece che in ordine alfabetico). Ecco un esempio di
// ordinamento personalizzato in Go.

package main

import "sort"
import "fmt"

// Al fine di poter ordinare tramite una funzione personalizzata
// in Go, abbiamo bisogno di un tipo corrispondente. Abbiamo
// creato il tipo `ByLength` che altro non è che un alias per
// il tipo built-in `[]string`.
type ByLength []string

// Implementiamo l'interfaccia `sort.Interface` implementando
// le funzioni `Len`, `Less` e `Swap` al fine di poter usare
// la funzione `Sort` generica dal package `sort`.
// `Len` e `Swap` saranno generalmente simili indipendentemente
// dal tipo, mentre `Less` conterrà la logica per l'ordinamento
// personalizzato. Nel nostro caso, dato che vogliamo ordinare
// in baso alla lunghezza, utilizzeremo `len(s[i])` e
// `len(s[j])` per effettuare il confronto.
func (s ByLength) Len() int {
    return len(s)
}
func (s ByLength) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
    return len(s[i]) < len(s[j])
}

// Dopo aver dichiarato queste funzioni, è possibile
// effettuare un ordinamento personalizzato tramite un cast
// dello slice `fruits` al tipo `ByLength` sul quale chiameremo
// la funzione `sort.Sort`.
func main() {
    fruits := []string{"pesca", "banana", "kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
