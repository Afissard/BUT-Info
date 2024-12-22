# |
Exécute la commande 2 après la commande 1.
```bash
echo > texte.txt "hello world!" | cat texte.txt
```
# &&
Exécute une la commande si la précédente à réussi
```bash
echo hello && echo world
```
# ||
Exécute la commande 1, si elle échoue, la commande 2 est exécuté
```bash
cat texte.txt || echo oups
```
# Combinaison 
```bash
echo > texte.txt "hello world!" | cat texte.txt && rm -f texte.txt || echo "Error !"
```