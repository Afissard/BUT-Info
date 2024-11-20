---
title: TD1
draft: 
description: 
tags:
  - Algorithmes-Efficaces
  - Mathématiques
---
# Exo 0

| fonction           | 1 : "comparaison"            | 2 : ordre de grandeur | 3 : catégorie              |
| ------------------ | ---------------------------- | --------------------- | -------------------------- |
| f1(n)=n²+7         | f1 et f2 "à peu près pareil" | f1 ∈ θ(n²)            | quadratique (car n²)       |
| f2(u)=u(u+1)/2     | f1 et f2 "à peu près pareil" | f2 ∈ θ(n²)            | quadratique                |
| f3(n)=n⁴+n³(n-7)   | f3 est "la plus grande"      | f3 ∈ θ(n⁴)            | est de degré 4 (quartique) |
| f4(n)=c\*n\*log(n) | f3 est "la plus petite"      | f4 ∈ θ(n\*log(n))     | pseudo-linéaire            |
les math sans latex c'est compliqué à noté ( #TOOD)
4: 2 boucles imbriqués sur une donnée de taille n -> coût en n² 
5: ... 
6: +16k jours
7: ... 

# Exo 1
1. Il y a 2n+1 comparaison
2. c'est linéaire : θ(n)
3. *pour compter dans les "si", toujours prendre le pire des cas*

|                    | algo 1                  | algo 2                          |
| ------------------ | ----------------------- | ------------------------------- |
| nombre comparaison | 2\*n+1                  | 2n+1                            |
| nombre affectation | 3+n(1{aff-si} + 1{aff}) | 3+n(2 *dans le pire des cas*)+1 |
