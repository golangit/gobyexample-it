# Quando eseguiamo il nostro processo verrà
# rimpiazzato da `ls`
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Nota che Go non offre la classica funzionalità di UNIX
# `fork`. Di solito, però, questo non è un problema,
# poiché l'uso di goroutine, lo spawn di processi
# e l'uso di `exec` coprono la maggior parte degli usi
# di `fork`.
