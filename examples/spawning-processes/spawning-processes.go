// Talvolta i programmi Go hanno bisogno di far partire
// altri processi, che possono non essere scritti in Go. Ad esempio
// l'highlight della sintassi in questo sito è [realizzata](https://github.com/golangit/gobyexample-it/blob/master/tools/generate.go)
// facendo partire un processo [`pygmentize`](http://pygments.org/)
// da un programma Go.
// Vediamo qualche esempio su come poter eseguire processi
// arbitrari da Go.
package main

import "fmt"
import "io/ioutil"
import "os/exec"

func main() {

    // Inizieremo con un programma semplice, che non ha argomenti
    // o input ma che stampa solamente sul stdout. La funzione
    // `exec.Command` è un helper che rappresenta il processo
    // che andremo ad eseguire
    dateCmd := exec.Command("date")

    // Tramite la funzione `.Output` è possibile eseguire il
    // processo, attendere che finisca ed ottenere il output.
    // Se non ci sono stati errori `dateOut` conterrà i byte
    // della data.
    dateOut, err := dateCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> date")
    fmt.Println(string(dateOut))

    // Adesso vedremo un esempio leggermente più complesso
    // dove forniremo alcuni dati sullo stdin di un processo,
    // e raccoglieremo il suo output sullo stdout.
    grepCmd := exec.Command("grep", "ciao")

    // Qui possiamo vedere come ottenere lo stdin e lo stdout
    // di un futuro comando, scrivere dei dati sullo stdin,
    // leggere dei dati dallo stdout e attendere che il processo
    // termini.
    grepIn, _ := grepCmd.StdinPipe()
    grepOut, _ := grepCmd.StdoutPipe()
    grepCmd.Start()
    grepIn.Write([]byte("ciao grep\naddio grep"))
    grepIn.Close()
    grepBytes, _ := ioutil.ReadAll(grepOut)
    grepCmd.Wait()

    // Abbiamo tralasciato il controllo degli errori nelle
    // righe precedenti, ma ricordiamoci che possiamo usare
    // il classico pattern `if err != nil` per tutti loro.
    // Ci limitiamo inoltre a leggere dallo `StdoutPipe`,
    // ma potremmo leggere anche da `StderrPipe` nello stesso
    // identico modo.
    fmt.Println("> grep hello")
    fmt.Println(string(grepBytes))

    // Nota che per eseguire processi con flag e parametri
    // è necessario passare un array che li contenga, non è
    // possibile passare semplicemente una stringa.
    // Se si vuole eseguire un processo passando la sua stringa
    // è possibile utilizzare il comando `bash` con il
    // flag `-c` e passare la stringa del processo da eseguire.
    lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
    lsOut, err := lsCmd.Output()
    if err != nil {
        panic(err)
    }
    fmt.Println("> ls -a -l -h")
    fmt.Println(string(lsOut))
}
