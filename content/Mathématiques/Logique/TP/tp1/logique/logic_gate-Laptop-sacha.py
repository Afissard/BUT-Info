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
    else : 
        return False

def non(x:bool)-> bool:
    if x :
        return False
    else : 
        return True