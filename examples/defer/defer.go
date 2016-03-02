// Il built-in _defer_ può essere utilizzato per assicurarci
// che una funzione venga eseguita in un secondo momento,
// ad esempio per scopi di pulizia o per rilasciare delle
// risorse. Può essere assimilata alle keywork `ensure` o
// `finally` di altri linguaggi.
package main

import "fmt"
import "os"

// Supponiamo di voler creare un file, ci scriviamo
// all'interno e dopo vogliamo chiuderlo. Ecco come possiamo
// farlo utilizzando la `defer`.
func main() {

	// Subito dopo aver creato un file con la funzione
	// `createFile`, utilizziamo la `defer` su `closeFile`.
	// Questa funzione sarà eseguita al termine della funzione
	// dentro la quale è stata chiamata (in questo caso
	// il `main`), appena la funzione `writeFile` avrà terminato.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
