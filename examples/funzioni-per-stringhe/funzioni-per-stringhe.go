// Il pacchetto `strings` della standard library contiene
// svariate funzioni per la gestione delle stringhe.
// Ecco qualche esempio di funzioni tratte da questo pacchetto

package main

import s "strings"
import "fmt"

// Diamo alla funzione `fmt.Println` un alias (`p`) in quanto
// l'andremo ad usare parecchio in questo esempio
var p = fmt.Println

func main() {

    // Ecco un esempio delle funzioni disponibili nel
    // pacchetto `strings`. Nota che queste sono tutte
    // funzioni prevenienti dal pacchetto (e non proprie
    // del tipo string). Questo significa che è necessario
    // passare la stringa come primo parametro ogni volta
    p("Contains:  ", s.Contains("test", "es"))
    p("Count:     ", s.Count("test", "t"))
    p("HasPrefix: ", s.HasPrefix("test", "te"))
    p("HasSuffix: ", s.HasSuffix("test", "st"))
    p("Index:     ", s.Index("test", "e"))
    p("Join:      ", s.Join([]string{"a", "b"}, "-"))
    p("Repeat:    ", s.Repeat("a", 5))
    p("Replace:   ", s.Replace("foo", "o", "0", -1))
    p("Replace:   ", s.Replace("foo", "o", "0", 1))
    p("Split:     ", s.Split("a-b-c-d-e", "-"))
    p("ToLower:   ", s.ToLower("TEST"))
    p("ToUpper:   ", s.ToUpper("test"))
    p()

    // Puoi trovare l'elenco di tutte le funzioni del pacchetto
    // `strings` sulla [documentazione ufficiale](http://golang.org/pkg/strings/)

    // Queste ultimi due comandi non fanno parte del pacchetto,
    // ma vale comunque la pena citarli: ottenere la lunghezza
    // di una stringa ed ottenere un carattere tramite indice
    p("Len: ", len("hello"))
    p("Char:", "hello"[1])
}
