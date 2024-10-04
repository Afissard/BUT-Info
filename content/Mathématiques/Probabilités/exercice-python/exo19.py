import math, numpy as np, matplotlib.pyplot as plt

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
        x[1][i] = (len(x)-i)*xp
    return x



if __name__ == "__main__":
    f = law(ballThrow(), FProba, FValue)
    e = law(ballThrow(), EProba, EValue)
    t = law(ballThrow(), TProba, TValue)
    showLaw(f, "F")
    showLaw(e, "E")
    showLaw(t, "T")