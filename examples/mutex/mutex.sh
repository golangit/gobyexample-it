# Eseguendo il nostro programma vedremo che abbiamo
# eseguito circa 4,000,000 di operazioni sul nostro
# `stato` sincronizzato da `mutex`.
$ go run mutexes.go
ops: 4272750
stato: map[4:29 3:50 0:39 1:71 2:51]

# Come prossima cosa, vedremo come come implementare
# questa gestione della memoria condivisa usando
# esclusivamente goroutine e channel.
