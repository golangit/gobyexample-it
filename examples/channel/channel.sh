# Quando eseguiremo il programma potremmo vedere che
# `"ping"` è passato senza problemi da una goroutine
# all'altra attraverso il nostro channel.
$ go run channels.go 
ping

# Di default l'invio e il ricevimento si bloccano finché
# sia il mittente che il destinatario non sono pronti.
# Questo ci ha permesso di aspettare alla fine del 
# programma per il messaggio `"ping"` senza dover 
# effetturare alcuna sincronizzazione.
