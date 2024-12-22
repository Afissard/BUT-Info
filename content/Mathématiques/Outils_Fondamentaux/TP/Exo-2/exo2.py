import math

def isPrime(nbr:int, primeList:list):
    """
    retourne si int donnée est premier
    """
    isPrime = False
    maxTest = math.sqrt(nbr)
    for i in primeList:
        if (i <= maxTest) and (nbr%i == 0): 
            return isPrime
    isPrime = True
    return isPrime

def getPrimeListUntil(max):
    primeList = [2]
    for i in range(3, max, 1) :
        if isPrime(i, primeList) :
            primeList.append(i)
    return primeList

def main():
    print(getPrimeListUntil(20))
    print(getPrimeListUntil(100))

if __name__ == "__main__":
    main()