Soit E un ensemble et *R* une relation sur E. On dit que *R* est une **relation d'ordre** si elle est : 
- [[Réflexive]]
- [[Transitive]]
- [[Anti-Symétrique]]
Lorsque E possède une relation d'ordre, on dit qu'il est ordonné.

# Remarques
- La notion d'ordre de croissance n'est définie que sur des ensembles ordonnées.
- Les notions de minimum et de maximum n'existe que sur des ensemble ordonnés.
## Ordre large (≼/≽) vs Ordre strict (≺/≻)
Attention la definition précédente d'une relation d'ordre correspond à un ordre large (≼/≽). Une relation binaire (≺/≻) est un ordre strict quand elle est [[Irreflexive]] et [[Transitive]] (et on peut alors montrer qu'un ordre strict est toujours [[Anti-Symétrique]] car on ne peut avoir : *(x ≺ y)∧(y ≺ x)*).
## Ordre total vs Ordre partiel
Un ordre (≼/≽) est dit total si deux éléments sont toujours comparables : *∀x, y ∈ E, (x ≼ y) ∨ (y ≼ x)*
Un ordre qui n'est pas total est dit partiel.
# Ordre : minorant / majorant
## Minorant
On dit que x ∈ E est un minorant de F si tout élément de F est plus grand que x pour : *∀y ∈ F, x ≼ y*. Si un minorant de F est un élément de F on dit que c'est le plus petit élément de F.
## Majorant
On dit que x ∈ E est un majorant de F si tout élément de F est plus petit que x pour : *∀y ∈ F, y ≼ x*. Si un majorant de F est un élément de F on dit que c'est le plus grand élément de F.
# Ordre : minimal / maximal
## Minimal
x ∈ F est un élément minimal de F quand aucun élément de F n'est strictement plus petit : *∀y ∈ F, y ≼ x ⇒ y = x*
## Maximal
x ∈ F est un élément maximal de F quand aucun élément de F n'est strictement plus grand : *∀y ∈ F, x ≼ y ⇒ y = x*