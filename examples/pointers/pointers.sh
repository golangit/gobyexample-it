# `zeroval` non cambia il valore di `i` in `main`,
# `zeroptr` invece sì perché ha un riferimento al valore
# nella memoria di quell'indirizzo.
$ go run pointers.go
iniziale:  1
zeroval:   1
zeroptr:   0
puntatore: 0x42131100
