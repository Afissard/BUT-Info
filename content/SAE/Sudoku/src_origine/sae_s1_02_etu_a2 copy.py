import numpy as np
import copy
import time

#TROIS PETITES FONCTIONS DE TEST UTILISEES PLUS BAS#
def test(mess,eval,res):
    print(mess,(eval==res)*'OK'+(eval!=res)*'Try again')
def test_determine_valuations(mess,list_var,res):
    test=mess+'Ok'
    list_testee=determine_valuations(list_var)
    for el in list_testee :
        if el not in res:
            test=mess+'Try again'
            return test
    for el in res:
        if el not in list_testee :
            test=mess+'Try again'
            return test
    for i in range(len(list_testee)-1):
        if list_testee[i] in list_testee[i+1:]:
            test=mess+'wowowow y a du doublon là-dedans'
            return test
    return test  

def test_for(mess,formu,res_for):
    res=True
    if (formu==[] and res_for!=[]) or (formu!=[] and res_for==[]):
        res=False
    for el1 in formu:
        for el2 in res_for:
            res=(set(el1)==set(el2))
            if res :
                break
        if not res :
            print(mess+'Try again !')
            return
    for el2 in res_for:
        for el1 in formu:
            res=(set(el2)==set(el1))
            if res :
                break
        if not res :
            print(mess+'Try again !')
            return
    res=False
    for i in range(len(formu)-1):
        for el in formu[i+1:]:
            if set(formu[i])==set(el):
                print(mess+'wowowow y a du doublon là-dedans')
                return 
    print(mess+'Ok')
    
#A VOUS DE JOUER#
def evaluer_clause(clause,list_var):
    '''Arguments : une liste d'entiers non nuls traduisant une clause,une liste de booléens informant de valeurs logiques connues (ou None dans le cas contraire) pour un ensemble de variables
    Renvoie : None ou booléen
'''
    
'''clause1=[1,-2,3,-4]
list_var1=[True,True,False,None]
test("essai cas 1 evaluer_clause : ",evaluer_clause(clause1,list_var1),True)
clause2=[1,-2,3,-4]
list_var2=[False,True,False,None]
test("essai cas 2 evaluer_clause : ",evaluer_clause(clause2,list_var2),None)
clause3=[1,-2,3,-4]
list_var3=[None,True,False,True]
test("essai cas 3 evaluer_clause : ",evaluer_clause(clause3,list_var3),None)
clause4=[1,-3]
list_var4=[False,False,True]
test("essai cas 4 evaluer_clause : ",evaluer_clause(clause4,list_var4),False)
clause5=[]
list_var5=[False,False,True]
test("essai cas 5 evaluer_clause : ",evaluer_clause(clause5,list_var5),False)
clause6=[1,2,3]
list_var6=[False,False,True]
test("essai cas 6 evaluer_clause : ",evaluer_clause(clause6,list_var6),True)
'''

def evaluer_cnf(formule,list_var):
    '''Arguments : une liste de listes d'entiers non nuls traduisant une formule,une liste de booléens informant de valeurs logiques connues (ou None dans le cas contraire) pour un ensemble de variables
    Renvoie : None ou booléen
'''
    
'''for1=[[1,2],[2,-3,4],[-1,-2],[-1,-2,-3],[1]]
list_var_for1_test1=[True,False,False,None]
test('test1 evaluer_cnf : ',evaluer_cnf(for1,list_var_for1_test1),True)
list_var_for1_test2=[None,False,False,None]
test('test2 evaluer_cnf : ',evaluer_cnf(for1,list_var_for1_test2),None)
list_var_for1_test3=[True,False,True,False]
test('test3 evaluer_cnf : ',evaluer_cnf(for1,list_var_for1_test3),False)'''

def determine_valuations(list_var):
    '''Arguments : une liste de booléens informant de valeurs logiques connues (ou None dans le cas contraire) pour un ensemble de variables
    Renvoie : La liste de toutes les valuations (sans doublon) envisageables pour les variables de list_var
'''
    
'''
list_var1=[True,None,False,None]
print(test_determine_valuations('res_test_determine_valuations cas 1 : ',list_var1,[[True, True, False, True], [True, False, False, True], [True, True, False, False], [True, False, False, False]]))
list_var2=[None,False,True,None,True,False]
print(test_determine_valuations('res_test_determine_valuations cas 2 : ',list_var2,[[True, False, True, True, True, False], [False, False, True, True, True, False], [True, False, True, False, True, False], [False, False, True, False, True, False]]))
list_var3=[False,True,True,False]
print(test_determine_valuations('res_test_determine_valuations cas 3 : ',list_var3,[[False, True, True, False]]))
list_var4=[None,None,None]
print(test_determine_valuations('res_test_determine_valuations cas 4 : ',list_var4,[[True, True, True], [False, True, True], [True, False, True], [False, False, True], [True, True, False], [False, True, False], [True, False, False], [False, False, False]]))
'''

