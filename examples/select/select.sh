# Riceveremo i valori `"uno"` e `"due"`, come ci si
# potrebbe aspettare.
$ time go run select.go 
ricevuto uno
ricevuto due

# Nota che il tempo totale di esecuzione si aggira a
# 2 secondi piuttosto che a 3, perché i due `Sleep` sono
# eseguiti concorrenzialmente.
real	0m2.245s
