# Prova ad eseguire il codice per la scrittura su file.
$ go run writing-files.go 
scritti 4 byte
scritti 6 byte
scritti 13 byte

# ... poi controlla i contenuti dei file scritti.
$ cat /tmp/dat1
ciao
go
$ cat /tmp/dat2
dei
write
bufferizzati

# Nel prossimo esempio vedremo come applicare i concetti
# che abbiamo appena visto riguardanti l'I/O agli stream
# `stdin` e `stdout`.
