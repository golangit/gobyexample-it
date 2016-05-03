# To experiment with command-line arguments it's best to
# build a binary with `go build` first.
# Per sperimentare con gli argomenti della linea di
# comando, è meglio creare un file binario con `go build`
# prima.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# Nel prossimo esempio vedremo come analizzare gli
# argomenti in un modo più avanzato con le flag.
