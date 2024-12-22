## table employé
| Nuempl | Nomempl | hebdo | salaire | affect |
| ------ | ------- | ----- | ------- | ------ |
| 20     | Marcel  | 35    | 2000    | 3      |
| 23     | Claude  | 20    | 2500    | 3      |
| 37     | Michèle | 35    | 3000    | 3      |
| 39     | Léon    | 35    | 1900    | 1      |
| 41     | Jules   | 35    | 2800    | 1      |
| 30     | Edith   | 30    | 4000    | 4      |
| 17     | Sophie  | 35    | 2800    | 2      |
| 57     | Anne    | 35    | 1300    | 2      |
| 68     | Casimir | 20    | 3000    | 4      |
| 10     | Martin  | null  | 1000    | 4      |
## table service
| Nuserv | Nomserv      | Chef |
| ------ | ------------ | ---- |
| 1      | achat        | 41   |
| 2      | vente        | 17   |
| 3      | informatique | 23   |
| 4      | mécanique    | 20   | 
## table projet
| Nuproj | Nomproj | resp |
| ------ | ------- | ---- |
| 103    | Cobra   | 30   |
| 237    | Zorro   | 30   |
| 370    | Erasmus | 57   |
| 492    | Commet  | 20   |
| 135    | Eureka  | 57   |
| 160    | Esprit  | 5     |
## table travail
| Nuempl | Nuproj | Durée |
| ------ | ------ | ----- |
| 17     | 135    | 5     |
| 20     | 492    | 10    |
| 23     | 237    | 15    |
| 23     | 135    | 10    |
| 30     | 103    | 5     |
| 30     | 370    | 5     |
| 30     | 492    | 5     |
| 30     | 135    | 5     |
| 30     | 160    | 5     |
| 30     | 237    | 10    |
| 37     | 160    | 30    |
| 37     | 135    | 5     |
| 39     | 237    | 10    |
| 39     | 135    | 15    |
| 41     | 103    | 20    |
| 41     | 492    | 20    |
| 57     | 103    | 20    |
| 57     | 370    | 10    |
| 68     | 370    | 25    | 
## 1 : Employé {NUEMPL, NOMEMPL}
| Nuempl | Nomempl |
| ------ | ------- |
| 20     | Marcel  |
| 23     | Claude  |
| 37     | Michèle |
| 39     | Léon    |
| 41     | Jules   |
| 30     | Edith   |
| 17     | Sophie  |
| 57     | Anne    |
| 68     | Casimir |
| 10     | Martin  |
## 2 : Employé \[Salaire > 2000 \]
| Nuempl | Nomempl | salaire |
| ------ | ------- | ------- |
| 23     | Claude  | 2500    |
| 37     | Michèle | 3000    |
| 41     | Jules   | 2800    |
| 30     | Edith   | 4000    |
| 17     | Sophie  | 2800    |
| 68     | Casimir | 3000    |
## 3 : Union(Employé, Travail)
| Nuempl | Nomempl | Nuproj | Durée |
| ------ | ------- | ------ | ----- |
| 17     | Sophie  | 135    | 5     |
| 23     | Claude  | 237    | 15    |
| 23     | Claude  | 135    | 10    |
| 30     | Edith   | 103    | 5     |
| 30     | Edith   | 370    | 5     |
| 30     | Edith   | 492    | 5     |
| 30     | Edith   | 135    | 5     |
| 30     | Edith   | 160    | 5     |
| 30     | Edith   | 237    | 10    |
| 37     | Michèle | 160    | 30    |
| 37     | Michèle | 135    | 5     |
| 41     | Jules   | 103    | 20    |
| 41     | Jules   | 492    | 20    |
| 68     | Casimir | 370    | 25    |
## 4 : Inter(Employé {nuempl },Travail {nuempl })
| Nuempl |
| ------ |
| 17     |
| 23     |
| 23     |
| 30     |
| 30     |
| 30     |
| 30     |
| 30     |
| 30     |
| 37     |
| 37     |
| 41     |
| 41     |
| 68     |
## 5 : Diff(Employé {nuempl },Travail {nuempl })
| Nuempl |
| ------ |
| 10     |
## 6 : Div(Travail {nuempl, nuproj }, Projet {nuproj })
????
Résultat => nmpl : 30
