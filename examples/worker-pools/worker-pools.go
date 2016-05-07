// In questo esempio vedremo come implementare
// un _worker pool_ usando le goroutine e i channel.

package main

import "fmt"
import "time"

// Questo è il worker, sul quale eseguiremo i nostri
// task concorrenti. Questi worker riceveranno i task da
// eseguire sul channel `jobs` ed invieranno i risultati sul
// channel `results`. In questo esempio abbiamo inserito una
// `Sleep` da un secondo per simulare un task oneroso.
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        fmt.Println("worker", id, "sta eseguendo task", j)
        time.Sleep(time.Second)
        results <- j * 2
    }
}

func main() {

    // Per poter utilizzare il pool di worker dobbiamo
    // poter inviare i task e poter ricevere i risultati,
    // creiamo dunque due channel per questo
    jobs := make(chan int, 100)
    results := make(chan int, 100)

    // Questo farà partire 3 worker, che saranno inizialmente
    // bloccati in quanto non hanno task da eseguire.
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }

    // Inviamo 9 task da eseguire sul channel `job` ed invochiamo
    // la `close` su quel canale, in modo da indicare che non
    // ci sono altri task da eseguire.
    for j := 1; j <= 9; j++ {
        jobs <- j
    }
    close(jobs)

    // Infine recuperiamo i risultati delle computazioni sul
    // channel `results`
    for a := 1; a <= 9; a++ {
        <-results
    }
}
