// Gli URL offrono un modo semplice ed univoco per [localizzare le risorse](https://it.wikipedia.org/wiki/Uniform_Resource_Locator).
// Ecco come è possibile effettuare il parsing
// degli URL con Go.

package main

import "fmt"
import "net"
import "net/url"

func main() {

    // Effettueremo il parsing di questo URL, che include
    // uno schema, informazioni sull'autenticazione, host,
    // porta, path, parametri della query e fragment.
    s := "postgres://user:pass@host.com:5432/path?k=v#f"

    // Effettuiamo il parsing per assicurarci che non
    // ci siano errori.
    u, err := url.Parse(s)
    if err != nil {
        panic(err)
    }

    // Ottenere lo schema è semplice.
    fmt.Println(u.Scheme)

    // `User` contiene tutte le informazioni sull'autenticazione.
    // Invoca `Username` e `Password` per ottenere i singoli valori
    fmt.Println(u.User)
    fmt.Println(u.User.Username())
    p, _ := u.User.Password()
    fmt.Println(p)

    // Il campo `Host` contiene informazioni su hostname e porta,
    // se presente. Utilizza `SplitHostPort` per estrarre la porta.
    fmt.Println(u.Host)
    host, port, _ := net.SplitHostPort(u.Host)
    fmt.Println(host)
    fmt.Println(port)

    // Qui estraiamo il `path` e il fragment dopo il
    // carattere `#`.
    fmt.Println(u.Path)
    fmt.Println(u.Fragment)

    // Per ottenere i parametri della query in un
    // formato string del tipo `k=v` è possibile utilizzare
    // `RawQuery`. È possibile effettuare il parsing
    // dei parametri in una map tramite `ParseQuery`.
    // Le chiavi della map saranno di tipo string, mentre
    // i valori saranno di tipo slice di string, per cui
    // è possibile recuperare il primo valore con `[0]`.
    fmt.Println(u.RawQuery)
    m, _ := url.ParseQuery(u.RawQuery)
    fmt.Println(m)
    fmt.Println(m["k"][0])
}
