// Le operazioni di invio e di ricezione nei channel sono
// bloccanti. Tuttavia, possiamo usare un `select` con una
// clausola per implementare invii, ricezioni e select
// a più clausole _non bloccanti_.

package main

import "fmt"

func main() {
    messaggi := make(chan string)
    segnali := make(chan bool)

    // Di seguito una ricezione non bloccante. Se un
    // valore è disponibile su `messages`, allora
    // `select` userà il `case` apposito. Se nessun
    // valore è disponibile, allora selezionerà
    // immediatamente il branch `default`.
    select {
    case msg := <-messaggi:
        fmt.Println("messaggio ricevuto", msg)
    default:
        fmt.Println("nessun messaggio ricevuto")
    }

    // Un invio non bloccante funziona più o meno alla
    // stessa maniera.
    msg := "ciao"
    select {
    case messaggi <- msg:
        fmt.Println("messaggio inviato", msg)
    default:
        fmt.Println("nessun messaggio inviato")
    }

    // Possiamo usare `case` multipli sopra la clausola
    // `default` per implementare un select a più
    // clausole non bloccanti. Qui facciamo una ricezione
    // non-bloccante sia su `messaggi` che su `segnali`.
    select {
    case msg := <-messaggi:
        fmt.Println("messaggio ricevuto", msg)
    case sig := <-segnali:
        fmt.Println("segnale ricevuto", sig)
    default:
        fmt.Println("nessuna attività")
    }
}
