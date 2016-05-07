// Il meccanismo principale per gestire lo stato in Go
// è quello di utilizzare i channel. Abbiamo visto
// questo meccanismo in azione nell'esempio dei
// [worker pool](worker-pools). Ci sono una serie di altri
// meccanismi per gestire lo stato. Adesso vedremo
// come si possa utilizzare il package `sync/atomic`
// per realizzare un _contatore atomico_ che viene
// letto e aggiornato da svariate goroutine.

package main

import "fmt"
import "time"
import "sync/atomic"
import "runtime"

func main() {

    // Utilizzermo un unsigned integer per rappresentare
    // il nostro contatore (che non sarà mai negativo).
    var ops uint64 = 0

    // Per simulare degli aggiornamenti concorrenti,
    // avvieremo 50 goroutine. Queste goroutine incrementeranno
    // il contatore circa ogni millisecondo.
    for i := 0; i < 50; i++ {
        go func() {
            for {
                // Per incrementare in modo atomico il
                // contatore utilizziamo la funzione `AddUint64`
                // passandogli il puntatore alla variabile `ops`
                // utilizzando l'operatore `&`.
                atomic.AddUint64(&ops, 1)

                // Permettiamo alle altre goroutine di procedere.
                runtime.Gosched()
            }
        }()
    }

    // Attendiamo un secondo per permettere alle goroutine
    // di effettuare qualche incremento.
    time.Sleep(time.Second)

    // Al fine di poter leggere il contatore in modo sicuro,
    // mentre viene ancora aggiornato dalle goroutine, possiamo
    // estrarne una copia del valore e salvarlo dentro
    // `opsFinal` tramite la funzione `LoadUint64`.
    // Come sopra, dobbiamo passare a questa funzione
    // il puntatore `&ops` dal quale leggere il valore
    opsFinal := atomic.LoadUint64(&ops)
    fmt.Println("ops:", opsFinal)
}
