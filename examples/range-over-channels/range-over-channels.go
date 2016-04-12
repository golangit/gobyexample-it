// In un [esempio precedente](range), abbiamo visto come
// `for` e `range` ci diano la possibilità di reiterare
// su strutture di dati semplici. Possiamo usare la stessa
// sintassi anche per reiterare sui valori ricevuti da un
// channel.

package main

import "fmt"

func main() {

	// Reitereremo su 2 valori nel canale `queue`.
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	// Questo `range` reitera su ogni elemento ogni qual
	// volta che viene ricevuto dal channel `queue`.
	// Poiché abbiamo effettuato un `close` sul canale
	// in precedenza, la reiterazione terminerà dopo aver
	// ricevuto i 2 elementi. Se non avessimo effettuato
	// il `close` sul canale, allora saremmo bloccati ad
	// attendere per un terzo ricevimento nel ciclo.
	for elem := range queue {
		fmt.Println(elem)
	}
}
