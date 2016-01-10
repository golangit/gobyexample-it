// Per default i channels vengono creati senza _buffer_. 
// Scrivere un valore in un channel (`chan <-`) sarà quindi bloccante
// non c'è nessuno in ascolto su quel channel (`<- chan`).
// I channel con _buffer_ accettano un numero limitato di valori e non
// risultano essere bloccanti se nessuno sta ascoltando sul
// channel in questione

package main

import "fmt"

func main() {

    // Qui creiamo un channel (con `make`) di string con un 
    // buffer di dimensione 2.
    messages := make(chan string, 2)

    // Dato che questo channel è bufferizzato, possiamo inviare questi 
    // due valori sul canale senza bloccarsi.
    messages <- "buffered"
    messages <- "channel"

    // Possiamo recuperare successivamente i valori lasciati sul canale.
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}
