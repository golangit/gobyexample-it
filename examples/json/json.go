// Go ha il supporto builtin per l'encoding e il decoding
// di JSON, con anche la possibilità di convertire a o da
// tipi builtin e personalizzati.

package main

import "encoding/json"
import "fmt"
import "os"

// Useremo questi due struct per dimostrare l'encoding e
// il decoding di tipi personalizzati.
type Response1 struct {
	Pagina int
	Frutti []string
}
type Response2 struct {
	Pagina int      `json:"pagina"`
	Frutti []string `json:"frutti"`
}

func main() {

	// Come prima cosa vedremo come fare l'encoding di
	// datatype basici per ottenere stringhe JSON. Ecco
	// degli esempi per dei valori atomici.
	bolB, _ := json.Marshal(true)
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	// E ora degli esempi con le slice e le map, che
	// ritornano degli array e oggetti JSON come ti
	// aspetteresti.
	slcD := []string{"mela", "pesca", "pera"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"mela": 5, "lattuga": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	// Il package JSON può encoddare automaticamente
	// dei data type personalizzati, ad esempio delle
	// struct. Includerà solo i field esportati
	// nell'output encoddato e di default userà i nomi
	// dei field per le chiavi JSON.
	res1D := &Response1{
		Pagina: 1,
		Frutti: []string{"mela", "pesca", "pera"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	// Puoi usare dei "tag" sulle dichiarazioni dei field
	// di una struct per personalizzare il nome della
	// chiave nel JSON finale. Vedi la dichiarazione di
	// `Response2` sopra per vedere un esempio di questi
	// tag.
	res2D := &Response2{
		Pagina: 1,
		Frutti: []string{"mela", "pesca", "pera"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	// Ora diamo un'occhiata al decoding di valori JSON
	// in valori builtin di Go. Di seguito un esempio di
	// una struttura di dati generica.
	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	// Abbiamo bisogno di definire una variabile dove
	// il package JSON metterà i dati decoddati. Questa
	// `map[string]interface{}` sarà una map di stringhe
	// che puntano a dati con datatype arbitrario.
	var dat map[string]interface{}

	// Qui avviene il vero decoding, dove controlleremo
	// anche se c'è stato un errore.
	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Println(dat)

	// Per poter usare i dati nella mappa decoddata,
	// avremo bisogno di fare un type assertion al loro
	// vero valore. Ad esempio, qui facciamo un type
	// assertion del valore in `num` in un `float64`.
	num := dat["num"].(float64)
	fmt.Println(num)

	// Accedere a dati annidati richiede una serie
	// di type assertion.
	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	// Possiamo anche fare il decode di JSON in data
	// type personalizzati. Questo ha il vantaggio di
	// aggiungere un'ulteriore thread-safety ai nostri
	// programmi e elimina la necessità dell'uso dei
	// type assertion quando accediamo a dati decoddati.
	str := `{"pagina": 1, "frutti": ["mela", "pesca"]}`
	res := Response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Println(res)
	fmt.Println(res.Frutti[0])

	// In the examples above we always used bytes and
	// strings as intermediates between the data and
	// JSON representation on standard out. We can also
	// stream JSON encodings directly to `os.Writer`s like
	// `os.Stdout` or even HTTP response bodies.
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"mela": 5, "lettuce": 7}
	enc.Encode(d)
}
