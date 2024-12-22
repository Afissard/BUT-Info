import numpy as np
import matplotlib.pyplot as plt
import math as mt


def BaseLagrange(x,listX,i):
    return #une valeur

def InterLagrange(x,listX,listY):
    return #une valeur

def dicho(a,b,f,e):
    return #liste_valeurs

def fausse_pos(a,b,f,e):
    return #liste_valeurs

def newton(a,f,df,e):
    return #liste_valeurs


###QUELQUES TESTS INDICATIFS

# print('test BaseLagrange : ',np.isclose(BaseLagrange(0.2,np.linspace(-1,1,10),6),0.37638881280000036))
# print('test2 BaseLagrange : ',np.isclose(BaseLagrange(0.2,np.linspace(-1,1,11),6),1))
# print('test3 BaseLagrange : ',np.isclose(BaseLagrange(0.2,np.linspace(-1,1,11),4),0))

# print('test1 InterLagrange : ',np.isclose(InterLagrange(np.pi/3,np.linspace(0,np.pi,7),np.cos(np.linspace(0,np.pi,7))),0.5))
# print('test2 InterLagrange : ',np.isclose(InterLagrange(np.pi/3,np.linspace(0,np.pi,5),np.cos(np.linspace(0,np.pi,5))),0.4969732592091243))

# list_res_dicho1=[1.5, 1.25, 1.375, 1.4375, 1.40625, 1.421875, 1.4140625, 1.41796875, 1.416015625, 1.4150390625, 1.41455078125, 1.414306640625, 1.4141845703125, 1.41424560546875, 1.414215087890625]
# print('test dicho : ',list(np.isclose(np.array(dicho(1,2,lambda x:x**2-2,10**(-5))),np.array(list_res_dicho1))).count(True)==len(list_res_dicho1))

# list_res_fausse_pos1=[1.3333333333333333, 1.4, 1.411764705882353, 1.4137931034482758, 1.414141414141414, 1.4142011834319526, 1.41421143847487]
# print('test fausse_pos : ',list(np.isclose(np.array(fausse_pos(1,2,lambda x:x**2-2,10**(-5))),np.array(list_res_fausse_pos1))).count(True)==len(list_res_fausse_pos1))

# list_res_newton=[1.5, 1.4166666666666667, 1.4142156862745099]
# print('test newton : ',list(np.isclose(np.array(newton(2,lambda x:x**2-2,lambda x:2*x,10**(-5))),np.array(list_res_newton))).count(True)==len(list_res_newton))
