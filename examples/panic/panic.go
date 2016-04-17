// A `panic` typically means something went unexpectedly
// wrong. Mostly we use it to fail fast on errors that
// shouldn't occur during normal operation, or that we
// aren't prepared to handle gracefully.
// Un `panic` in genere significa che qualcosa è andato
// fin troppo storto. In genere lo usiamo per fallimenti
// immediati su errori che non dovrebbero succedere
// durante una normale esecuzione del programma, o che non
// siamo pronti a gestire con grazia.

package main

import "os"

func main() {

	// Da ora utilizzeremo i `panic`` su questo sito per
	// controllare errori inaspettati. Questo è l'unico
	// programma sul sito fatto apposta perché risulti
	// in un `panic`.
	panic("un problema")

	// Un uso comune dei panic è per interrompere il
	// programma quando ci troviamo davanti ad un errore
	// ritornato da una funzione che non sappiamo come
	// gestire o non vogliamo gestire per niente. Di
	// seguito un esempio di `panic` quando riceviamo
	// un errore inaspettato dopo aver tentato di creare
	// un nuovo file.
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
