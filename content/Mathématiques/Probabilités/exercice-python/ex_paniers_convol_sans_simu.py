import numpy as np


f=np.array([[0,1,2],[4/25,12/25,9/25]])
e=np.array([[0,2,4],[1/4,1/2,1/4]])
t=np.array([[0,3,6],[9/16,6/16,1/16]])

def convol(x1,x2):
    res1=np.zeros((2,1))
    for i in range(x1.shape[1]):
        for j in range(x2.shape[1]):
            if (x1[0][i]+x2[0][j] not in res1[0]):
                res1=np.append(res1,[[x1[0][i]+x2[0][j]],[x1[1][i]*x2[1][j]]],axis=1)
            else :
                k=np.where(res1[0]==x1[0][i]+x2[0][j])
                res1[1][k]=res1[1][k]+x1[1][i]*x2[1][j]
    ind=np.argsort(res1[0])#pas indispensable
    res1[0]=res1[0][ind]#pas indispensable
    res1[1]=res1[1][ind]#pas indispensable
    return res1

float_formatter = "{:.3f}".format#pas indispensable
np.set_printoptions(formatter={'float_kind':float_formatter})#pas indispensable
res_h=convol(convol(f,e),t)
print('loi H : ',res_h)
#print('verif somme probas : ',res_h[1].sum())
espe=(res_h[0]*res_h[1]).sum()
print('espérance=',espe)
variance=((res_h[0]-espe)**2*res_h[1]).sum()
print('variance=',variance)
print('écart type=',variance**(1/2))

