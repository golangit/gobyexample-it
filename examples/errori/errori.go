// In Go si è soliti comunicare errori attraverso un
// valore restituito esplicito e separato. Questo
// diversamente sia da quanto accade con le eccezioni
// usate in linguaggi come Java e Ruby che dall'overloading
// di un singolo valore come risultato / errore
// che è a volte usato in C. L'approccio
// usato da Go rende facile capire quali funzioni possono
// generare errori e gestirli attraverso gli stessi
// costrutti utilizzati per qualsiasi altra attività che
// non riguardi la gestione degli errori.

package main

import "errors"
import "fmt"

// Per convenzione, gli errori sono l'ultimo valore
// restituito, ed il loro tipo è `error`,
// un'interfaccia built-in.
func f1(arg int) (int, error) {
    if arg == 42 {

        // `errors.New` crea un errore base con il
        // messaggio di errore dato.
        err := errors.New("impossibile calcolare con 42")
        return -1, err

    }

    // Un `nil` nella posizione dell'errore indica che non
    // vi è stato alcun errore.
    return arg + 3, nil
}

// È possibile usare altri tipi come `error`
// implementando il metodo `Error()` su di essi. Di
// seguito una variante del precedente esempio che
// usa un tipo a sé per rappresentare esplicitamente un
// errore di parametri.
type argError struct {
    arg  int
    prob string
}

func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
    if arg == 42 {

        // In questo caso, usiamo la sintassi
        // `&argError` per costruire una nuova struct ed
        // inserire i valori dei due campi, `arg` e `prob`.
        return -1, &argError{arg, "impossibile calcolare"}
    }
    return arg + 3, nil
}

func main() {

    // I due cicli che seguono provano ognuna delle nostre
    // funzioni che restituiscono un errore. Nota che l'uso di
    // un controllo di errori sulla stessa linea dell'`if`
    // è una pratica comune in Go.
    for _, i := range []int{7, 42} {
        if r, e := f1(i); e != nil {
            fmt.Println("f1 ha fallito   :", e)
        } else {
            fmt.Println("f1 ha funzionato:", r)
        }
    }
    for _, i := range []int{7, 42} {
        if r, e := f2(i); e != nil {
            fmt.Println("f2 ha fallito   :", e)
        } else {
            fmt.Println("f2 ha funzionato:", r)
        }
    }

    // Se vuoi utilizzare qualche campo specifico di
    // un errore personalizzato, devi convertire
    // l'errore in un'istanza dell'errore personalizzato
    // tramite un type assertion.
    _, e := f2(42)
    if ae, ok := e.(*argError); ok {
        fmt.Println(ae.arg)
        fmt.Println(ae.prob)
    }
}
