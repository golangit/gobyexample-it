// Un `panic` in genere significa che qualcosa è andato
// fin troppo storto. In genere lo usiamo per fallimenti
// immediati su errori che non dovrebbero succedere
// durante una normale esecuzione del programma, o per
// cui non è possibile effettuare procedure di recovery

package main

import "os"

func main() {

    // D'ora in poi utilizzeremo i `panic` su questo
    // sito per controllare errori inaspettati. Questo
    // è l'unico programma sul sito fatto apposta
    // perché risulti in un `panic`.
    panic("un problema")

    // Un uso comune del panic è l'interruzione del
    // programma quando ci troviamo davanti ad un errore
    // restituito da una funzione che non sappiamo come
    // gestire o non vogliamo gestire per niente. Di
    // seguito un esempio di `panic` quando riceviamo
    // un errore inaspettato dopo aver tentato di creare
    // un nuovo file.
    _, err := os.Create("/tmp/file")
    if err != nil {
        panic(err)
    }
}
