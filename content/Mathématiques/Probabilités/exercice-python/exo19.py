import math, numpy as np, matplotlib.pyplot as plt, random

def showLaw(x, xName=""):
    temp = xName + "\n"
    for i in range(len(x)):
        for j in range(len(x[i])):
            temp += str(x[i][j]) + "\t"
        temp += "\n"
    print(temp)

"""
F: points obtenus grace à 2 lancers francs : 1 points chacun, p=3/5
E: points obtenus grace à 2 lancers excentré : 2 points chacun, p=1/2
T: points obtenus grace à 2 lancers : 3 points chacun, p=1/4

Attendu:
F=(
    (0,      1,      2),
    (4/25,   12/25,  9/25)
)

E=(
    (0,     1,      2),
    (1/4,   1/2,    1/4)
)

F+E=(
    (0,     2,      4       1,      3,      5,      6),
    (1/25,  17/100, 11/50,  3/25,   6/25,   3/25,   9/100)
)

H = convol(convol(F,E),T)
"""

FProba = 3/5
FValue = 1
EProba = 1/2
EValue = 2
TProba = 1/4
TValue = 3

def ballThrow():
    height = 2
    width = 3
    x = np.zeros((height,width))
    x[0]=np.arange(width)
    return x

def law(x, xp, xv):
    for i in range(len(x[0])):
        x[0][i] *= xv 
        x[1][i] = xp**i
    x[1][0] -= sum(x[1][1:])
    return x

def sumLaw(l1, l2):
    res = np.zeros((2,1))
    for i in range(l1.shape[1]):
        for j in range(l2.shape[1]):
            if (l1[0][i]+l2[0][j] not in res[0]):
                res = np.append(res, [ [l1[0][i] + l2[0][j]], [l1[1][i] + l2[1][j]] ], axis=1)
            else : 
                k = np.where(res[0] == l1[0][i] + l2[0][j])
                res[1][k] = res[1][k] + l1[1][i] * l2[1][j]
    return res

def mergeResByMean(a, b):
    # potentiellement hautement inutile
    res = np.zeros((2,1))
    for i in range(a):
        for j in range(a[i]):
            res[i][j] = np.mean([a[i][j], b[i][j]])
    return res

def simLaw(l, n):
    # nop
    res = []
    for _ in range(n):
        ball = random.uniform(0, 1)
        point = np.sum(l[0]) # supérieur au nombre de points possible
        for i in range(len(l)):
            if ball <= l[1][i] : # point >= l[0][i] and
                point = l[0][i]
                ball = l[1][i]
                print(point, ball)
            res.append(point)
    return res
    
def simCorriger(h, n=1000):
    # cassé
    simH = np.zeros(n)
    alea = np.random.rand(n)
    for i in range(n):
        for j in range(len(h.cum(0))):
            if alea[i]<h.cum[1][j]:
                simH[i] = h.cum[0][j]
                break
    return simH

if __name__ == "__main__":
    f = law(ballThrow(), FProba, FValue)
    e = law(ballThrow(), EProba, EValue)
    t = law(ballThrow(), TProba, TValue)
    showLaw(f, "F")
    showLaw(e, "E")
    showLaw(t, "T")
    fe = sumLaw(f, e)
    showLaw(fe, "FE")
    h = sumLaw(fe, t)
    showLaw(h, "H")
    
    # print(simLaw(h, 10))
    # print(simCorriger(h, 1000)) # cassé