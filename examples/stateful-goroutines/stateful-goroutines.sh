# Eseguendo il nostro codice possiamo vedere che
# sono state eseguite circa 800.000 operazioni in
# un secondo.
$ go run stateful-goroutines.go
ops: 807434

# Per questo caso particolare l'approccio a goroutine 
# risulta migliore rispetto a quello basato su mutex.
# Questo approccio risulta vincente nel caso in cui si
# debbano gestire svariati channel oppure nei casi in
# cui la gestione di troppe mutex potrebbe risultare
# complesso. Il consiglio è quello di utilizzare
# l'approccio che risulta più naturale per ogni singolo
# caso e che mantiene alta la leggibilità del proprio
# codice.