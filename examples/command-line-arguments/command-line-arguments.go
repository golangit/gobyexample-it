// Gli _argomenti della linea di comando_ sono un metodo
// comune per passare parametri ai programmi.
// Ad esempio, `go run hello.go` usa `run` e `hello.go`
// come parametri al programma `go`.

package main

import "os"
import "fmt"

func main() {

    // `os.Args` fornisce accesso dirretto agli argomenti
    // della linea di comando. Nota che il primo valore
    // in questa slice è il percorso del programma, e
    // `os.Args[1:]` contiene i veri argomenti del
    // programma.
    argsConProg := os.Args
    argsSenzaProg := os.Args[1:]

    // Puoi ottenere gli argomenti individualmente con
    // l'indexing.
    arg := os.Args[3]

    fmt.Println(argsConProg)
    fmt.Println(argsSenzaProg)
    fmt.Println(arg)
}
