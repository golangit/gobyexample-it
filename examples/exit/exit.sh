# Se esegui `exit.go` usando `go run`, il valore di uscita
# verrà intercettato da Go e stampato.
$ go run exit.go
exit status 3

# Compilando e eseguendo il programma potrai anche vedere
# il valore restituito nel terminale.
$ go build exit.go
$ ./exit
$ echo $?
3

# Nota che il `!` del nostro programma non è mai stato
# stampato.