def resol_sat_force_brute(formule,list_var):
    '''Arguments : une liste de listes d'entiers non nuls traduisant une formule,une liste de booléens informant de valeurs logiques connues (ou None dans le cas contraire) pour un ensemble de variables
    Renvoie : SAT,l1
    avec SAT : booléen indiquant la satisfiabilité de la formule
          l1 : une liste de valuations rendant la formule vraie ou une liste vide
'''
'''
for1=[[1,2],[2,-3,4],[-1,-2],[-1,-2,-3],[1],[-1,2,3]]
list_var_for1=[None,None,None,None]
test('test1 resol_sat_force_brute : ',resol_sat_force_brute(for1,list_var_for1),(True,[True, False, True, True]))

for2=[[1,4,-5],[-1,-5],[2,-3,5],[2,-4],[2,4,5],[-1,-2],[-1,2,-3],[-2,4,-5],[1,-2]]
list_var_for2=[None,None,None,None,None]
test('test2 resol_sat_force_brute : ',resol_sat_force_brute(for2,list_var_for2),(False,[]))

for3=[[-1,-2],[-1,2,-3,4],[2,3,4],[3],[1,-4],[-1,2],[1,2]]
list_var_for3=[None,None,None,None]
test('test3 resol_sat_force_brute : ',resol_sat_force_brute(for3,list_var_for3),(True,[False, True, True, False]))

for4=[[-1,-2],[-1,2,-3,4],[2,3,4],[3],[1,-4],[-1,2],[1,2]]
list_var_for4=[None,None,None,True]
test('test4 resol_sat_force_brute : ',resol_sat_force_brute(for4,list_var_for4),(False,[]))


'''



def enlever_litt_for(formule,litteral):
    '''Arguments :
formule : comme précédemment
litteral : un entier non nul traduisant la valeur logique prise par une variable
    Renvoie : la formule simplifiée
'''
    
'''for1=[[1,2,4,-5],[-1,2,3,-4],[-1,-2,-5],[-3,4,5],[-2,3,4,5],[-4]]
litt1=4
test('essai cas 1 enlever_litt_for : ',enlever_litt_for(for1,litt1),[[-1, 2, 3], [-1, -2, -5], []])'''

def init_formule_simpl_for(formule_init,list_var):
    '''
    Renvoie : La formule simplifiée en tenant compte des valeurs logiques renseignées dans list_var
'''

'''
list_var_for1=[False, None, None, False, None]
for1=[[-5, -3, 4, -1], [3], [5, -2], [-2, 1, -4], [1, -3]]
cor_for1=[[3], [5, -2], [-3]]
test_for('test1_init_formule_simpl_for : ',init_formule_simpl_for(for1,list_var_for1),cor_for1)

list_var_for2= [False, True, False, True, False]
for2= [[3, 2, 1], [-1, -2, 5]]
cor_for2=[]
test_for('test2_init_formule_simpl_for : ',init_formule_simpl_for(for2,list_var_for2),cor_for2)

list_var_for3= [None, None, None, True, None]
for3= [[-5, -1], [-1, -3], [4], [-4, 1], [-2, -1, 3]]
cor_for3=[[-5, -1], [-1, -3], [1], [-2, -1, 3]]
test_for('test3_init_formule_simpl_for : ',init_formule_simpl_for(for3,list_var_for3),cor_for3)
'''

def retablir_for(formule_init,list_chgmts):
    '''Arguments : une formule initiale et une liste de changements à apporter sur un ensemble de variables (chaque changement étant une liste [i,bool] avec i l'index qu'occupe la variable dans list_var et bool la valeur logique qui doit lui être assignée) 
    Renvoie : la formule simplifiée en tenant compte de l'ensemble des changements
'''
    

'''
formule_init=[[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
list_chgmts1=[[0, True], [1, True], [2, False]]
form1=[[-5], [4, 5], [-4, 5]]

list_chgmts2=[[0, True], [1, True], [2, False], [3, True], [4, False]]
form2=[[]]

list_chgmts3=[[0, True], [1, True], [2, False], [3, False]]
form3=[[-5], [5]]
test('essai cas 1 retablir_for : ',retablir_for(formule_init,list_chgmts1),form1)
test('essai cas 2 retablir_for : ',retablir_for(formule_init,list_chgmts2),form2)
test('essai cas 3 retablir_for : ',retablir_for(formule_init,list_chgmts3),form3)
'''

