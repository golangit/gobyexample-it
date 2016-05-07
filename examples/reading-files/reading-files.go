// Leggere e scrivere su file è un task basilare
// necessario ad una marea di programmi. Vediamo prima
// qualche esempio di lettura da file.

package main

import (
    "bufio"
    "fmt"
    "io"
    "io/ioutil"
    "os"
)

// La lettura da file richiede il controllo degli errori
// di svariate funzioni. Questo helper ci aiuterà a
// gestire gli errori che si potrebbero generare.
func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Il task più semplice è la lettura dell'intero
    // file in memoria
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    fmt.Print(string(dat))

    // Generalmente si vuole più controllo sul modo con
    // viene letto un file. Per queste funzionalità si
    // deve iniziare aprendo il file (con `os.Open`) al fine
    // di ottenere un valore di tipo `os.File`.
    f, err := os.Open("/tmp/dat")
    check(err)

    // Con `Read` è possibile leggere byte dall'inizio del file.
    // Qui si permette di leggere fino a 5 byte, ma si ottiene
    // anche il numero di byte effettivamente letti (nel
    // parametro n1).
    b1 := make([]byte, 5)
    n1, err := f.Read(b1)
    check(err)
    fmt.Printf("%d bytes: %s\n", n1, string(b1))

    // È possibile posizionarsi in un punto preciso del file
    // tramite `Seek` ed effettuare una `Read` da quel punto
    o2, err := f.Seek(6, 0)
    check(err)
    b2 := make([]byte, 2)
    n2, err := f.Read(b2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2))

    // Il pacchetto `io` offre svariate funzioni che possono
    // essere utili durante la lettura da file. Ad esempio,
    // una read come quella poco sopra può essere realizzata
    // tramite la funzione `ReadAtLeast`
    o3, err := f.Seek(6, 0)
    check(err)
    b3 := make([]byte, 2)
    n3, err := io.ReadAtLeast(f, b3, 2)
    check(err)
    fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

    // Per posizionarsi all'inizio del file è possibile
    // utilizzare `Seek(0, 0)`
    _, err = f.Seek(0, 0)
    check(err)

    // Il pacchetto `bufio` offre una serie di metodi
    // per la lettura bufferizzata su file.
    // Questo pacchetto può essere utile per la sua
    // efficienza nel caso di svariate piccole read, e per
    // le funzioni addizionali che offre.
    r4 := bufio.NewReader(f)
    b4, err := r4.Peek(5)
    check(err)
    fmt.Printf("5 bytes: %s\n", string(b4))

    // Ricordati di chiudere il file quando hai terminato.
    // Questo può essere effettuato subito dopo la `Open`
    // utilizzando una `defer`.
    f.Close()
}
