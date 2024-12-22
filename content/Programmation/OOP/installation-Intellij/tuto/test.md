# Exécuter des cas de tests

Notez la présence du dossier `src/test/kotlin/` qui pourra contenir des cas de tests.

Placez vous dans `Calcul.kt` et selectionnez une des deux fonctions. Selectionnez Code > Generate... > Test...

![](junit.png)


Selectionnez les deux méthodes `sum(...)` et `product(...)`, puis `OK`.

Un nouveau fichier nommé `CalculKtTest.kt` est apparu dans `src/test/kotlin/`. Modifiez le comme suit 

    @Test
    fun sum() {
        assertEquals(42, sum(40,1))
    }

    @Test
    fun product() {
        assertEquals(42, product(21, 2))
    }

__ATTENTION :__ l'erreur pour `sum(40,1)` est volontaire.


![](junit_run.png)

Exécutez tous les cas de tests, en appuyant sur la "double flèche verte" ou en via un "clic droit" sur `CalculKtTest.kt`, puis `Run 'CalculKtTest'`.

Vous pouvez aussi exécuter un seul cas de test en appuyant sur la "flèche verte".

![](junit_result.png)

Le résultats de l'exécution de cas de tests sont donnés dans un onglet spécifique.
- En selectionnant un cas de tests en échec, le détail de l'erreur s'affiche ;
- On peut afficher ou non les cas de tests ok, trier les cas de tests ;
- Re-exécuter seulement les cas de tests en échec.


[Deboguer un programme](debug.md)
