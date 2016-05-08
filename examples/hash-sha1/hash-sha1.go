// La funzione [_SHA-1_](https://it.wikipedia.org/wiki/Secure_Hash_Algorithm)
// viene frequentemente utilizzata per calcolare il _digest_
// (l'impronta) dei file binari o di testo.
// Ad esempio il sistema di versionamento [git](http://git-scm.com/)
// utilizza gli SHA-1 in modo intensivo per identificare
// i file e le cartelle versionate. Ecco come calcolare
// gli hash SHA-1 in Go.

package main

// Si possono trovare implementazioni di svariati algoritmi
// crittografici dentro il pacchetto `crypto/*`.
import "crypto/sha1"
import "fmt"

func main() {
    s := "sha1 this string"

    // Il pattern per creare un nuovo hash è `sha1.New()`,
    // `sha1.Write(bytes)`, ed infine `sha1.Sum([]byte{})`.
    // Iniziamo creando un nuovo hash.
    h := sha1.New()

    // La funzione `Write` lavora con i bytes.
    // Se si deve calcolare l'hash di una stringa `s`
    // è possibile utilizzare `[]byte(s)` per ottenerne i bytes.
    h.Write([]byte(s))

    // Con la funzione `Sum` è possibile ottenere il valore
    // dell'hash. L'argomento può essere utilizzato per
    // accodare il valore calcolato ad uno slice di byte
    // esistente, non è generalmente necessario.
    bs := h.Sum(nil)

    // Gli SHA-1 sono generalmente stampati in formato
    // esadecimale (ad esempio nei commit di git). Utilizza
    // '%x' per stampare il valore in esadecimale.
    fmt.Println(s)
    fmt.Printf("%x\n", bs)
}
