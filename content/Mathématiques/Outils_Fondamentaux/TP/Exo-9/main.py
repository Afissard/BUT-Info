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

def main():
    matriceA = np.array([[2,3], [4,5]])
    matriceB = np.array([[6,7], [8,9]])
    print("diagonale de :\n", matriceA, "\n=>\n", diagonale(matriceA))
    print("transposition de :\n", matriceA, "\n=>\n", transpose(matriceA))

if __name__ == "__main__":
    main()