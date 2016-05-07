[Readme in Inglese](https://github.com/golangit/gobyexample-it/blob/master/README.md)

## Go by Example - IT (gobyexample.it)

Versione italiana di [Go by Example](https://gobyexample.com),
un sito che insegna a programmare in Go utilizzando
degli esempi pratici commentati.

La versione italiana di Go by Example può essere consultata su 
[http://gobyexample.it](http://gobyexample.it).

### Traduzione

La traduzione italiana è stata realizzata dalla [community italiana di golang (golang-it)](http://golangit.github.io).
Sentiti libero di [aprire un bug](https://github.com/golangit/gobyexample-it/issues) oppure una [pull request](https://github.com/golangit/gobyexample-it/pulls) per riportare eventuali problemi od imprecisioni nella traduzione.

### Overview

Il sito di Go by Example viene generato estraendo il codice
e i commenti dai file contenuti nella cartella `examples`.
Questi contenuti vengono composti utilizzando i file .tmpl
nella cartella `templates` per generare dei files statici
(cartella `public`).
Il programma che implementa il build process si trova dentro
la cartella `tools`, mentre nella cartella `vendor` si trovano
le dipendenze necessarie per buildare il sito.

La cartella `public` può essere servita tramite un sistema
per la condivisione di contenuti statici. Il progetto originale
utilizzava S3 e CloudFront.

### Compilazione

Per compilare il progetto è necessario avere Go e Python installati.
Invoca i comandi:

```console
$ go get github.com/russross/blackfriday
$ tools/build
$ open public/index.html
```

Per compilare in modo continuo in loop esegui:

```console
$ tools/build-loop
```

### Licenza

Traduzione italiana: copyright Nicola Corti e Morgan Bazalgette, licenza
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Progetto originale: copyright Mark McGranaghan, licenza
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).

Go Gopher: copyright [Renée French](http://reneefrench.blogspot.com/), licenza
[Creative Commons Attribution 3.0 Unported License](http://creativecommons.org/licenses/by/3.0/).
