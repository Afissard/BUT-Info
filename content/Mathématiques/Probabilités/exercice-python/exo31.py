import math, numpy as np, matplotlib.pyplot as plt, random

"""
Simulons le tirage successif de quatre boules avec remise
dans une urne contenant 7 boules blanches et 3 noires. Par 
la simulation d'un grand nombre de cette expérience et en 
renvoyant une certaine fréquence observé retrouvé des 
valeurs approximatives des probabilités des événement 
suivants :
    - Tirage contenant exactement deux boules blanches
    - Tirage contenant au moins une boule blanche
    
Déterminer les probabilité associé aux deux événements étudiés.
---

X~B(4, 0.7)
P(x=2)=C(2, 4)*(0.7**2)*(1-0.7)

Y = {1 si X=2 ; 0 sinon}

def tirage(n, nbr_blanches):
    print("todo")

Yn -> E[Y] # Yn barre tend vers E[Y]
"""

###################################################

def tirage_successif(urne, nb=4):
    main = []
    for _ in range(nb):
        boule = random.randint(0, len(urne)-1)
        main.append(urne[boule])
    return main

def simulation():
    urne = ["b", "b", "b", "b", "b", "b", "b", "n", "n", "n"]
    mUrne = [1, 1, 1, 1, 1, 1, 1, 0, 0, 0]
    print(tirage_successif(urne=mUrne))

###################################################

def tirage(n, nbr_blanches):
    res=[]
    for _ in range(n):
        x = tirage_successif([1, 1, 1, 1, 1, 1, 1, 0, 0, 0], 4)
        res.append(np.count_nonzero(x))
        
    return np.mean(res)

###################################################

def tirage_corriger1(n, nbb):
    res = np.zeros((n, 4))
    count = np.zeros(n)
    for i in range(n):
        res[i] = np.random.randint(1, 11, 4)
        count[i] = np.count_nonzero(res[i]<8)
    return np.count_nonzero(count==nbb)/n


def tirage_corriger2(n, nbb):
    simu = np.random.rand(n, 4)
    test = np.count_nonzero(simu<0.7, axis=1)
    # print(test)
    return np.count_nonzero(test==nbb)/n
    
if __name__ == "__main__":
    # print(tirage(100, 2))
    
    print(tirage_corriger1(100,2))
    print(1-tirage_corriger1(100,2))
    
    print(tirage_corriger2(100,2))
    print(1-tirage_corriger2(100,2))
