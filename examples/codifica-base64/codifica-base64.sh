# La stringa viene codificata in un valore differente 
# se si utilizza il formato URL-compatible.
# Entrambi i valori però decodificano alla stessa stringa
$ go run base64-encoding.go
YWJjMTIzIT8kKiYoKSctPUB+
abc123!?$*&()'-=@~

YWJjMTIzIT8kKiYoKSctPUB-
abc123!?$*&()'-=@~
