# Eseguendo il programma verrà stampato il valore di
# `FOO`, impostato da noi nel sorgente, mentre il valore
# di `BAR` sarà vuoto.
$ go run environment-variables.go
FOO: 1
BAR: 

# La lista di chiavi nell'ambiente dipenderà dal tuo
# sistema
TERM_PROGRAM
PATH
SHELL
...

# Se prima di eseguire il programma assegnamo un valore
# a `BAR`, allora il programma lo stamperà.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
