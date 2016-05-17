// Effettuare il parsing da una stringa ad un numero
// è un task tanto semplice quanto comune in molti programmi.
// Vedremo adesso come effettuare il parsing dei numeri in Go.

package main

// Il pacchetto `strconv` della standard library offre
// le funzioni necessarie al parsing
import "strconv"
import "fmt"

func main() {

    // Tramite la funzione `ParseFloat` è possibile indicare il
    // numero di bit della precisione da utilizzare (in questo
    // caso 64).
    f, _ := strconv.ParseFloat("1.234", 64)
    fmt.Println(f)

    // Tramite la funzione `ParseInt` è possibile indicare la
    // base da utilizzare per il parsing, tramite il secondo
    // parametro. Il valore `0` sta ad indicare che la base
    // deve essere inferita automaticamente dalla stringa.
    // Il parametro `64` richiede che il risultato sia
    // rappresentabile con 64 bit.
    i, _ := strconv.ParseInt("123", 0, 64)
    fmt.Println(i)

    // La funzione `ParseInt` riconosce anche i numeri esadecimali
    d, _ := strconv.ParseInt("0x1c8", 0, 64)
    fmt.Println(d)

    // È disponibile anche la funzione `ParseUint`
    u, _ := strconv.ParseUint("789", 0, 64)
    fmt.Println(u)

    // La funzione `Atoi` è una funzione di convenienza per
    // convertire da numeri in base 10 a `int`
    k, _ := strconv.Atoi("135")
    fmt.Println(k)

    // Le funzioni di parsing possono restituire un errore
    _, e := strconv.Atoi("wat")
    fmt.Println(e)
}
