import numpy as np

def q1():
    print(np.zeros(7)) # créer une liste de 0, de taille 7
    print(np.ones(6)) # créer une liste de 1, de taille 6
    print(np.identity(3)) # matrice identitaire de 3*3
    print(np.array([3,7,-1])) # explicite
    print(np.arange(10, 30, 5)) # créé liste de 10 à 30, par pas de 5
    print(np.linspace(0, 2, 9)) # créé liste de 0 à 2, de 9 nombre à écart égaux
    print(np.sin(np.linspace(0, 2*np.pi, 20)))

def q2():
    # soit deux matrice a et b
    a = np.array([[1, 3], [0, 4]])
    b = np.array([[0, 4], [-1, 1]])
    print("a :\n", a, "\nb :\n", b, "\n")

    print("\naddition :")
    print("+ :\n", a + b)
    print("add :\n", np.add(a, b))

    print("\n? :")
    print("dot :\n", a.dot(b)) 
    print("@ :\n", a @ b)

    print("\ntransposition :")
    print(a.transpose())

    print("\nalgèbre linéaire : matrice puissance (2) :")
    print(np.linalg.matrix_power(a, 2))

    print("\nalgèbre linéaire : déterminant :")
    print(np.linalg.det(a))

    print("\nalgèbre linéaire : inversion :")
    print(np.linalg.inv(a))

    print("\nshape : dimension de la matrice")
    print(a.shape)

    print("\nsum :")
    print("sum", a.sum())
    print("sum axe 0", a.sum(axis=0))
    print("sum axe 1", a.sum(axis=1))

    print("\nmin / max :")
    print("min :", a.min(), "\nmax :", a.max)

    print("index de matrice ?")
    print("a[1] :", a[1], "\na[0, 1] :", a[0, 1], "\na[0][1] :", a[0][1])

def main():
    # q1()
    q2()

if __name__ == "__main__":
    main()
