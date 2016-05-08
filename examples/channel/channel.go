// I _channel_ sono i <a href="https://it.wikipedia.org/wiki/Pipe_(informatica)">pipe</a>
// che connettono goroutine concorrenti. Puoi mandare
// e ricevere dei valori da una goroutine a un'altra
// goroutine attraverso un channel.

package main

import "fmt"

func main() {

    // Puoi creare un nuovo channel con la sintassi
    // `make(chan tipo)`. I channel sono dichiarati con i
    // tipi dei valori che possono veicolare.
    messages := make(chan string)

    // Puoi _inviare_ un valore in un channel con la
    // sintassi `channel <-`. Di seguito mandiamo la
    // stringa `"ping"` al channel `messages` che abbiamo
    // dichiarato in precedenza, in una nuova goroutine.
    go func() { messages <- "ping" }()

    // Puoi _ricevere_ un valore da un channel con la
    // sintassi `<-channel`. Di seguito riceveremo il
    // messaggio `"ping"` che abbiamo inviato a messages
    // precedentemente nella goroutine e lo stamperemo.
    msg := <-messages
    fmt.Println(msg)
}
