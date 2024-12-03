from exo1 import correction_list_prime as getPrimeUntil
from exo2 import bezout as euclide
from data_exo5 import *
import random
from math import gcd, sqrt

primes = getPrimeUntil(1000)

def key_creation():
    """
    FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF
    """
    n = primes[random.randint(26, len(primes)-1)]
    d = primes[random.randint(26, len(primes)-1)]
    while d == n : d = primes[random.randint(26, len(primes)-1)]
    e = random.randint(0, 1000)
    return n,e,d

def encryption_int(n,e,m):
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

def decryption_int(n,d,c):
    """
    Pour déchiffrer le message C reçu, Alice utilise d, car
    on peut montrer qu’elle peut ainsi retrouver M :
    
        M≡c**d  [n]
    """
    return (c**d)%n
    
if __name__ == "__main__":
    n,e,d = key_creation()
    print(n,e,d)
    m = 50
    en = encryption_int(n,e,m)
    de = decryption_int(n,d,en)
    print(m, "->", en, "->", de)