def progress(list_var,list_chgmts):
    '''Arguments : list_var, list_chgmts définies comme précédemment
    Renvoie : l1,l2
    l1 : nouvelle list_var 
    l2 : nouvelle list_chgmts 
'''
'''
list_var=[True, None, None, None, None]
list_chgmts=[[0, True]]
l1=[True, True, None, None, None]
l2=[[0, True], [1, True]]
test("essai cas 1 progress : ",progress(list_var,list_chgmts),(l1,l2))

list_var=[True, False, False, None, None]
list_chgmts=[[0, True], [1, False], [2, False]]
l1=[True, False, False, True, None]
l2=[[0, True], [1, False], [2, False], [3, True]]
test("essai cas 2 progress : ",progress(list_var,list_chgmts),(l1,l2))

list_var=[None, None, None, None, None]
list_chgmts=[]
l1=[True, None, None, None, None]
l2=[[0, True]]
test("essai cas 3 progress : ",progress(list_var,list_chgmts),(l1,l2))

list_var=[False, None, None, None, None]
list_chgmts=[[0, False]]
l1=[False, True, None, None, None]
l2=[[0, False], [1, True]]
test("essai cas 4 progress : ",progress(list_var,list_chgmts),(l1,l2))

list_var=[True, False, None, None, None]
list_chgmts=[]
l1=[True, False, True, None, None]
l2=[[2, True]]
test("essai cas 5 progress : ",progress(list_var,list_chgmts),(l1,l2))

list_var=[True, False, False, None, None]
list_chgmts=[[2, False]]
l1=[True, False, False, True, None]
l2=[[2, False], [3, True]]
test("essai cas 6 progress : ",progress(list_var,list_chgmts),(l1,l2))

'''

def progress_simpl_for(formule,list_var,list_chgmts):
    '''Arguments : formule,list_var, list_chgmts définies comme précédemment
    Renvoie : form,l1,l2
    form : nouvelle formule
    l1 : nouvelle list_var 
    l2 : nouvelle list_chgmts 
'''
    
    
'''
formule= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
list_var= [None, None, None, None, None] 
list_chgmts= []
cor_form,cor_l1,cor_l2= ([[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]],[True, None, None, None, None],[[0, True]])
test('essai1_progress_simpl_for : ',progress_simpl_for(formule,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
 
 
formule= [[-5], [5]] 
list_var= [True, True, True, False, None] 
list_chgmts= [[0, True], [1, True], [2, True], [3, False]]
cor_form,cor_l1,cor_l2= ([[]],[True, True, True, False, True],[[0, True], [1, True], [2, True], [3, False], [4, True]])
test('essai2_progress_simpl_for : ',progress_simpl_for(formule,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))

formule= [[3, -4], [-3, 4, 5], [-4, 5]] 
list_var= [True, False, None, None, None] 
list_chgmts= [[0, True], [1, False]]
cor_form,cor_l1,cor_l2= ([[4, 5], [-4, 5]],[True, False, True, None, None],[[0, True], [1, False], [2, True]])
test('essai3_progress_simpl_for : ',progress_simpl_for(formule,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
'''    

def progress_simpl_for_dpll(formule,list_var,list_chgmts,list_sans_retour):
    '''Arguments : list_sans_retour contient l'ensemble des numéros de variables auxquelles on a affecté une valeur logique sur laquelle on ne reviendra pas
    renvoie :form,l1,l2,l3 avec :
    form : la formule simplifiée
    l1 : la liste actualisée des valeurs attribuées aux variables après le changement effectué
    l2 : la liste actualisée de l'ensemble des changements effectués
    l3 : la liste éventuellement actualisée des numéros de variables auxquelles une affectation a été attribuée sur laquelle on ne reviendra pas
    '''

'''
formule= [[-5], [4, 5], [-4, 5]] 
list_var= [True, True, False, None, None] 
list_chgmts= [[0, True], [1, True], [2, False]] 
list_sans_retour= []
cor_for,cor_l1,cor_l2,cor_l3= ([[4], [-4]],[True, True, False, None, False],[[0, True], [1, True], [2, False], [4, False]],[4])
test('essai1_progress_simpl_for_dpll : ',progress_simpl_for_dpll(formule,list_var,list_chgmts,list_sans_retour),(cor_for,cor_l1,cor_l2,cor_l3))

formule= [[-5,4], [2,4, 5], [-2, 5]]
list_var= [True, None, None, None, None] 
list_chgmts= [[0, True]] 
list_sans_retour= [0]
cor_for,cor_l1,cor_l2,cor_l3= ([[-2,5]],[True, None, None, True, None],[[0, True],[3, True]],[0,3])
test('essai2_progress_simpl_for_dpll : ',progress_simpl_for_dpll(formule,list_var,list_chgmts,list_sans_retour),(cor_for,cor_l1,cor_l2,cor_l3))

formule=[[1,2,4,-5],[-1,2,3,-4],[-1,-2,-5],[-3,4,5],[-2,3,4,5],[-4,5]]
list_var=[None,None,None,None,None]
list_chgmts= [] 
list_sans_retour= []
cor_for,cor_l1,cor_l2,cor_l3=([[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]], [True, None, None, None, None], [[0, True]], [])
test('essai3_progress_simpl_for_dpll : ',progress_simpl_for_dpll(formule,list_var,list_chgmts,list_sans_retour),(cor_for,cor_l1,cor_l2,cor_l3))
'''  

