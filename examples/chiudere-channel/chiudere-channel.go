// Chiudere un channel serve ad indicare che non verranno
// più inviati su quel channel.
// Questo è utile per comunicare il completamento di
// un operazione a chi sta ascoltando sul channel.

package main

import "fmt"

// In questo esempio useremo un channel `jobs` per
// inviare i task che devono essere eseguiti dalla
// goroutine `main()` a una goroutine che svolge il
// ruolo di `worker`. Quando non si hanno più task
// da eseguire si chiuderà il channel `jobs` tramite
// il built-in `close`.
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // Questa è la goroutine worker, che riceverà i task
    // in modo continuativo dal channel `jobs` tramite il
    // comando `j, more := <-jobs`. Utilizzando questa forma
    // speciale per ricevere i valori, nella variabile
    // `more` avremo il valore `false` se il channel `jobs`
    // è stato chiuso e tutti i valori sono stati ricevuti.
    // Utilizzeremo questo valore per indicare quando la
    // goroutine ha terminato di eseguire tutti i task che
    // ha ricevuto.
    go func() {
        for {
            j, more := <-jobs
            if more {
                fmt.Println("received job", j)
            } else {
                fmt.Println("received all jobs")
                done <- true
                return
            }
        }
    }()

    // Questo invia 3 job alla goroutine worker sul channel
    // `jobs` e lo chiude subito dopo.
    for j := 1; j <= 3; j++ {
        jobs <- j
        fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    // Attendiamo la goroutine worker utilizzando l'approccio
    // di [sincronizzazione](sincronizzazione-dei-channel) che abbiamo
    // visto negli esempi precedenti
    <-done
}
