import numpy as np

def diagonale (matriceA):
    if len(matriceA) != len(matriceA[0]):
        matriceB = 0
    else :
        matriceB = np.zeros(matriceA.shape, int)
    return matriceB

def transpose(matrice):
    tMatrice = np.zeros([matrice.shape[1], matrice.shape[0]], int) # créer la matrice transposé
    for i in range(len(matrice)) :
        for j in range(len(matrice[i])) :
            tMatrice[j][i] = matrice[i][j] # rempli la matrice
    return tMatrice  

def produit(matA, matB):
    if matA.shape[1] != matB.shape[0] :
        return 0
    # else implicite
    resMat = np.zeros([len(matB), len(matB)], int) # création de la matrice de résultat
    for i in resMat : for j in i : for y in range(len(matA)) : for x in range(len(matA[y])) : j = matA[y][x]*matB[x][y]
    return resMat

def main():
    matriceA = np.array([[2,3], [4,5]])
    matriceB = np.array([[6,7], [8,9]])
    print("diagonale de :\n", matriceA, "\n=>\n", diagonale(matriceA))
    print("transposition de :\n", matriceA, "\n=>\n", transpose(matriceA))
    print("produit matriciel de :\n", matriceA, "\net\n", matriceB, "\n=>\n", produit(matriceA, matriceB))

if __name__ == "__main__":
    main()