def retour(list_var,list_chgmts):
    '''
    renvoie :l1,l2 avec :
    l1 : la liste actualisée des valeurs attribuées aux variables 
    l2 : la liste actualisée de l'ensemble des changements effectués depuis une formule initiale
    
    '''
'''
list_var= [True, True, None, None, None]
list_chgmts= [[0, True], [1, True]]
l1= [True, False, None, None, None]
l2= [[0, True], [1, False]]
test("essai cas 1 retour : ",retour(list_var,list_chgmts),(l1,l2))

list_var= [True, False, None, None, None]
list_chgmts= [[0, True], [1, False]]
l1= [False, None, None, None, None]
l2= [[0, False]]
test("essai cas 2 retour : ",retour(list_var,list_chgmts),(l1,l2))

list_var= [True, False, False, True, None]
list_chgmts= []
l1= [True, False, False, True, None]
l2= []
test("essai cas 3 retour : ",retour(list_var,list_chgmts),(l1,l2))

list_var= [True, False, False, False, False]
list_chgmts= [[0, True], [1, False], [2, False], [3, False], [4, False]]
l1= [False, None, None, None, None]
l2= [[0, False]]
test("essai cas 4 retour : ",retour(list_var,list_chgmts),(l1,l2))

list_var= [True, True, False, True, None]
list_chgmts= [[1, True]]
l1= [True, False, False, True, None]
l2= [[1, False]]
test("essai cas 5 retour : ",retour(list_var,list_chgmts),(l1,l2))

list_var= [True, False, False, True, None]
list_chgmts= [[1, False]]
l1= [True, None, False, True, None]
l2= []
test("essai cas 6 retour : ",retour(list_var,list_chgmts),(l1,l2))
'''

def retour_simpl_for(formule_init,list_var,list_chgmts):
    '''
Renvoie : form,l1,l2
    form : nouvelle formule
    l1 : nouvelle list_var 
    l2 : nouvelle list_chgmts 
'''

'''
formule_init= [[-2, 1, -5, -4], [2, 4, -1], [-5, 4], [1, 4, -2], [-4, -2, 5]] 
list_var= [True, True, False, False, True] 
list_chgmts= [[0, True], [4, True]]
cor_form,cor_l1,cor_l2= ([[2, 4], [-4, -2]],[True, True, False, False, False],[[0, True], [4, False]])
test('essai1_retour_simpl_for : ',retour_simpl_for(formule_init,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))

formule_init= [[5, 4, -3], [-2, -5, 3], [-1]] 
list_var= [False, True, True, False, False] 
list_chgmts= [[2, True]]
cor_form,cor_l1,cor_l2= ([[-2, -5], [-1]],[False, True, False, False, False],[[2, False]])
test('essai2_retour_simpl_for : ',retour_simpl_for(formule_init,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
'''    

def retour_simpl_for_dpll(formule_init,list_var,list_chgmts,list_sans_retour):
    '''
Renvoie : form,l1,l2,l3
    form : nouvelle formule
    l1 : nouvelle list_var 
    l2 : nouvelle list_chgmts
    l3 : nouvelle list_sans_retour
'''
'''
formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
list_var= [True, True, False, True, False] 
list_chgmts= [[0, True], [1, True], [2, False], [4, False], [3, True]] 
list_sans_retour= [4, 3]
cor_form,cor_l1,cor_l2,cor_l3= ([[3, -4], [-3, 4, 5], [-4, 5]], [True, False, None, None, None], [[0, True], [1, False]], [])
test('essai1_retour_simpl_for_dpll : ',retour_simpl_for_dpll(formule_init,list_var,list_chgmts,list_sans_retour),(cor_form,cor_l1,cor_l2,cor_l3))

formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
list_var= [True, True, True, True, False] 
list_chgmts= [[0, True], [1, True], [2, True], [3, True], [4, False]] 
list_sans_retour= []
cor_form,cor_l1,cor_l2,cor_l3= ([[-5], [5]], [True, True, True, False, None], [[0, True], [1, True], [2, True], [3, False]], [])
test('essai2_retour_simpl_for_dpll : ',retour_simpl_for_dpll(formule_init,list_var,list_chgmts,list_sans_retour),(cor_form,cor_l1,cor_l2,cor_l3))

formule_init= [[3, 1], [1], [-2, 3, -5], [-1, 3], [-4, -3, -2]] 
list_var= [True, None, False, None, True] 
list_chgmts= [[0, True]] 
list_sans_retour= [0]
cor_form,cor_l1,cor_l2,cor_l3= ([[3, 1], [1], [-2, 3, -5], [-1, 3], [-4, -3, -2]], [None, None, False, None, True], [], [])
test('essai3_retour_simpl_for_dpll : ',retour_simpl_for_dpll(formule_init,list_var,list_chgmts,list_sans_retour),(cor_form,cor_l1,cor_l2,cor_l3))
'''


