import math

def bezout(a,b): # algo d'euclide étendu
    # if not(a>b):
    # if (a<=b):
    #     return ValueError
    
    u0, u1, v0, v1 = 1, 0, 0, 1
    
    # r , q = a%b, a//b
    q, r = divmod(a, b) 

    while r != 0 :
        a, b = b, r
        u2 = u0 - q*u1
        v2 = v0 - q*v1
        u0, v0, u1, v1 = u1, v1, u2, v2
        q, r = divmod(a, b) 
    
    return (b, u1, v1)


if __name__ == "__main__":
    print(bezout(3315, 154))

    # for i in range(1, 10):
    #     for j in range(1, 10):
    #         print(bezout(i,j))
