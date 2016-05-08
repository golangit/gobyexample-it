# Se si esegue questo programma si potrà notare come 
# la prima operazione andrà in timeout, mentre la seconda
# verrà eseguita correttamente.
$ go run timeouts.go 
timeout 1
risultato 2

# Usare questo pattern con il `select` per implementare
# i timeout e la comunicazione attraverso i channel 
# risulta essere particolarmente vantaggioso. 
# Svariate altre funzionalità di Go sono basate su 
# channel e `select`, vedremo due esempi di questo a 
# breve: i _timer_ e i _ticker_.
