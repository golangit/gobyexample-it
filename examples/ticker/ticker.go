// I [timer](timer) possono essere utilizzati
// quando è necessario effettuare qualcosa nel futuro
// nel futuro _solo una volta_. I _ticker_, d'altra parte,
// possono essere utilizzati quando è necessario svolgere
// un'azione qualcosa ripetutamente a intervalli regolari.
// Di seguito un esempio di un ticker che "ticchetta"
// periodicamente finché non lo fermiamo.

package main

import "time"
import "fmt"

func main() {

    // I ticker usano, come i timer, dei channel per
    // notificare che il tempo è scaduto. Di seguito
    // useremo un `range` su un canale per iterare
    // sui valori, che arriveranno ogni 500
    // millisecondi.
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick", t)
        }
    }()

    // I ticker possono essere fermati. Una volta che un
    // ticker è fermato non invierà più alcun valore
    // sul suo canale. Fermeremo il nostro ticker dopo
    // 1600ms.
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    fmt.Println("Ticker fermato")
}
