# Il nostro programma in eseguzione mostra i 9 task
# che vengono eseguiti dai vari worker. Il programma
# termina in circa 3 secondi, anche se il totale dei
# task avrebbe richiesto 9 secondi. Questo perchè ci
# sono 3 worker che eseguono in parallelo
$ time go run worker-pools.go 
worker 1 sto eseguendo il task job 1
worker 2 sto eseguendo il task job 2
worker 3 sto eseguendo il task job 3
worker 1 sto eseguendo il task job 4
worker 2 sto eseguendo il task job 5
worker 3 sto eseguendo il task job 6
worker 1 sto eseguendo il task job 7
worker 2 sto eseguendo il task job 8
worker 3 sto eseguendo il task job 9

real	0m3.149s
