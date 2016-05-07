// Quando passi un channel come parametro di funzione, puoi indicare
// se il channel deve essere utilizzato solamente per ricevere o per inviare messaggi.
// Questa indicazione aumenta la _type-safety_ dell'intero programma.

package main

import "fmt"

// Questa funzione `ping` accetta un channel per il solo invio.
// Provare a ricevere un valore su questo canale genererà un errore a compile-time
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// La funzione `pong` accetta un channel per la sola ricezione (`pings`) ed uno per
// il solo invio (`pongs`).
func pong(pings <-chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "messaggio passato")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}
