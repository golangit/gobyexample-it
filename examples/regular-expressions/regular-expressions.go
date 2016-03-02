// Go offre il supporto nativo alle [espressioni regolari](https://it.wikipedia.org/wiki/Espressione_regolare).
// Ecco qualche esempio di operazioni comuni con le regexp in Go.

package main

import "bytes"
import "fmt"
import "regexp"

func main() {

	// Questo controlla se una stringa soddisfa una regexp
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	// Prima abbiamo utilizzato le regexp in modo diretto.
	// Nel caso generale è necessario invocare prima il
	// `Compile` su una regexp, al fine di ottenere una
	// struct `Regexp` ottimizzata.
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Su questa struttura abbiamo a disposizione svariati
	// metodi. Ecco un test simile al nostro primo esempio.
	fmt.Println(r.MatchString("peach"))

	// Questo ritorna il primo match (una stringa) che
	// soddisfa la regexp.
	fmt.Println(r.FindString("peach punch"))

	// Anche questo metodo ritorna il primo match,
	// ma invece di ritornarne la stringa, ritorna l'indice
	// di inizio e di fine della stringa.
	fmt.Println(r.FindStringIndex("peach punch"))

	// La variante `Submatch` include informazioni
	// relativi al match del pattern e anche ai match di
	// eventuali sotto-pattern. In questo caso ritornerebbe
	// i match per i pattern `p([a-z]+)ch` e `([a-z]+)`.
	fmt.Println(r.FindStringSubmatch("peach punch"))

	// In modo simile questo metodo ritornerà gli indici
	// di inizio e di fine del match e dei sotto-match
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))

	// La variante `All`, invece di ritornare solo la prima
	// occorrenza del match, ritorna tutti i match.
	fmt.Println(r.FindAllString("peach punch pinch", -1))

	// Questa variante `All` è disponibile per tutte le
	// funzioni che abbiamo visto prima.
	fmt.Println(r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	// Se si passa un interno non negativo come secondo parametro
	// di queste funzioni si sta limitando il numero
	// di match da ritornare. Ad esempio in questo caso
	// si è interessati solamente ai primi 2 match.
	fmt.Println(r.FindAllString("peach punch pinch", 2))

	// Il nostro primo esempio mostrava la funzione
	// `MatchString` che lavorava su stringhe. È possibile
	// passare anche parametro di tipo `[]string` ed
	// utilizzare la funzione `Match`.
	fmt.Println(r.Match([]byte("peach")))

	// Se si vuole creare una costante a partire da
	// una regexp si può utilizzare la funzione `MustCompile`
	// (invece di `Compile`). La classica `Compile` non
	// funzionerà in quanto ha 2 valori di ritorno.
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(r)

	// Il pacchetto `regexp` può essere utilizzato anche per
	// effettuare sostituzioni di stringhe/sottostringhe.
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>"))

	// La variante `Func` permette di trasformare i match
	// trovati passandogli una funzione che verrà applicata
	// su ogni occorrenza
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out))
}
