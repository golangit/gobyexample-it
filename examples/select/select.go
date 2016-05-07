// _Select_ ti permette di elaborare valori ricevuti
// da più canali con una sola goroutine. Combinare le
// goroutine e i canali con select è uno dei punti di
// forza di Go.

package main

import "time"
import "fmt"

func main() {

    // Per il nostro esempio utilizzeremo select su due
    // canali.
    c1 := make(chan string)
    c2 := make(chan string)

    // Ogni canale riceverà un valore dopo un po' di
    // tempo, giusto per simulare ad esempio operazioni
    // RPC bloccanti eseguite su multiple goroutine.
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "uno"
    }()
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "due"
    }()

    // Utilizzeremo `select` per ricevere dai due canali
    // simultaneamente, stampando ogni valore che arriva.
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("ricevuto", msg1)
        case msg2 := <-c2:
            fmt.Println("ricevuto", msg2)
        }
    }
}
