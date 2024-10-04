import math, numpy as np, matplotlib.pyplot as plt, random

########################################################

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

########################################################

def getX(width=4, height=2):
    x = np.zeros((height,width))
    x[0]=np.arange(width)
    return x
    
def P(x):
    for i in range(len(x[0])):
        x[1][i] = (math.comb(3, int(x[0][i])) * math.comb(5, 4-int(x[0][i]))) / math.comb(8, 4)
    return x

def diagX(x):
    xCumul = x.copy()
    for i in range(len(x[0])): xCumul[1][i] = x[1][:i+1].sum()
    
    xFr = xCumul.copy()
    xFr = np.insert(xFr, 0, [3,0], axis=1)
    
    print(xCumul,"\n", xFr)
    
    for i in range (len(xFr[0])): plt.plot(([xFr[0][i], xFr[0][i+1]], [xFr[1][i], xFr[1][i]]))
    
    return True
    

if __name__ == "__main__":
    x = getX()
    x = P(x)
    diagX(x)