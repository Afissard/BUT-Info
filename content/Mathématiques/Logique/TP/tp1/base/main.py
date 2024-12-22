import math

# Les base de python

def exo2_1(x:float):
    if x < 0:
        return math.sqrt(x*(-1))
    else :
        return math.sqrt(x)

def main():
    print("racine de 4 =", exo2_1(4), "\nracine de -20 =", exo2_1(-20))

if __name__ == "__main__":
    main()