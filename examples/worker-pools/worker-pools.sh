# Il nostro programma in esecuzione mostra i 9 task
# che vengono eseguiti dai vari worker. Il programma
# termina in circa 3 secondi, anche se il totale dei
# task avrebbe richiesto 9 secondi. Questo perchè ci
# sono 3 worker che vengono eseguiti in parallelo
$ time go run worker-pools.go 
worker 1 sta eseguendo task job 1
worker 2 sta eseguendo task job 2
worker 3 sta eseguendo task job 3
worker 1 sta eseguendo task job 4
worker 2 sta eseguendo task job 5
worker 3 sta eseguendo task job 6
worker 1 sta eseguendo task job 7
worker 2 sta eseguendo task job 8
worker 3 sta eseguendo task job 9

real	0m3.149s
