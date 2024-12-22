from logic_gate import *

def et_table():
    et_ver_tab = [[]]
    for x in range(1):
        for y in range(1):
            et_ver_tab[x][y].append(et(x, y)) #BUG something out of range
    print(et_ver_tab)
    

def main():
    #print(et(False, True), et(False, False), et(True, True))
    et_table()

if __name__=="__main__":
    main()