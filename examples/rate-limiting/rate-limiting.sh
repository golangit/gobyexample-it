# Eseguendo il nostro programma vedremo il gruppo
# di richieste che verranno gestite ognuna ogni 200
# millisecondi.
$ go run rate-limiting.go
richiesta 1 2016-04-17 19:05:37.736132953 +0200 CEST
richiesta 2 2016-04-17 19:05:37.936138961 +0200 CEST
richiesta 3 2016-04-17 19:05:38.136209943 +0200 CEST
richiesta 4 2016-04-17 19:05:38.336145582 +0200 CEST
richiesta 5 2016-04-17 19:05:38.536120745 +0200 CEST

# Invece, per il nostro secondo gruppo di richieste
# vedremo che le prime tre vengono eseguite
# all'istante, mentre le altre 2 a distanza di 200
# millisecondi l'una dall'altra.
richiesta 1 2016-04-17 19:05:38.536251527 +0200 CEST
richiesta 2 2016-04-17 19:05:38.536266185 +0200 CEST
richiesta 3 2016-04-17 19:05:38.536277665 +0200 CEST
richiesta 4 2016-04-17 19:05:38.736385724 +0200 CEST
richiesta 5 2016-04-17 19:05:38.936385957 +0200 CEST
