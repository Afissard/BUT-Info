import numpy as np

def diagonale (matriceA):
    if len(matriceA) != len(matriceA[0]):
        matriceB = 0
    else :
        matriceB = np.zeros(matriceA.shape, int)
        for i in range(len(matriceA)) :
            matriceB[i][i] = matriceA[i][i]
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
    resMat = np.zeros((matA.shape[0], matB.shape[1]), int) # création de la matrice de résultat
    matB = transpose(matB) # ou resMat[y][x] = matA[y][x]*matB[x][Y]

    for y in range(resMat.shape[0]) : 
        for x in range(resMat.shape[1]) :
            # print(matA[y]*matB[x])
            resMat[y][x] = np.sum(matA[y]*matB[x])

    return resMat

def venderMonde(vec):
    matrice = np.zeros([vec.shape[0], vec.shape[0]], int)
    for y in range(matrice.shape[0]) : 
        for x in range(matrice.shape[1]) :
            matrice[y][x] = (x+1)*(y+1) # v(i)
    return matrice

def main():
    matriceA = np.array([[2,3], [4,5]])
    matriceB = np.array([[6,7], [8,9]])
    vector = np.array([2,3])
    print("diagonale de :\n", matriceA, "\n=>\n", diagonale(matriceA))
    print("transposition de :\n", matriceA, "\n=>\n", transpose(matriceA))
    print("produit matriciel de :\n", matriceA, "\net\n", matriceB, "\n=>\n", produit(matriceA, matriceB))
    print("VenderMonde de :\n", vector, "\n=>\n", venderMonde(vector))

if __name__ == "__main__":
    main()

