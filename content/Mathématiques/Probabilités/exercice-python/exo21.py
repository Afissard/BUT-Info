import numpy as np

##############################################################################

def d6(n): return np.random.randint(1, 6+1, size=n)

def d6_nb_res6(n=1000):
    nb6 = []
    for _ in range(n):
        res = d6(4)
        nb6.append(np.count_nonzero(res==6))
    return nb6

def mod1():
    print("On jette un d6 équilibré 4 fois et on note le nombre de 6 obtenus")
    print("Univers image de X(omega):", [0,1,2,3,4])
    print("Loi: X~B(4, 1/6)")
    print("E(X)= n.p = 4/6 =", 4/6)
    simu = d6_nb_res6(1000)
    # print(simu)
    print("moyenne simulation sur 1000 tirages:",np.mean(simu))       

##############################################################################

def urne(n): return np.random.randint(1, 10+1, size=n)

def actionA(): return(urne(1))

def actionB():
    res = urne(4)
    return np.count_nonzero(res==5)

def actionC():
    res = urne(4)
    return np.count_nonzero(2 in res or 3 in res or 4 in res) #todo fix with sum ???

def simuMod2(n=1000):
    resA = []
    resB = []
    resC = []
    for _ in range(n):
        resA.append(actionA)
        resB.append(actionB)
        resC.append(actionC)
    return np.mean(resA), np.mean(resB), np.mean(resC)

def mod2():
    print("Dans une urne contenant 10 boule indiscernables au toucher et numérotées de 1 à 10...")
    print(simuMod2())

if __name__ == "__main__":
    mod1()
    mod2()