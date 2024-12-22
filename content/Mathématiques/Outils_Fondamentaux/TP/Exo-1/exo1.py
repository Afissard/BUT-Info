import math

def divEuclide(a:int, b:int) :
    """
    Programmer l'algorithme prenant en arguments deux entiers 
    naturels a et b et renvoyant le couple formé du quotient 
    et du reste de la division euclidienne de a par b (sans 
    faire appel aux opérateurs déjà intégrés dans Python 
    pour le renvoi de ces résultats)
    """
    quotient, reste = 0, 0

    if not(a>b) : return SyntaxError

    while a>=b :
        quotient +=1
        if a - b >= 0:
            a-=b
    reste = a
    return quotient, reste

def algoEuclide(a:int, b:int):
    """
    Programmer l'algorithme d'Euclide étendu prenant en arguments 
    deux entiers naturels a et b et renvoyant le triplet formé du 
    pgcd 'd' de ces deux entiers et de deux entiers 'u' et 'v' 
    vérifiant : a*u + b*v = d.
    """
    if not(a>b) : return SyntaxError

    # initialisation
    u0, u1, v0, v1 = 1, 0, 0, 1
    quotient, reste = divEuclide(a, b)

    # boucle algo
    while reste != 0 :
        a, b = b, reste
        u2 = u0 - quotient*u1
        v2 = v0 - quotient*v1
        u0, v0 = u1, v1
        u1, v1 = u2, v2
        quotient, reste = divEuclide(a, b)
    
    # fin algo
    pgcd = b
    coefBézout = (u1, v1)
    return pgcd, coefBézout


def main():
    print(divEuclide(4,2))
    a, b = 210, 55
    test = algoEuclide(a, b)
    print(test)
    verifyPGCD = a*test[1][0] + b*test[1][1] 
    if verifyPGCD == test[0]:
        print("ok")
    else :
        print("nop : a*u + b*v =",verifyPGCD, "mais !=", test[0])

if __name__ == "__main__":
    main()