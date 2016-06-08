// I _Timeout_ sono fondamentali per i programmi che si
// connettono a risorse esterne o che hanno comunque
// bisogno di limitare il tempo d'esecuzione.
// Implementare i timeout in Go è semplice grazie ai
// channel e al costrutto `select`.

package main

import "time"
import "fmt"

func main() {

    // Ai fini del nostro esempio, supponiamo di star
    // eseguendo una chiamata esterna che restituisce il suo
    // risultato sul channel `c1` dopo 2 secondi.
    c1 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c1 <- "risultato 1"
    }()

    // Qui mostriamo il `select` che implementa un timeout.
    // Il comando `res := <-c1` attende il risultato della funzione
    // mentre il comando `<-Time.After` emette un risultato dopo
    // un timeout di 1 secondo.
    // Dal momento che il comando `select` procede con il primo
    // canale per cui è disponibile un valore, eseguiremo il caso
    // del timout se la nostra chiamata esterna richiede più di
    // 1 secondo.
    select {
    case res := <-c1:
        fmt.Println(res)
    case <-time.After(time.Second * 1):
        fmt.Println("timeout 1")
    }

    // In questo caso possiamo vedere come impostare un timeout di
    // 3 secondi sia sufficiente a far restituire la nostra goroutine
    // che scrive sul canale `c2`. Riusciremo infatti a vedere il
    // risultato a schermo e a non far scattare il timeout.
    c2 := make(chan string, 1)
    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "risultato 2"
    }()
    select {
    case res := <-c2:
        fmt.Println(res)
    case <-time.After(time.Second * 3):
        fmt.Println("timeout 2")
    }
}
