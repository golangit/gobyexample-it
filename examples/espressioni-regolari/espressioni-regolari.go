// Go offre il supporto nativo alle [espressioni regolari](https://it.wikipedia.org/wiki/Espressione_regolare).
// Ecco qualche esempio di operazioni comuni con le regexp in Go.

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

    // Questo controlla che una stringa soddisfi una regexp
    match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
    fmt.Println(match)

    // Prima abbiamo utilizzato le regexp in modo diretto.
    // In generale si invoca prima il `Compile` su una
    // regexp, al fine di ottenere una struct `Regexp`
    // ottimizzata.
    r, _ := regexp.Compile("p([a-z]+)ch")

    // Su questa struttura abbiamo a disposizione svariati
    // metodi. Ecco un test simile al nostro primo esempio.
    fmt.Println(r.MatchString("peach"))

    // Questo restituisce il primo match (una stringa) che
    // soddisfa la regexp.
    fmt.Println(r.FindString("peach punch"))

    // Anche questo metodo restituisce il primo match,
    // ma invece di restituirne la stringa, restituisce l'indice
    // di inizio e di fine della stringa.
    fmt.Println(r.FindStringIndex("peach punch"))

    // La variante `Submatch` include informazioni
    // relativi al match del pattern e anche ai match di
    // eventuali sotto-pattern. In questo caso restituirebbe
    // i match per i pattern `p([a-z]+)ch` e `([a-z]+)`.
    fmt.Println(r.FindStringSubmatch("peach punch"))

    // In modo simile questo metodo restituirà gli indici
    // di inizio e di fine del match e dei sotto-match
    fmt.Println(r.FindStringSubmatchIndex("peach punch"))

    // La variante `All`, invece di restituire solo la prima
    // occorrenza del match, restituisce tutti i match.
    fmt.Println(r.FindAllString("peach punch pinch", -1))

    // Questa variante `All` è disponibile per tutte le
    // funzioni che abbiamo visto prima.
    fmt.Println(r.FindAllStringSubmatchIndex(
        "peach punch pinch", -1))

    // Se si passa un intero non negativo come secondo parametro
    // di queste funzioni si sta limitando il numero
    // di match da restituire. Ad esempio in questo caso
    // si è interessati solamente ai primi 2 match.
    fmt.Println(r.FindAllString("peach punch pinch", 2))

    // Il nostro primo esempio mostrava la funzione
    // `MatchString` che lavorava su stringhe. È possibile
    // passare anche parametro di tipo `[]byte` ed
    // utilizzare la funzione `Match`.
    fmt.Println(r.Match([]byte("peach")))

    // Se si vuole creare una costante a partire da
    // una regexp si può utilizzare la funzione `MustCompile`
    // (invece di `Compile`). La classica `Compile` non
    // funzionerà in quanto ha 2 valori restituiti.
    r = regexp.MustCompile("p([a-z]+)ch")
    fmt.Println(r)

    // Il pacchetto `regexp` può essere utilizzato anche per
    // effettuare sostituzioni di stringhe/sottostringhe.
    fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

    // La variante `Func` permette di trasformare i match
    // trovati passandogli una funzione che verrà applicata
    // su ogni occorrenza
    in := []byte("a peach")
    out := r.ReplaceAllFunc(in, bytes.ToUpper)
    fmt.Println(string(out))
}
