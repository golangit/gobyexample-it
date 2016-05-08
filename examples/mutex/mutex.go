// Nell'esempio precedente abbiamo visto come gestire
// semplici addizioni utilizzando delle operazioni
// atomiche. Per una gestione dello stato più complessa
// usiamo le _[mutex](https://it.wikipedia.org/wiki/Mutex)_
// per accedere con sicurezza a dati condivisi
// attraverso multiple goroutine.

package main

import (
    "fmt"
    "math/rand"
    "runtime"
    "sync"
    "sync/atomic"
    "time"
)

func main() {

    // Per il nostro esempio lo `stato` sarà una
    // mappa.
    var stato = make(map[int]int)

    // Questa `mutex` sincronizzerà l'accesso a `stato`.
    var mutex = &sync.Mutex{}

    // Per confrontare l'approccio basato sulle mutex
    // piuttosto che altri, `ops` conterà quante
    // operazioni effettuiamo su `stato`.
    var ops int64 = 0

    // Di seguito facciamo partire 100 goroutine che
    // dovranno eseguire ripetute operazioni di lettura
    // sullo stato.
    for r := 0; r < 100; r++ {
        go func() {
            totale := 0
            for {

                // Per ogni lettura scegliamo una chiave
                // per accedere alla map, facciamo un
                // `Lock()` sulla `mutex` per avere
                // controllo esclusivo sullo `stato`,
                // leggiamo il valore alla chiave
                // scelta, sblocchiamo (`Unlock()`) la
                // mutex, e incrementiamo il numero di
                // operazioni (`ops`).
                chiave := rand.Intn(5)
                mutex.Lock()
                totale += stato[chiave]
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)

                // Per essere sicuri che questa goroutine
                // non blocchi lo scheduler, cediamo
                // esplicitamente il controllo dopo ogni
                // operazione utilizzando
                // `runtime.Gosched()`. In genere viene
                // eseguito automaticamente ad esempio
                // con le operazioni sui channel o con
                // funzioni come `time.Sleep`, ma in
                // questo caso abbiamo bisogno di farlo
                // manualmente.
                runtime.Gosched()
            }
        }()
    }

    // Faremo inoltre partire altre 10 goroutine per
    // simulare la scrittura, con lo stesso pattern
    // che abbiamo usato per la lettura.
    for w := 0; w < 10; w++ {
        go func() {
            for {
                chiave := rand.Intn(5)
                valore := rand.Intn(100)
                mutex.Lock()
                stato[chiave] = valore
                mutex.Unlock()
                atomic.AddInt64(&ops, 1)
                runtime.Gosched()
            }
        }()
    }

    // Lasciamo ora lavorare le 10 goroutine per un
    // secondo.
    time.Sleep(time.Second)

    // Prendiamo e stampiamo il numero di operazioni
    // finale.
    opsFinal := atomic.LoadInt64(&ops)
    fmt.Println("ops:", opsFinal)

    // Con un lock finale sullo `stato`, vediamo come
    // si è evoluto lo stato.
    mutex.Lock()
    fmt.Println("stato:", stato)
    mutex.Unlock()
}
