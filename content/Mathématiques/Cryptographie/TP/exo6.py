from exo1 import correction_list_prime as getPrimeUntil
from exo2 import bezout as euclide
from data_exo5 import *
import random
from math import gcd, sqrt
import hashlib as hash

primes = getPrimeUntil(1000)
primes = primes[25:] # retire les 25 premiers qui sont trop petits

def key_creation():
    # Choisir deux nombres premiers distincts p et q
    p = random.choice(primes)
    q = random.choice([x for x in primes if x != p])
    
    # Calculer n et phi(n)
    n = p * q
    phi_n = (p - 1) * (q - 1)
    
    # Choisir un e tel que 1 < e < phi(n) et gcd(e, phi(n)) = 1
    e = random.choice([x for x in range(2, phi_n) if gcd(x, phi_n) == 1])
    
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
    
def testRSA(msg:int):
    ok = True
    n,e,d = key_creation()
    en = encryption_int(n,e,msg)
    de = decryption_int(n,d,en)
    if msg != de :
        print(f"msg: {msg}\t-> {en}\t-> {de}\t=> {msg==de} | n:{n} e:{e}\td:{d}\t")
        ok = False
    else :
        print(f"{msg}\tpassed")
    return ok

def xTestsRAS(x:int):
    ok = True
    for i in range(x):
        test = testRSA(i)
        if test == False : ok = False
    return ok

if __name__ == "__main__":
    res = []
    res.append(xTestsRAS(100))
    res.append(xTestsRAS(200))
    print(res)
