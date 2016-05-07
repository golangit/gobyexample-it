# Eseguendo il programma si nota che viene calcolato
# l'hash della funzione e stampato in esadecimale 
# a schermo
$ go run sha1-hashes.go
sha1 this string
cf23df2207d99a74fbe169e3eba035e633b65d94

# È possibile calcolare altre funzioni hash in un modo
# molto simile. Ad esempio per calcolare l'MD5 è 
# sufficiente importare `crypto/md5` ed usare
# `md5.New()`.

# Nota che se necessiti di funzioni hash che siano sicure
# dal punto di vista crittografico, ti consigliamo di 
# approfondire il tema dello [hash strength](http://en.wikipedia.org/wiki/Cryptographic_hash_function)!
