from exo1 import correction_list_prime as getPrimeUntil
from exo2 import bezout as euclide
from data_exo5 import *
import random
from math import gcd

# démonstration td : pour tout g E[2;p-1], g^(p-1) === 1[p]
# slide cour petit theorème de Fermat : 1 ≡a^(p−1)[p]

p = getP()
list_gen = getList_gen()


def trouveGenerateur(p:int):
    """
    Ecrire une fonction renvoyant, pour un nombre premier p passé en argument,
    un générateur de (Z/pZ)* choisi aléatoirement (choisir un nombre aléatoirement et tester
    s’il est générateur et reproduire l’opération jusqu’à ce que ce soit le cas)
    """
    
    generateur = None
    isNotGen = True
    testedVal = []
    
    while len(testedVal) < p-1-2:
        candidat = random.randint(2, p - 1)
        if not(candidat in testedVal) :
            testedVal.append(candidat)
            classes = set(range(candidat))
            engendre = set((k * p) % candidat for k in range(candidat))

            if not(classes == engendre) : 
                isNotGen = False
                print(p,"\t",candidat,"\t",testedVal)
                return candidat
    
    print(p,"\t",generateur,"\t",testedVal)
    return generateur

def logDiscret(g, y, p):
    """
    Ecrire une fonction log_discret(y,g,p) pour un élément y de (Z/pZ)* et un
    générateur g de (Z/pZ)*, l’élément x tel que gx ≡y[p].
    
    Tests : 
    3,6,7 -> 3
    3,3,7 -> 4
    """
    # Vérification des entrées
    if gcd(g, p) != 1 or gcd(y, p) != 1:
        raise ValueError("g et y doivent être dans (Z/pZ)*.")
    
    # Recherche brute pour x
    for x in range(p - 1):  # L'ordre de (Z/pZ)* est p-1
        if pow(g, x, p) == y:
            return x
    
    return None  # Si aucun x trouvé

def Shanks():
    """
    Ecrire une fonction, avec l’algorithme de Shanks des pas de bébé, pas de géant
    - Soit s = 1 + ⌊√p⌋
    - Calculer g**(-s)
    - Créer deux liste
        - L1 : 1, g, g2, g3,··· ,g**(s-1)
        - L2 : y, y.g−s,, y.g−2s,, y.g−3s,··· ,, y.g−(s−1)×s
    - Trouver une occurrence commune gr0 et y.g−k0s respectivement dans les listes L1 et L2
    - x = r0 + k0.s est une solution
    """
    return None

def testLogDiscret():
    res = logDiscret(3,6,7)
    if res == 3: print("3,6,7 -> 3")
    else : print("error",res)
    res = logDiscret(3,4,7)
    if res == 4: print("3,4,7 -> 4")
    else : print("error",res)


if __name__ == "__main__" :
    trouveGenerateur(7)
    trouveGenerateur(13)
    trouveGenerateur(31)
    trouveGenerateur(100)
    trouveGenerateur(10)
    trouveGenerateur(9)

    testLogDiscret()