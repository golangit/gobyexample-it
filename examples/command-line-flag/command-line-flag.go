// I _Command-line flag_ sono uno dei modi prioritari
// per specificare delle opzioni in una qualsiasi
// [interfaccia a riga di comando](https://it.wikipedia.org/wiki/Interfaccia_a_riga_di_comando)
// Ad esempio in `wc -l`, il `-l` è un flag.

package main

// Go offre il pacchetto `flag` per supportare il parsing
// dei flag da linea di comando. Utilizzeremo questo pacchetto
// per implementare il nostro programma di esempio.
import "flag"
import "fmt"

func main() {

    // Il pacchetto offre semplici funzioni per parsare
    // stringhe, interi e booleani. Qui dichiariamo un
    // flag di tipo string chiamato `word` che ha come
    // valore di default `"foo"` ed una breve descrizione.
    // Questa funzione `flag.String` restituisce il puntatore
    // ad una stringa (non un valore di tipo stringa).
    // Vedremo dopo come utilizzare questo puntatore.
    wordPtr := flag.String("word", "foo", "a string")

    // Qui dichiariamo due flag, `numb` e `fork`,
    // utilizzando un approccio simile a `word`.
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // È inoltre possibile dichiarare un flag che vada ad
    // utilizzare una variabile precedentemente dichiarata
    // nel programma con `var`. Nota che è necessario passare
    // il puntatore a tale variabile.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

    // Una volta che tutti i flag sono stati dichiarati è
    // possibile invocare `flag.Parse()` per eseguire il
    // parsing dei flag dalla linea di comando.
    flag.Parse()

    // Qui stampiamo i valori dei flag. Otterremo i valori
    // di default se i flag non sono presenti sulla linea
    // di comando. Nota che è necessario dereferenziare il
    // puntatore (vedi `*wordPtr`) per ottenere il valore
    // del flag.
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}
