# Per fare qualche prova con la linea di comando è meglio
# prima compilare il programma ed eseguire direttamente
# il compilato.
$ go build command-line-flags.go

# Proviamo ad eseguire passando tutti i flag
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# Nota che se non si settano i flag, vengono caricati
# i valori di default
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# I parametri posizionali possono essere passati in
# fondo dopo tutti i flag
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# Nota che il pacchetto `flag` richiede che tutti i flag
# siano possizionati necessarimente prima dei parametri.
# Eventuali flag che si trovano in coda vengono trattati
# come parametri e non riconosciuti correttamente
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
trailing: [a1 a2 a3 -numb=7]

# Usa `-h` o `--help` per ottenere i suggerimenti 
# automatici sull'utilizzo dei flag (che vengono 
# generati a partire dalle definizioni che mettiamo
# nel programma).
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# Se si utilizza un flag che non è stato dichiarato,
# il programma mostrerà un messaggio di errore e stamperà
# di nuovo il messaggio di informazione.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# Vedremo adesso le variabili d'ambiente. Un altro modo
# comune per rendere parametrici i programmi.
