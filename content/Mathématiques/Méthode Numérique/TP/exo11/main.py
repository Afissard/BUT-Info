def f1(x):
    return x**2-2

def df1(x):
    return 2*x

def dichotomie(a, b, f, e):
    """
    Args:
        a (float): borne 1 de l’intervale de recherche, avec f(a) et f(b) de signe opposés.
        b (float): borne 1 de l’intervale de recherche, avec f(a) et f(b) de signe opposés.
        f (function): la fonction dont on cherche une racine
        e (int): la tolérance qui détermine la précision de la solution
    """
    
    if f(a)*f(b) >= 0 : # a et b doivent être de signe opposés
        return SyntaxError
    
    m = (a + b)/2
    # print(m)
    
    if abs(f(m)) < e : 
        return m
    else :
        if f(a) * f(m) > 0 : 
            # a = m
            return dichotomie(m, b, f, e)
        else : 
            # b = m
            return dichotomie(a, m, f, e)

def fausse_pos(a, b, f, e):
    m = (a*f(b) - b*f(a)) / (f(b) - f(a))
    if abs(f(m)) < e : 
        return m
    else :
        if f(a)*f(m) > 0: 
            return fausse_pos(m, b, f, e)
        else : 
            return fausse_pos(a, m, f, e)

def newtown(a, f, df, e):
    m = a - f(a)/df(a)
    if abs(f(m)) < e : return m
    else : return newtown(m, f, df, e)

def main():
    e = 10**(-8)
    
    print("Dichotomie")
    # print("f1 : (-1;4):\t", dichotomie(-1, 4, f1, e))
    print("f1 : (1;4):\t", dichotomie(1, 4, f1, e))
    
    print("fausse position")
    print("f1 : (1;4):\t", fausse_pos(1, 4, f1, e))
    
    print("méthode de Newtown")
    print("f1 : (1):\t", newtown(1, f1, df1, e))
    
    
if __name__ == "__main__":
    main()
