import numpy as np
import matplotlib.pyplot as plt


def baseLagrange(x, listX, i):
    """
    X : vecteur
    listeX : liste de nombre
    i : entier positif < len(listX)
    
    retourne image de X par la fonction (voir fiche)
    """
    res = 1
    for j in range(len(listX)):
        if j != i:
            res *= ((x - listX[j]) / (listX[i] - listX[j]))
    return res

def testBaseLarange():
    x = np.linspace(start=-1, stop=3, num=50)
    listX = [0, 1, 2]
    plt.figure()
    for i in range(len(listX)):
        plt.subplot(1, len(listX),i+1)
        plt.plot(x, baseLagrange(x, listX, i))
    plt.show()

def interlagrange1(x, listX, listY):
    assert len(listX) == len(listY), "taille listX et listY différente !"
    res = 0
    for j in range(len(listY)):
        res += baseLagrange(x, listX, j) * listY[j]
    return res

def testInterLagrange():
    x = np.linspace(start=-1, stop=3, num=50)
    listX = [0, 1, 2]
    listY = [-1, 3, 2]
    res = interlagrange1(x, listX, listY)
    plt.figure()
    plt.plot(res)
    plt.show()


def main():
    testBaseLarange()
    testInterLagrange()

if __name__=="__main__":
    main()