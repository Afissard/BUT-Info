import math as mt
import time

##Exercice 1##
list_nb_a_tester=[13,1009,10007,100003,1000003,10000019,100000007,1000000007,10000000019,100000000003,1000000000039,100000000000031]

def isPrime(n, pl=[2])->bool:
    if n <=1 : return False
    for i in pl :
        if n % i == 0 : return False
    for i in range(pl[len(pl)-1], int(mt.sqrt(n)+1)): 
        if n % i == 0 : return False
    return True

def listPrimeUntil(n)->list:
    res_list = [2]
    for i in range(res_list[0]+1, n):
        if isPrime(i, res_list) :
            res_list.append(i)
    return res_list

def correction_list_prime(n):
    nombre_teste=3
    premiers=[2]
    while nombre_teste<=n:
        prime = True
        limite = mt.floor(mt.sqrt(nombre_teste))
        for i in premiers:
            if i>limite:break
            r = nombre_teste%i
            if r==0 :
                prime = False
                break
        if prime : premiers.append(nombre_teste)
        nombre_teste+=1
    return premiers


def main():
    print("1.1")
    for i in list_nb_a_tester :
        print(i,":",isPrime(i))
    
    print("1.2")
    for i in [(x**2) for x in range(12)]:
        print(i, listPrimeUntil(i))
    
    print("1.3")
    primeList = []
    for i in list_nb_a_tester :
        start = time.time()
        res = listPrimeUntil(i)
        end = time.time()
        print(i,":", end - start, "last prime is :", res[len(res)-1])
        primeList = res
    print(primeList)
if __name__ == "__main__":
    main()