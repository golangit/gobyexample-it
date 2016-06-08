// Usa `os.Exit` per terminare immediatamente l'esecuzione
// di un programma con un determinato [valore di uscita](https://it.wikipedia.org/wiki/Valore_di_uscita).

package main

import "fmt"
import "os"

func main() {

    // I `defer` _non_ saranno eseguiti con `os.Exit`,
    // quindi questo `fmt.Println` non verrà mai eseguito.
    defer fmt.Println("!")

    // Terminiamo il programma col valore di uscita 3.
    os.Exit(3)
}

// Nota che a differenza di C e simili, Go non usa un
// valore restituito da `main()` per indicare il valore
// di uscita. Se vuoi terminare il programma con un valore
// non-zero, usa `os.Exit`.
