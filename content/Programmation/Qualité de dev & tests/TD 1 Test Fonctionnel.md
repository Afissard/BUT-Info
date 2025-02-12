---
title: TD 1 Test Fonctionnel
draft: 
description: 
tags:
---
![[Oracle-at-Delphi.png]]
# Partie 1 test fonctionnel
## Quels sont les états valides d’une instance d’AlarmClock
DT (ring, hour, min), l'oracle contrôle les 5 attributs ou la levé d’exception
`is_enable => is_ringing`
## Concevez les tests fonctionnels du constructeur
### Identification des variables qui formes chaque données de tests

| nom        | type    | interval (nominal) |
| ---------- | ------- | ------------------ |
| ring       | int     | \[1;10]            |
| hour       | int     | \[0;23]            |
| min        | int     | \[0;59]            |
| is_enable  | boolean |                    |
| is_ringing | boolean |                    |
### Analyse partitionnelle

| nom  | type                  | interval nominal | valeurs exeptionnel | partition fonctionnel |
| ---- | --------------------- | ---------------- | ------------------- | --------------------- |
| ring | `int [2^-31, 2^31-1]` | `[1, 10]`        |                     | void                  |
| hour | `int [2^-31, 2^31-1]` | `[0, 23]`        |                     | void                  |
| min  | `int [2^-31, 2^31-1]` | `[0, 59]`        |                     | void                  |
### Table de décision

| Donnée de test       |     |     |     |     |     |     |     |
| -------------------- | --- | --- | --- | --- | --- | --- | --- |
| ring `[2^-31, 1[`    | X   |     |     |     |     |     |     |
| ring `[1, 10]`       |     |     |     |     |     |     | X   |
| ring `]10, 2^31-1]`  |     | X   |     |     |     |     |     |
| hour `[2^-31, 1[`    |     |     | X   |     |     |     |     |
| hour `[0, 23]`       |     |     |     |     |     |     | X   |
| hour `]10, 2^31-1]`  |     |     |     | X   |     |     |     |
| min `[2^-31, 1[`     |     |     |     |     | X   |     |     |
| min `[0, 59]`        |     |     |     |     |     |     | X   |
| min `]10, 2^31-1]`   |     |     |     |     |     | X   |     |
| Contrôle de l'oracle |     |     |     |     |     |     |     |
| ring `[1, 10]`       |     |     |     |     |     |     | X   |
| hour `[0, 23]`       |     |     |     |     |     |     | X   |
| min `[0, 59]`        |     |     |     |     |     |     | X   |
| is_ringing `F`       |     |     |     |     |     |     | X   |
| is_enable `F`        |     |     |     |     |     |     | X   |
| AlarmClockExeption   | X   | X   | X   | X   | X   | X   |     |
### Cas de test
1. ((-5, 17, 50); ACE)
2. ((50, 17, 50); ACE)
3. ((5, -17, 50); ACE)
4. ((5, 170, 50); ACE)
5. ((5, 17, -50); ACE)
6. ((5, 17, 500); ACE)
7. ((5, 17, 50); (5, 17, 50, F, F))
## Concevez les tests fonctionnels de la fonction selectRing()
### Identification des variables qui formes chaque données de tests
| nom        | type    | interval (nominal) |
| ---------- | ------- | ------------------ |
| ring       | int     | `[1, 10]`          |
| ringtone   | int     | `[1, 10]`          |
| is_ringing | boolean |                    |
### Analyse partitionnelle
| nom      | type                  | interval nominal | valeurs exeptionnel | partition fonctionnel |
| -------- | --------------------- | ---------------- | ------------------- | --------------------- |
| ring     | `int [2^-31, 2^31-1]` | `[1, 10]`        |                     | void                  |
| ringtone | `int [2^-31, 2^31-1]` | `[1, 10]`        | ringtone != ring    | void                  |
### Table de décision

| Donnée de test          |     |     |     |     |     |     |
| ----------------------- | --- | --- | --- | --- | --- | --- |
| is_ringing `T`          | X   | X   |     |     |     |     |
| is_ringing `F`          |     |     | X   | X   | X   | X   |
| ring `[1, 10]`          | X   |     |     |     | X   |     |
| ring `[1, 10]`          |     | X   |     |     |     |     |
| ringtone `[2^-31, 1[`   |     |     | X   |     |     |     |
| ringtone `[1, 10]`      | X   | X   |     |     | X   | X   |
| ringtone `]10, 2^31-1]` |     |     |     | X   |     |     |
| Contrôle de l'oracle    |     |     |     |     |     |     |
| ring <- ringtone        | X   | X   |     |     | X   | X   |
| changement `T`          |     |     |     |     |     |     |
| changement `F`          | X   | X   |     |     |     |     |
| AlarmClockExeption      |     |     | X   | X   |     |     |
### Cas de test
1. (T, 1, 5), (1, F)
2. (T, 5, 5), (5, F)
3. (F, 1, -5), (ACE)
4. (F, 1, 50), (ACE)
5. (F, 1, 5), (5, T)
6. (F, 5, 5), (5, F)
## Concevez les tests fonctionnels de la fonction checkTimeAndRing()
spécificité si allumé alors que le réveil doit sonner -> sonne et est éteint si nécessaire
### Identification des variables qui formes chaque données de tests
is_enable
hour
min
### Analyse partitionnelle
### Table de décision
bla bla, big table
### Cas de test
1. (F, 12, 2O, F), (in-changer:F, F)
2. (T, 17, 00), (F, T)
3. (T, 12, 00), (F, T)
4. (T, 17, 20), (F, T)
5. (T, 12, 20), (T, F)
