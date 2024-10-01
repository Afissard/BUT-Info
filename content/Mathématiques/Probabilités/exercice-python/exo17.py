import math, numpy as np, matplotlib.pyplot, random


def tirage_simultané(urne, nb=4):
    main = []
    for _ in range(nb):
        boule = random.randint(0, len(urne)-1)
        main.append(urne[boule])
        urne.remove(urne[boule])
    return main

def simulation():
    urne = ["r", "r", "r", "b", "b", "b", "b", "b"]
    print(tirage_simultané(urne=urne))

def P(x):
    for i in range(len(x[0])-1):
        x[1][i] = (math.comb(3, int(x[0][i])) * math.comb(5, 4-int(x[0][i]))) / math.comb(8, 4)
    return x
    
if __name__ == "__main__":
    X = np.zeros((2,4))
    X[0]=np.arange(4)
    
    print(X)
    print(P(X))