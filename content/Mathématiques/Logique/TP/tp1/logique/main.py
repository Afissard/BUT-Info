from logic_gate import *

def F1(p:bool, q:bool, r:bool)-> bool:
    return implique(non(p),(et(q,r)))

def F2(p:bool, q:bool, r:bool)-> bool:
    return implique(q, (et(non(p), non(r))))

def satisfiabilité():
    """
    Version casual-bourrin
    """
    satisfiable = False
    for p in range(2):
        for q in range(2):
            for r in range(2):
                if et(F1(p,q,r), F2(p,q,r)) :
                    print("F1 et F2 satisfiables pour : p:",p, "q",q, "r",r)
                    satisfiable = True
    return satisfiable

input_table_3v = [
    [False,     False,  False],
    [False,     False,  True],
    [False,     True,   False],
    [False,     True,   True],
    [True,      False,  False],
    [True,      False,  True],
    [True,      True,   False],
    [True,      True,   True]
]

def satisfiabilité_v2():
    for i in range(len(input_table_3v)):
        if et(F1(input_table_3v[i]), F2(input_table_3v[i])) :
                print("F1 et F2 satisfiables pour :", input_table_3v[i])

def main():
    #print(et(False, True), et(False, False), et(True, True))
    #print("table de vérité de ET :", et_table())
    print(satisfiabilité())

if __name__=="__main__":
    main()