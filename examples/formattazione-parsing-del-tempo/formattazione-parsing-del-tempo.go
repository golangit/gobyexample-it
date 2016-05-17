// Go supporta la formattazione di struct di tipo `time.Time`
// che permettono di descrivere istanti temporali tramite
// layout personalizzabili.

package main

import "fmt"
import "time"

func main() {
    p := fmt.Println

    // Ecco un esempio di come è possibile formattare un `time`
    // secondo l'RFC3339, utilizzando la costante corrispondente.
    t := time.Now()
    p(t.Format(time.RFC3339))

    // Per effettuare il parsing di un `time` è possibile utilizzare
    // gli stessi parametri di layout di `Format`.
    t1, e := time.Parse(
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)

    // Sia `Format` che `Parse` utilizzano layout basati su esempi.
    // Generalmente si utilizzando costanti di `time` invece di
    // specificare il layout per esteso. È comunque possibile
    // utilizzare layout personalizzati. I layout personalizzati
    // devono utilizzare un tempo fissato di riferimento
    // `Mon Jan 2 15:04:05 MST 2006` per costruire il pattern da
    // utilizzare per formattare/parsare un `time`.
    // Il tempo di riferimento non può essere variato, deve rimanere
    // fissato al 2006, ore 15, etc. pena il fallimmento delle
    // funzioni.
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2, e := time.Parse(form, "8 41 PM")
    p(t2)

    // Se si desidera solamente la rappresentazione numerica
    // di una data, è possibile utilizzare la formattazione
    // standard delle stringhe. La stringa verrà popolata
    // con i valori numerici estratti dal `time`.
    fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
        t.Year(), t.Month(), t.Day(),
        t.Hour(), t.Minute(), t.Second())

    // `Parse` restituirà un errore sugli input malformati,
    // che spiegherà la natura degli errori.
    ansic := "Mon Jan _2 15:04:05 2006"
    _, e = time.Parse(ansic, "8:41PM")
    p(e)
}