def resol_parcours_arbre(formule_init,list_var,list_chgmts):
    '''Renvoie : SAT,l1
    avec SAT : booléen indiquant la satisfiabilité de la formule
          l1 : une liste de valuations rendant la formule vraie ou une liste vide'''
    
'''
formule_init= [[1, 4, -5], [-1, -5], [2, -3, 5], [2, -4], [2, 4, 5], [-1, -2], [-1, 2, -3], [-2, 4, -5], [1, -2]] 
list_var= [True, True, False, True, None] 
list_chgmts= [[1, True]]
cor_resol=(False, [])
test('essai1_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)

formule_init= [[5], [3, -5, -1, -2], [1, -2, -5], [2, -5, 1, -3], [3]] 
list_var= [True, False, None, False, True] 
list_chgmts= [[0, True]]
cor_resol=(True,[True, False, True, False, True])
test('essai2_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)

formule_init= [[-5, 2, -3, -4], [1, -5], [5, 2], [3, -2, 4], [5, -2, -1]] 
list_var= [False, True, False, None, None] 
list_chgmts= [[1, True]]
cor_resol=(True,[False, True, False, True, False])
test('essai3_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)
'''   

def resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts):#la même distinction peut être faite entre formule et formule_init
    '''
    Renvoie SAT,l1 avec :
SAT=True ou False
l1=une liste de valuations rendant la formule vraie ou une liste vide
''' 
        #Initialisation du parcours
    if list_chgmts==[]:
        if [] in formule:
            return False,[]
        if formule==[]:
            return True,list_var
        form,list_var_init,list_chgmts_init=progress_simpl_for(formule,list_var,[])
        return resol_parcours_arbre_simpl_for(formule_init,form,list_var_init,list_chgmts_init)
    #Reste du parcours à implémenter :
'''
formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
formule= [[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
list_var= [True, None, None, None, None] 
list_chgmts= [[0, True]]
cor_resol=(True, [True, False, True, True, True])
test('essai1_resol_parcours_arbre_simpl_for : ',resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts),cor_resol)

formule_init= [[5], [3, -5, -1, -2], [1, -2, -5], [2, -5, 1, -3], [3]] 
formule= [[5], [-5]] 
list_var= [False, True, True, False, None] 
list_chgmts= [[2, True]]
cor_resol=(False, [])
test('essai2_resol_parcours_arbre_simpl_for : ',resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts),cor_resol)

formule_init= [[-5, 2, -3, -4], [1, -5], [5, 2], [3, -2, 4], [5, -2, -1]] 
formule= [[-5], [4]] 
list_var= [False, True, False, None, None] 
list_chgmts= [[1, True]]
cor_resol=(True, [False, True, False, True, False])
test('essai3_resol_parcours_arbre_simpl_for : ',resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts),cor_resol)
'''

def resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour):
    '''
    Renvoie SAT,l1 avec :
SAT=True ou False
l1=une liste de valuations rendant la formule vraie ou une liste vide
'''
    #Initialisation du parcours
    if list_chgmts==[]:
        if [] in formule:
            return False,[]
        if formule==[]:
            return True,list_var
        form,list_var_init,list_chgmts_init,list_sans_retour_init=progress_simpl_for_dpll(formule,list_var,[],[])
        return resol_parcours_arbre_simpl_for_dpll(formule_init,form,list_var_init,list_chgmts_init,list_sans_retour_init)

'''
formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
formule= [[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
list_var= [True, None, None, None, None] 
list_chgmts= [[0, True]] 
list_sans_retour= []
cor_resol=(True, [True, False, True, None, True])
test('essai1_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)

formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]] 
formule= [[3, -4]] 
list_var= [True, False, None, None, True] 
list_chgmts= [[0, True], [1, False], [4, True]] 
list_sans_retour= [4]
cor_resol=(True, [True, False, True, None, True])
test('essai2_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)

formule_init= [[-5, 2, -3, -4], [1, -5], [5, 2], [3, -2, 4], [5, -2, -1]] 
formule= [[2], [-2, 4]] 
list_var= [False, None, False, None, False] 
list_chgmts= [[4, False]] 
list_sans_retour= [4]
cor_resol=(True, [False, True, False, True, False])
test('essai3_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)

formule_init= [[5], [3, -5, -1, -2], [1, -2, -5], [2, -5, 1, -3], [3]] 
formule= [[-2],[2,-3],[3]] 
list_var= [False, None, None, False, True] 
list_chgmts= [[4, True]]
list_sans_retour=[4]
cor_resol=(False, [])
test('essai4_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)
'''
        
def ultim_resol(formule_init,list_var):
    '''
    Renvoie SAT,l1 avec :
SAT=True ou False
l1=une liste de valuations rendant la formule vraie ou une liste vide

    Affichage possible du temps mis pour la résolution
'''

