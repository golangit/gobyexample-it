// Go offre il supporto built-in per la codifica [base64](https://it.wikipedia.org/wiki/Base64).

package main

// Con questa sintassi si imposta il pacchetto
// `encoding/base64` con il nome `b64` invece del default `base64`.
// Questo ci permetterà di risparmiare spazio successivamente.
import b64 "encoding/base64"
import "fmt"

func main() {

    // Ecco una `string` che andremo a codificare in base64
    data := "abc123!?$*&()'-=@~"

    // Go supporta base64 sia standard che URL-compatible.
    // Qui si mostra coma utilizzare lo standard encoding.
    // L'encoder richiede un `[]byte` quindi effettuiamo
    // un cast della nostra `string` a quel tipo.
    sEnc := b64.StdEncoding.EncodeToString([]byte(data))
    fmt.Println(sEnc)

    // Il decoding può restituire un errore, che può essere
    // utile analizzare se non si è sicuri che il valore
    // in input sia una string base64 valida.
    sDec, _ := b64.StdEncoding.DecodeString(sEnc)
    fmt.Println(string(sDec))
    fmt.Println()

    // Questo codifica/decodifica utilizzando un formato
    // base64 URL-compatible
    uEnc := b64.URLEncoding.EncodeToString([]byte(data))
    fmt.Println(uEnc)
    uDec, _ := b64.URLEncoding.DecodeString(uEnc)
    fmt.Println(string(uDec))
}
