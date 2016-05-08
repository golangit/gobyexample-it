// Talvolta è necessario che i nostri programmi Go
// gesticano in modo corretto i <a href="https://it.wikipedia.org/wiki/Segnale_(informatica)">segnali Unix</a>
// Ad esempio, potremmo volere che il nostro programma
// termini in modo corretto se riceve un `SIGTERM` oppure
// vorremmo dire ad un programma a linea di comando di
// interrompersi se riceve un `SIGINT`.
// In questo esempio vedremo come è semplice gestire i segnali
// con i channel di Go.

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

    // I segnali su Go vengono gestiti tramite l'invio
    // di valori di tipo `os.Signal` su un channel.
    // Qui creiamo un channel per ricevere queste notifiche,
    // ed uno che utilizzeremo per la terminazione del programma.
    sigs := make(chan os.Signal, 1)
    done := make(chan bool, 1)

    // La funzione `signal.Notify` registra questo channel
    // per la ricezione di specifici segnali.
    signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

    // Questa goroutine si blocca sul channel `sigs` in
    // attesa di un segnale. Quando lo riceve, lo stamperà
    // a schermo e aggiungerà un messaggio nel channel
    // `done`, per notificare la goroutine main che può
    // terminare.
    go func() {
        sig := <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    // Il programma si bloccherà qui in attesa di un valore
    // sul channel `done` e poi terminerà.
    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
