// Scrivere file in Go segue dei pattern simili a quelli
// che abbiamo visto in precedenza per la lettura.

package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "os"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {

    // Come prima cosa, vediamo come si può scrivere una
    // stringa (o dei byte) in un file.
    d1 := []byte("ciao\ngo\n")
    err := ioutil.WriteFile("/tmp/dat1", d1, 0644)
    check(err)

    // Per delle scritture più "scandite", apriamo un file
    // per la scrittura.
    f, err := os.Create("/tmp/dat2")
    check(err)

    // È idiomatico fare un defer di `Close` subito dopo
    // aver aperto un file (ed aver controllato
    // eventuali errori, ovviamente.)
    defer f.Close()

    // Puoi usare `Write` per scrivere slice di byte,
    // come ci si aspetterebbe.
    d2 := []byte{100, 101, 105, 10}
    n2, err := f.Write(d2)
    check(err)
    fmt.Printf("scritti %d byte\n", n2)

    // Per comodità, c'è anche `WriteString`.
    n3, err := f.WriteString("write\n")
    fmt.Printf("scritti %d byte\n", n3)

    // Usa un `Sync` per assicurarti che i dati siano
    // scritti su disco.
    f.Sync()

    // Il package `bufio` dispone di writer
    // bufferizzati, oltre ai reader bufferizzati che
    // abbiamo visto in precenza.
    w := bufio.NewWriter(f)
    n4, err := w.WriteString("bufferizzati\n")
    fmt.Printf("scritti %d byte\n", n4)

    // Anche in questo caso, usa `Flush` per essere sicuro
    // che tutte le operazioni siano state apportate sul
    // writer vero e proprio.
    w.Flush()

}
