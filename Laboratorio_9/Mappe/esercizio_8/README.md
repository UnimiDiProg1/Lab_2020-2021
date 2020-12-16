# Estratto conto

Scrivere un programma che, tramite **redirizione dello standard input**,  legga un file di testo `operazioni.txt` memorizzato nella stessa directory in cui è memorizzato il programma.

La prima riga del file `operazioni.txt` è una stringa nel formato

`IMPORTO`

dove `IMPORTO` è un valore reale che specifica il saldo iniziale di un conto corrente.

Ogni riga del file `operazioni.txt` successiva alla prima è una stringa nel formato

`DATA;TIPO;IMPORTO`

e specifica un'operazione bancaria effettuata sul conto corrente:

* `DATA`: Una stringa senza spazi che codifica la data in cui è avvenuta l'operazione bancaria, specificata in uno dei seguenti possibili formati:
    1. aaaa/m/g
    2. aaaa/m/gg
    3. aaaa/mm/g
    4. aaaa/mm/gg
    5. g/m/aaaa
    6. gg/m/aaaa
    7. g/mm/aaaa
    8. gg/mm/aaaa
* `TIPO`: Un carattere che specifica la natura dell'operazione bancaria: `'V'` per versamento e `'P'` per prelievo.
* `IMPORTO`: Un valore reale maggiore di 0 che specifica l'ammontare dell'importo versato o prelevato.

Come mostrato nell'**Esempio di esecuzione**, il programma deve stampare a video:
1. Il saldo iniziale del conto corrente.
2. Per ogni giorno in cui è stata effettuata almeno un'operazione bancaria:
    1. la data del giorno considerato nel formato gg-mm-aaaa
    2. la lista delle operazioni effettuate nel giorno considerato;
    3. il saldo giornaliero: la differenza tra l'ammontare degli importi versati e l'ammontare degli importi prelevati con operazioni effettuate nel giorno considerato.
3. Il saldo finale del conto corrente.
 
In particolare, per quanto riguarda il punto 2., si noti che le informazioni relative ai giorni in cui è stata effettuata almeno un'operazione bancaria devono essere stampate considerando l'ordine cronologico dei giorni (dal giorno relativo alla data più antica al giorno relativo alla data più recente).

Si assuma che:
* ogni riga del file `operazioni.txt` sia nel formato corretto;
* i valori presenti in ogni riga del file `operazioni.txt` successiva alla prima specifichino correttamente un'operazione bancaria;
* il file `operazioni.txt` non sia vuoto.

##### Esempio d'esecuzione:

```text
$ cat operazioni_1.txt 
12.5
2019/06/04;P;34
2020/01/05;V;45.34
2019/6/04;P;10
2019/6/24;P;23
2019/06/04;P;456.10
2019/06/4;V;26.1
2019/06/8;P;74.23
2020/1/05;V;178.30
2020/01/7;V;194
2020/02/05;V;56

$ go run estratto_conto.go < operazioni_1.txt
SALDO INIZIALE:                12.50

DATA: 04-06-2019
Prelievo:                      34.00
Prelievo:                      10.00
Prelievo:                     456.10
Versamento:                    26.10
                         ___________
SALDO GIORNALIERO:           -474.00
                         ===========

DATA: 08-06-2019
Prelievo:                      74.23
                         ___________
SALDO GIORNALIERO:            -74.23
                         ===========

DATA: 24-06-2019
Prelievo:                      23.00
                         ___________
SALDO GIORNALIERO:            -23.00
                         ===========

DATA: 05-01-2020
Versamento:                    45.34
Versamento:                   178.30
                         ___________
SALDO GIORNALIERO:            223.64
                         ===========

DATA: 07-01-2020
Versamento:                   194.00
                         ___________
SALDO GIORNALIERO:            194.00
                         ===========

DATA: 05-02-2020
Versamento:                    56.00
                         ___________
SALDO GIORNALIERO:             56.00
                         ===========

SALDO FINALE:                 -85.09
                         ===========

$ cat operazioni_2.txt
1255.55
2019/10/2;V;130.11
2019/10/3;P;10.2
2019/11/01;V;560.9
2019/12/01;P;126
2019/10/02;P;30
2019/10/03;V;110.49
2019/11/1;P;560
2019/12/1;V;120.11

$ go run estratto_conto.go < operazioni_2.txt
SALDO INIZIALE:              1255.55

DATA: 02-10-2019
Versamento:                   130.11
Prelievo:                      30.00
                         ___________
SALDO GIORNALIERO:            100.11
                         ===========

DATA: 03-10-2019
Prelievo:                      10.20
Versamento:                   110.49
                         ___________
SALDO GIORNALIERO:            100.29
                         ===========

DATA: 01-11-2019
Versamento:                   560.90
Prelievo:                     560.00
                         ___________
SALDO GIORNALIERO:              0.90
                         ===========

DATA: 01-12-2019
Prelievo:                     126.00
Versamento:                   120.11
                         ___________
SALDO GIORNALIERO:             -5.89
                         ===========

SALDO FINALE:                1450.96
                         ===========
```