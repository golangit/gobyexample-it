# Per provare il line filter creaiamo prima un file con 
# alcune linee in minuscolo.
$ echo 'ciao'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# Utilizziamo il line filter per ottenere le stringhe
# in maiuscolo.
$ cat /tmp/lines | go run line-filters.go
CIAO
FILTER
