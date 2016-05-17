// Nel precedente esempio abbiamo visto come [far partire
// processi esterni](eseguire-sottoprocessi). Facciamo ciò
// quando abbiamo bisogno di avere un processo esterno
// accessibile a un processo Go in esecuzione. A volte
// potremmo aver bisogno di rimpiazzare completamente
// il processo corrente con un altro (magari non scritto
// in Go). Per fare ciò utilizzaremo la implemntazione in
// Go della classica syscall
// <a href="https://en.wikipedia.org/wiki/Exec_(system_call)"><code>exec</code></a>.

package main

import "syscall"
import "os"
import "os/exec"

func main() {

    // Per il nostro esempio "execuiremo" `ls`. Go
    // richiede un percorso assoluto per il file binario
    // che vogliamo eseguire, quindi eseguiremo
    // `exec.LookPath` per trovarlo (probabilmente
    // `/bin/ls`.)
    binary, lookErr := exec.LookPath("ls")
    if lookErr != nil {
        panic(lookErr)
    }

    // `Exec` richiede che gli argomenti siano in una
    // slice (piuttosto che in una grande stringa).
    // Daremo a `ls` un po' di argomenti comuni. Nota che
    // il primo argomento dovrebbe essere il nome del
    // programma.
    args := []string{"ls", "-a", "-l", "-h"}

    // `Exec` ha anche bisogno di una serie di [variabili
    // d'ambiente](variabili-dambiente) da usare. Di
    // seguito forniamo semplicemente il nostro ambiente
    // corrente.
    env := os.Environ()

    // Qui facciamo finalmente la chiamata a `syscall.Exec`.
    // Se questa chiamata ha successo, l'esecuzione del
    // nostro processo terminerà qui e verrà sostituita dal
    // processo `ls -a -l -h`. Se c'è qualche errore
    // otterremo un valore restituito.
    execErr := syscall.Exec(binary, args, env)
    if execErr != nil {
        panic(execErr)
    }
}
