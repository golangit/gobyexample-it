// I _range_ statement reiterano sugli elementi di una
// varietà di strutture di dati, similmente ai foreach di
// altri linguaggi. Vediamo come usare `range` con alcune
// delle strutture di dati che abbiamo già visto.

package main

import "fmt"

func main() {

	// Di seguito utilizziamo un `range` per sommare i
	// numeri di uno slice. (Con gli array viene fatto
	// allo stesso modo).
	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("somma:", sum)

	// `range` sugli array e gli slice ritorna sia l'indice
	// sia il valore di ognuno degli elementi. Prima non
	// avevamo bisogno di utilizzare l'indice, quindi
	// l'avevamo ignorato utilizzando il blank identifier
	// `_`. Qualche volta potremmo avere anche solo bisogno
	// dell'indice, e ignorare il suo valore.
	for i, num := range nums {
		if num == 3 {
			fmt.Println("indice:", i)
		}
	}

	// `range` su una map reitera sulle coppie
	// chiave-valore.
	kvs := map[string]string{"a": "arancia", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// `range` sulle stringhe reitera sugli Unicode code
	// point. Il primo valore è l'indice del `rune` nella
	// stringa, il secondo è il `rune` in sé.
	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
