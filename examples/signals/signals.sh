# Quando eseguiamo questo programma si bloccherà in 
# attesa di un segnale. Premendo `ctrl-C` (che il 
# terminale mostrerà come `^C`) inviamo un segnale 
# di tipo `SIGINT`, causando la scrittura di 
# `interrupt` e la successiva terminazione del 
# programma.
$ go run signals.go
awaiting signal
^C
interrupt
exiting
