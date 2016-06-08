// Le [variabili d'ambiente](https://it.wikipedia.org/wiki/Variabile_d%27ambiente)
// sono un sistema universale per dare informazioni di
// configurazione ai programmi UNIX. Diamo un occhiata a
// come ottenere, impostare e elencare le variabili
// d'ambiente.

package main

import "os"
import "strings"
import "fmt"

func main() {

    // Per impostare una variabile d'ambiente tramite una
    // coppia chiave/valore `os.Setenv`. Per ottenere il
    // valore di una chiave, usa `os.Getenv`. Questo
    // restituirà una stringa vuota nel caso la chiave non
    // sia presente nell'ambiente.
    os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    // Usa `os.Environ` per elencare tutte le coppie
    // chiave/valore nell'ambiente. Questo restituirà una
    // slice di stringhe nella forma `CHIAVE=valore`. Puoi
    // usare `strings.Split` sulle stringhe per ottenere una
    // slice contente al primo posto la chiave e al
    // secondo il valore. Qui stampiamo tutte le chiavi.
    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.Split(e, "=")
        fmt.Println(pair[0])
    }
}
