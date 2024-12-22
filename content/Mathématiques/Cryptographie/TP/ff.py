import math, random
from data_exo5 import *

def isPrime(n, pl=[2])->bool:
    if n <=1 : return False
    for i in pl :
        if n % i == 0 : return False
    for i in range(pl[len(pl)-1], int(math.sqrt(n)+1)): 
        if n % i == 0 : return False
    return True

def listPrime(n):
    return [prime for prime in range(n) if isPrime(n)]

def euclide(a,b): # algo d'euclide étendu    
    u0, u1, v0, v1 = 1, 0, 0, 1
    q, r = divmod(a, b) 

    while r != 0 :
        a, b = b, r
        u2 = u0 - q*u1
        v2 = v0 - q*v1
        u0, v0, u1, v1 = u1, v1, u2, v2
        q, r = divmod(a, b) 
    
    return (b, u1, v1)

##############################################################################


# DOIT ÊTRE PREMIER SINON NE MARCHE PAS
dict_cod={'a':1,'b':2,'c':3,'d':4,'e':5,'f':6,'g':7,'h':8,'i':9,'j':10,'k':11,'l':12,'m':13,'n':14,'o':15,'p':16,'q':17,\
                   'r':18,'s':19,'t':20,'u':21,'v':22,'w':23,'x':24,'y':25,'z':26,' ':27,',':28,'.':29,'?':30, ':':0}

encoded_dict = dict_cod

def encode(text, a, b): # chiff_afine
    encodedVal = []
    encodedText = ""
    
    for i in text :
        encodedVal.append((a*dict_cod[i]+b))
    for i in encodedVal :
        for key, value in dict_cod.items() :
            # print(key, value, i, i%31)
            if value == i%len(dict_cod) :
                encodedText+=key

    return encodedText

def decode(mess, a, b):
    encodedVal = []
    decodedText = ""
    
    euclide_etendu = euclide(a, len(dict_cod))
    a = euclide_etendu[1] # coef de bezout associé

    for char in mess :
        # m = a**(-1) * (c-b) [32] # a-> coef de bezout
        charCode = a*(dict_cod[char]-b)%len(dict_cod)

        wasDecode = False
        for key, value in dict_cod.items() :
            # print(key, value, i, i%31)
            if value == charCode%len(dict_cod) :
                decodedText+=key
                wasDecode = True
        if not(wasDecode) :
            print("failed to decode", char, ":", charCode)

    return decodedText

##############################################################################


p = getP()
list_gen = getList_gen()


def isGenerator(p:int):
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
    if math.gcd(g, p) != 1 or math.gcd(y, p) != 1:
        raise ValueError("g et y doivent être dans (Z/pZ)*.")
    
    # Recherche brute pour x
    for x in range(p - 1):  # L'ordre de (Z/pZ)* est p-1
        if pow(g, x, p) == y:
            return x
    
    return None  # Si aucun x trouvé

##############################################################################

primes = listPrime(1000)
primes = primes[25:] # retire les 25 premiers qui sont trop petits

def key_creation():
    # Choisir deux nombres premiers distincts p et q
    p = random.choice(primes)
    q = random.choice([x for x in primes if x != p])
    
    # Calculer n et phi(n)
    n = p * q
    phi_n = (p - 1) * (q - 1)
    
    # Choisir un e tel que 1 < e < phi(n) et gcd(e, phi(n)) = 1
    e = random.choice([x for x in range(2, phi_n) if math.gcd(x, phi_n) == 1])
    
    # Calculer d tel que (d * e) % phi(n) = 1 (inverse modulaire)
    d = pow(e, -1, phi_n)
    
    # Retourner la clé publique (n, e) et la clé privée (d)
    return (n, e, d)

def encryption_int(n:int,e:int,m:int):
    """
    On note M le message que souhaite transmettre Bob.
    M est un entier naturel strictement inférieur à n. Le message chiffré
    et envoyé par Bob sera représenté par l’entier naturel C strictement
    inférieur à n et tel que : 

        C≡m**e [n]
    """
    if m>=n : return ValueError
    res = (m**e)%n
    if res>=n : return ValueError
    return res 

def decryption_int(n:int,d:int,c:int):
    """
    Pour déchiffrer le message C reçu, Alice utilise d, car
    on peut montrer qu’elle peut ainsi retrouver M :
    
        M≡c**d  [n]
    """
    return (c**d)%n