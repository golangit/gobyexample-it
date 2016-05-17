// Go offre la possibilità di formattare le stringhe con
// il classico sistema di `printf`. Qui di seguito alcuni
// esempi di operazioni comuni di string formatting.

package main

import "fmt"
import "os"

type punto struct {
    x, y int
}

func main() {

    // Go offre numerosi modi per stampare dati, alcuni
    // disegnati appositamente per stampare strutture
    // di dati del nostro programma. Per esempio, `%v`
    // stampa la composizione della nostra struct `punto`.
    p := punto{1, 2}
    fmt.Printf("%v\n", p)

    // Se il valore è una struct, la variante `%+v`
    // includerà i nomi dei campi della struct in sé.
    fmt.Printf("%+v\n", p)

    // La variante `%#v` stampa una rappresentazione nella
    // sintassi di Go, ovvero ciò che viene stampato
    // potrebbe potenzialmente essere usato come codice
    // per produrre quel valore.
    fmt.Printf("%#v\n", p)

    // Per stampare il tipo di un valore, usa `%T`.
    fmt.Printf("%T\n", p)

    // Puoi, inoltre, stampare bool con `%t`.
    fmt.Printf("%t\n", true)

    // Ci sono, inoltre, numerose opzioni per stampare
    // integer. Puoi usare `%d` per avere una
    // rappresentazione standard in base 10.
    fmt.Printf("%d\n", 123)

    // `%b` stampa la rappresentazione binaria.
    fmt.Printf("%b\n", 14)

    // `%c` stampa il carattere corrispondente
    // all'integer dato.
    fmt.Printf("%c\n", 33)

    // Con `%x` puoi stampare il valore in esadecimale.
    fmt.Printf("%x\n", 456)

    // Ci sono anche numerose opzioni per il formatting
    // dei float. Per stampare un float normalmente, usa
    // `%f`.
    fmt.Printf("%f\n", 78.9)

    // `%e` e `%E` stampano il float in (versioni
    // leggermente differenti di) notazione
    // scientifica.
    fmt.Printf("%e\n", 123400000.0)
    fmt.Printf("%E\n", 123400000.0)

    // Per stampare stringhe normalmente, usa `%s`.
    fmt.Printf("%s\n", "\"string\"")

    // Per mettere le virgolette alle stringhe come nel
    // sorgente di Go, puoi usare `%q`.
    fmt.Printf("%q\n", "\"string\"")

    // Come per gli interi, si può stampare una versione
    // esadecimale di una stringa sempre con `%x`, con
    // due caratteri esadecimale per ogni byte passato.
    fmt.Printf("%x\n", "esadecimalami")

    // Per stampare la rappresentazione di un putatore,
    // usa `%p`.
    fmt.Printf("%p\n", &p)

    // Quando devi stampare dei numeri puoi spesso
    // aver bisogno di controllare la lunghezza e la
    // precisione della stringa risultante. Per
    // specificare la lunghezza di un integer, dai un
    // numero dopo il `%`. Normalmente il risultato
    // sarà allineato a destra con degli spazi.
    fmt.Printf("|%6d|%6d|\n", 12, 345)

    // Puoi anche specificare la lunghezza dei float
    // che vuoi stampare, anche se spesso vorrai anche
    // dare la lunghezza della precisione decimale
    // allo stesso tempo della precisione del numero
    // intero.
    fmt.Printf("|%6.2f|%6.2f|\n", 1.2, 3.45)

    // Per allineare alla sinistra, usa la flag `-`.
    fmt.Printf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)

    // Potrai anche aver bisogno di controllare la
    // lunghezza quando stampi delle stringhe, soprattutto
    // quando vuoi che si allineino per fare un output
    // "a mò di tabella". Puoi sempre specificare un
    // numero per allineare alla destra.
    fmt.Printf("|%6s|%6s|\n", "foo", "b")

    // Per allineare alla destra, usa la flag `-` come
    // per i numeri.
    fmt.Printf("|%-6s|%-6s|\n", "foo", "b")

    // Fino ad ora abbiamo visto `Printf`, che stampa
    // la stringa formattata a `os.Stdout`. `Sprintf`
    // formatta la stringa e la restituisce senza stamparla
    // da nessuna parte.
    s := fmt.Sprintf("una %s", "stringa")
    fmt.Println(s)

    // Puoi, inoltre, formattare+stampare a degli
    // `io.Writer` oltre a `os.Stdout` tramite
    // `Fprintf`.
    fmt.Fprintf(os.Stderr, "un %s\n", "errore")
}
