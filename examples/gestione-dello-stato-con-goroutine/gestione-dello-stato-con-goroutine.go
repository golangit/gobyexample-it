// Nell'esempio precedente abbiamo usato il lock esplicito
// insieme alle mutex per sincronizzare l'accesso allo
// stato condiviso da più goroutine. Un'altra opzione è
// quella di utilizzare le funzionalità di sincronizzazione
// native di Go, sfruttando goroutine e channel per ottenere
// lo stesso risultato. Questo approcchio channel-based
// risulta più in linea con il principio di Go di avere la
// memoria condivisa tramite la comunicazione, ed ogni
// goroutine con la propria memoria privata.

package main

import (
    "fmt"
    "math/rand"
    "sync/atomic"
    "time"
)

// In questo esempio lo stato verrà gestito da una singola
// goroutine. Questo ci garantirà che i dati non verranno
// mai corrotti da accessi concorrenti. Per poter leggere
// ed aggiornare lo stato le altre goroutine dovranno inviare
// dei messaggi alla goroutine proprietaria che risponderà
// con un messaggio contenente il dato richiesto.
// Le `struct` `readOp` e `writeOp` incapsulano queste
// richieste e contengono il field `resp` per indicare
// la risposta della goroutine proprietaria.
type readOp struct {
    key  int
    resp chan int
}
type writeOp struct {
    key  int
    val  int
    resp chan bool
}

func main() {

    // Come prima terremo conto del numero di operazioni
    // che andremo ad eseguire
    var ops int64 = 0

    // I channel `reads` e `writes` saranno utilizzati
    // dalle altre goroutine per effettuare richieste
    // di lettura e di scrittura dello stato.
    reads := make(chan *readOp)
    writes := make(chan *writeOp)

    // Questa sarà la goroutine che gestirà lo stato,
    // rappresentato da una map come nell'esempio precedente,
    // ma in questo caso visibile solo alla goroutine.
    // Questa goroutine effettuerà delle `select` sui due
    // channel e risponderà alle richieste che arrivano.
    // La goroutine si occuperà di gestire la richiesta e
    // restituirà un valore sul channel `resp` per indicare
    // un successo (nel caso di una richiesta `reads` verrà
    // restituito il valore richiesto).
    go func() {
        var state = make(map[int]int)
        for {
            select {
            case read := <-reads:
                read.resp <- state[read.key]
            case write := <-writes:
                state[write.key] = write.val
                write.resp <- true
            }
        }
    }()

    // Quì facciamo partire 100 goroutine che eseguiranno
    // delle richieste di lettura dello stato verso la
    // goroutine proprietaria tramite il channel `reads`.
    // Ogni richiesta deve essere eseguita tramite la struct
    // `readOp`, inviata sul channel `reads` e si attende
    // il suo risultato sul channel `resp`.
    for r := 0; r < 100; r++ {
        go func() {
            for {
                read := &readOp{
                    key:  rand.Intn(5),
                    resp: make(chan int)}
                reads <- read
                <-read.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // Avviamo anche altre 10 goroutine che effettueranno
    // delle scritture utilizzando un approccio simile.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                write := &writeOp{
                    key:  rand.Intn(5),
                    val:  rand.Intn(100),
                    resp: make(chan bool)}
                writes <- write
                <-write.resp
                atomic.AddInt64(&ops, 1)
            }
        }()
    }

    // Facciamo eseguire le goroutine per un secondo.
    time.Sleep(time.Second)

    // Infine leggiamo il valore della variabile `ops`.
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)
}
