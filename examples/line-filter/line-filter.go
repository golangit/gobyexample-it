// Un _line filter_ è un tipo di programma molto comune
// che legge l'input da stdin, lo processa, e ne stampa
// l'output su stdout. `grep` e `sed` sono comuni esempi
// di _line filter_.

// Ecco un esempio di line filter scritto in Go, che scrive
// in maiscuolo tutto il testo che riceve in input. Puoi
// utilizzare questo scheletro di line filter per realizzare
// il tuo line filter.
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {

    // Tramite `NewScanner` creiamo un wrapper intorno a
    // `os.Stdin` (che è non bufferizzato), permettendoci
    // di invocare la funzione `Scan` per ottenere il
    // prossimo token (in questo caso la prossima linea).
    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        // `Text` restituisce il token corrente, in questo caso
        // la prossima linea dall'input.
        ucl := strings.ToUpper(scanner.Text())

        // Stampa la linea in maiuscolo.
        fmt.Println(ucl)
    }

    // Controlla eventuali errori durante lo `Scan`.
    // Un _End of file_ è atteso dallo scanner e non viene
    // quindi considerato un errore
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "error:", err)
        os.Exit(1)
    }
}