def ultim_resol_simpl_for(formule_init,list_var):
    '''
    Renvoie SAT,l1 avec :
SAT=True ou False
l1=une liste de valuations rendant la formule vraie ou une liste vide

    Affichage possible du temps mis pour la résolution
'''

def ultim_resol_simpl_for_dpll(formule_init,list_var):
    '''
    Renvoie SAT,l1 avec :
SAT=True ou False
l1=une liste de valuations rendant la formule vraie ou une liste vide

    Affichage possible du temps mis pour la résolution
'''

def creer_grille_init(list_grille,n):
    '''Arguments : une liste de listes(contenant les coordonnées à renseigner et le nombre correspondant) et un entier donnant la taille de la grille
        Renvoie : une liste (list_grille_complete) avec les valeurs qui devront s'afficher dans la grille en la parcourant ligne après ligne de haut en bas et de gauche à droite
'''
'''
list_grille3=[[1,3,2],[1,6,5],[2,5,4],[2,8,9],[2,9,3],[3,2,7],[3,9,6],[4,3,1],[4,4,8],[4,8,3],[5,1,7],[5,2,2],[5,5,6],[5,8,8],[5,9,4],[6,2,4],[6,6,2],[6,7,5],[7,1,3],[7,8,1],[8,1,4],[8,2,6],[8,5,7],[9,4,9],[9,7,8]]
cor_grille3=[0, 0, 2, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 9, 3, 0, 7, 0, 0, 0, 0, 0, 0, 6, 0, 0, 1, 8, 0, 0, 0, 3, 0, 7, 2, 0, 0, 6, 0, 0, 8, 4, 0, 4, 0, 0, 0, 2, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 1, 0, 4, 6, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 8, 0, 0]
test("essai creer_grille_init : ",creer_grille_init(list_grille3,3),cor_grille3)
'''
def creer_grille_final(list_var,n):
    '''
    Renvoie : une liste (list_grille_complete) avec les valeurs qui devront s'afficher dans la grille (en fonction des valeurs logiques prises par les variables de list_var) en la parcourant ligne après ligne de haut en bas et de gauche à droite
'''
'''
list_var_fin=[False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, False]
cor_grille_final=[9, 3, 2, 6, 1, 5, 4, 7, 8, 5, 8, 6, 2, 4, 7, 1, 9, 3, 1, 7, 4, 3, 8, 9, 2, 5, 6, 6, 9, 1, 8, 5, 4, 7, 3, 2, 7, 2, 5, 1, 6, 3, 9, 8, 4, 8, 4, 3, 7, 9, 2, 5, 6, 1, 3, 5, 9, 4, 2, 8, 6, 1, 7, 4, 6, 8, 5, 7, 1, 3, 2, 9, 2, 1, 7, 9, 3, 6, 8, 4, 5]
test("essai creer_grille_final : ",creer_grille_final(list_var_fin,3),cor_grille_final)
'''
def afficher_grille(grille,n):
    '''
    ça affiche la grille
'''

def for_conj_sudoku(n):
    '''
    Renvoie : la formule (liste de listes) associée à une grille de sudoku de taille n selon les attentes formulées dans le sujet
    '''
