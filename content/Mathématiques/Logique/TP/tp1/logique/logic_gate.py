"""
All with NAND : https://en.wikipedia.org/wiki/NAND_logic
"""

def et(x:bool, y:bool)-> bool:
    if x :              # if x == True
        return y
    else :
        return x

def ou(x:bool, y:bool)-> bool:
    if x :
        return x
    elif y :
        return y
    #else : 
    return False

def non(x:bool)-> bool:
    if x :
        return False
    #else : 
    return True

def implique(x:bool, y:bool)-> bool:
    return et(non(x), y)

###############################################

def et_table():
    ver_table = []
    for x in range(2): # 0 et 1 inclus
        for y in range(2):
            ver_table.append(et(x, y))
            #print(et(x, y))
    return ver_table

def ou_table():
    ver_table = []
    for x in range(2):
        for y in range(2):
            ver_table.append(ou(x,y))
    return ver_table