'''
corrige_for2=[[-1, -21], [-1, -5], [-1, -17], [-1, -9], [-1, -33], [-1, -13], [-1, -49], [-1, -2], [-1, -3], [-1, -4], [-5, -17], [-5, -9], [-17, -33], [-5, -13], [-17, -49], [-5, -6], [-5, -7], [-5, -8], [-9, -13], [-33, -49], [-9, -10], [-9, -11], [-9, -12], [-13, -14], [-13, -15], [-13, -16], [1, 5, 9, 13], [1, 17, 33, 49], [1, 5, 17, 21], [-9, -29], [-17, -21], [-5, -21], [-17, -25], [-5, -37], [-17, -29], [-5, -53], [-17, -18], [-17, -19], [-17, -20], [-13, -25], [-21, -25], [-21, -37], [-21, -29], [-21, -53], [-21, -22], [-21, -23], [-21, -24], [-25, -29], [-37, -53], [-25, -26], [-25, -27], [-25, -28], [-29, -30], [-29, -31], [-29, -32], [17, 21, 25, 29], [5, 21, 37, 53], [9, 13, 25, 29], [-33, -53], [-33, -37], [-9, -25], [-33, -41], [-9, -41], [-33, -45], [-9, -57], [-33, -34], [-33, -35], [-33, -36], [-37, -49], [-37, -41], [-25, -41], [-37, -45], [-25, -57], [-37, -38], [-37, -39], [-37, -40], [-41, -45], [-41, -57], [-41, -42], [-41, -43], [-41, -44], [-45, -46], [-45, -47], [-45, -48], [33, 37, 41, 45], [9, 25, 41, 57], [33, 37, 49, 53], [-41, -61], [-49, -53], [-13, -29], [-49, -57], [-13, -45], [-49, -61], [-13, -61], [-49, -50], [-49, -51], [-49, -52], [-45, -57], [-53, -57], [-29, -45], [-53, -61], [-29, -61], [-53, -54], [-53, -55], [-53, -56], [-57, -61], [-45, -61], [-57, -58], [-57, -59], [-57, -60], [-61, -62], [-61, -63], [-61, -64], [49, 53, 57, 61], [13, 29, 45, 61], [41, 45, 57, 61], [-2, -22], [-2, -6], [-2, -18], [-2, -10], [-2, -34], [-2, -14], [-2, -50], [-2, -3], [-2, -4], [-6, -18], [-6, -10], [-18, -34], [-6, -14], [-18, -50], [-6, -7], [-6, -8], [-10, -14], [-34, -50], [-10, -11], [-10, -12], [-14, -15], [-14, -16], [2, 6, 10, 14], [2, 18, 34, 50], [2, 6, 18, 22], [-10, -30], [-18, -22], [-6, -22], [-18, -26], [-6, -38], [-18, -30], [-6, -54], [-18, -19], [-18, -20], [-14, -26], [-22, -26], [-22, -38], [-22, -30], [-22, -54], [-22, -23], [-22, -24], [-26, -30], [-38, -54], [-26, -27], [-26, -28], [-30, -31], [-30, -32], [18, 22, 26, 30], [6, 22, 38, 54], [10, 14, 26, 30], [-34, -54], [-34, -38], [-10, -26], [-34, -42], [-10, -42], [-34, -46], [-10, -58], [-34, -35], [-34, -36], [-38, -50], [-38, -42], [-26, -42], [-38, -46], [-26, -58], [-38, -39], [-38, -40], [-42, -46], [-42, -58], [-42, -43], [-42, -44], [-46, -47], [-46, -48], [34, 38, 42, 46], [10, 26, 42, 58], [34, 38, 50, 54], [-42, -62], [-50, -54], [-14, -30], [-50, -58], [-14, -46], [-50, -62], [-14, -62], [-50, -51], [-50, -52], [-46, -58], [-54, -58], [-30, -46], [-54, -62], [-30, -62], [-54, -55], [-54, -56], [-58, -62], [-46, -62], [-58, -59], [-58, -60], [-62, -63], [-62, -64], [50, 54, 58, 62], [14, 30, 46, 62], [42, 46, 58, 62], [-3, -23], [-3, -7], [-3, -19], [-3, -11], [-3, -35], [-3, -15], [-3, -51], [-3, -4], [-7, -19], [-7, -11], [-19, -35], [-7, -15], [-19, -51], [-7, -8], [-11, -15], [-35, -51], [-11, -12], [-15, -16], [3, 7, 11, 15], [3, 19, 35, 51], [3, 7, 19, 23], [-11, -31], [-19, -23], [-7, -23], [-19, -27], [-7, -39], [-19, -31], [-7, -55], [-19, -20], [-15, -27], [-23, -27], [-23, -39], [-23, -31], [-23, -55], [-23, -24], [-27, -31], [-39, -55], [-27, -28], [-31, -32], [19, 23, 27, 31], [7, 23, 39, 55], [11, 15, 27, 31], [-35, -55], [-35, -39], [-11, -27], [-35, -43], [-11, -43], [-35, -47], [-11, -59], [-35, -36], [-39, -51], [-39, -43], [-27, -43], [-39, -47], [-27, -59], [-39, -40], [-43, -47], [-43, -59], [-43, -44], [-47, -48], [35, 39, 43, 47], [11, 27, 43, 59], [35, 39, 51, 55], [-43, -63], [-51, -55], [-15, -31], [-51, -59], [-15, -47], [-51, -63], [-15, -63], [-51, -52], [-47, -59], [-55, -59], [-31, -47], [-55, -63], [-31, -63], [-55, -56], [-59, -63], [-47, -63], [-59, -60], [-63, -64], [51, 55, 59, 63], [15, 31, 47, 63], [43, 47, 59, 63], [-4, -24], [-4, -8], [-4, -20], [-4, -12], [-4, -36], [-4, -16], [-4, -52], [-8, -20], [-8, -12], [-20, -36], [-8, -16], [-20, -52], [-12, -16], [-36, -52], [4, 8, 12, 16], [4, 20, 36, 52], [4, 8, 20, 24], [-12, -32], [-20, -24], [-8, -24], [-20, -28], [-8, -40], [-20, -32], [-8, -56], [-16, -28], [-24, -28], [-24, -40], [-24, -32], [-24, -56], [-28, -32], [-40, -56], [20, 24, 28, 32], [8, 24, 40, 56], [12, 16, 28, 32], [-36, -56], [-36, -40], [-12, -28], [-36, -44], [-12, -44], [-36, -48], [-12, -60], [-40, -52], [-40, -44], [-28, -44], [-40, -48], [-28, -60], [-44, -48], [-44, -60], [36, 40, 44, 48], [12, 28, 44, 60], [36, 40, 52, 56], [-44, -64], [-52, -56], [-16, -32], [-52, -60], [-16, -48], [-52, -64], [-16, -64], [-48, -60], [-56, -60], [-32, -48], [-56, -64], [-32, -64], [-60, -64], [-48, -64], [52, 56, 60, 64], [16, 32, 48, 64], [44, 48, 60, 64]]
test_for('test_for_conj_sudoku : ',for_conj_sudoku(2),corrige_for2)
'''



def init_list_var(grille,n):
    '''
    Renvoie : une liste list_var initialisant une valuation tenant compte des valeurs non nulles déjà renseignées dans list_grille_complete
'''
'''
grille2= [0, 1, 0, 0, 4, 2, 0, 0, 0, 0, 2, 0, 0, 3, 0, 0]
cor_list_var_grille2= [None, None, None, None, True, False, False, False, None, None, None, None, None, None, None, None, False, False, False, True, False, True, False, False, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, False, True, False, False, None, None, None, None, None, None, None, None, False, False, True, False, None, None, None, None, None, None, None, None]
test('test_init_list_var : ',init_list_var(grille2,2),cor_list_var_grille2)
'''







'''#test ultim_resol
for2=[[1,4,-5],[-1,-5],[2,-3,5],[2,-4],[2,4,5],[-1,-2],[-1,2,-3],[-2,4,-5],[1,-2]]
list_var_for2=[None,None,None,None,None]
boo_for2,lilifor2=ultim_resol(for2,list_var_for2)
print('boo_for2 : ',boo_for2)
print('lilifor2 : ',lilifor2)'''


'''#test ultim_resol_simpl_for
#Cas grille Taille 2
formul_sudok2=for_conj_sudoku(2)
list_grille2=[[1,2,1],[2,1,4],[2,2,2],[3,3,2],[4,2,3]]
list_grille2_f=[[1,2,4],[2,1,4],[2,2,2],[3,3,2],[4,2,3]]
grille2=creer_grille_init(list_grille2,2)
afficher_grille(grille2,2)
list_var_grille2=init_list_var(grille2,2)
boo_2,lili2=ultim_resol_simpl_for(formul_sudok2,list_var_grille2)
#corrigé lili2=[False, False, True, False, True, False, False, False, False, False, False, True, False, True, False, False, False, False, False, True, False, True, False, False, False, False, True, False, True, False, False, False, True, False, False, False, False, False, False, True, False, True, False, False, False, False, True, False, False, True, False, False, False, False, True, False, True, False, False, False, False, False, False, True]
if boo_2:
    afficher_grille(creer_grille_final(lili2,2),2)
grille2f=creer_grille_init(list_grille2_f,2)
afficher_grille(grille2f,2)
list_var_grille2f=init_list_var(grille2f,2)
boo_2f,lili2f=ultim_resol_simpl_for(formul_sudok2,list_var_grille2f)
if boo_2f:
    afficher_grille(creer_grille_final(lili2f,2),2)'''

'''
#test ultim_resol_simpl_for
#Cas grille Taille 3
formul_sudok=for_conj_sudoku(3)
list_grille3=[[1,3,2],[1,6,5],[2,5,4],[2,8,9],[2,9,3],[3,2,7],[3,9,6],[4,3,1],[4,4,8],[4,8,3],[5,1,7],[5,2,2],[5,5,6],[5,8,8],[5,9,4],[6,2,4],[6,6,2],[6,7,5],[7,1,3],[7,8,1],[8,1,4],[8,2,6],[8,5,7],[9,4,9],[9,7,8]]
grille3=creer_grille_init(list_grille3,3)
afficher_grille(grille3,3)
list_var_grille3=init_list_var(grille3,3)
boo_3,lili3=ultim_resol_simpl_for(formul_sudok,list_var_grille3)
if boo_3:
    afficher_grille(creer_grille_final(lili3,3),3)
'''



'''#test ultim_resol_simpl_for_dpll cas3
formul_sudok=for_conj_sudoku(3)
list_grille3=[[1,3,2],[1,6,5],[2,5,4],[2,8,9],[2,9,3],[3,2,7],[3,9,6],[4,3,1],[4,4,8],[4,8,3],[5,1,7],[5,2,2],[5,5,6],[5,8,8],[5,9,4],[6,2,4],[6,6,2],[6,7,5],[7,1,3],[7,8,1],[8,1,4],[8,2,6],[8,5,7],[9,4,9],[9,7,8]]
grille3=creer_grille_init(list_grille3,3)
afficher_grille(grille3,3)
list_var_grille3=init_list_var(grille3,3)
boo_3,lili3=ultim_resol_simpl_for_dpll(formul_sudok,list_var_grille3)
if boo_3:
    afficher_grille(creer_grille_final(lili3,3),3)'''




