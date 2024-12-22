import numpy as np
import math


class GrapheValue:
    """
    Classe qui représente des graphes valués non orientés.
    Les graphes valués sont représentés par leur matrice de valuation.
    Les sommets sont numérotés de 0 à n-1 (n étant le nombre de sommets).
    """
    
    def __init__(self, mat):
        """
        Constructeur d'un graphe valué, à partir de sa matrice de valuation.
        Paramètres :
            m : matrice de valuation du graphe.
        """
        self.matrice = mat
 

    def __str__(self):
        """
        Représentation du graphe valué, par une chaîne de caractères.
        Retour :
            chaîne de caractères contenant les valeurs de la matrice 
            de valuation du graphe.
        """
        return str(self.matrice)
    
    
    def nb_sommets(self):
        """
        Calcul le nombre de sommets du graphe.
        Retour : 
            nombre de sommets du graphe.
        """
        return 0
    
    
    def nb_aretes(self):
        """
        Calcul le nombre d'arêtes du graphe.
        Retour : 
            nombre d'arêtes du graphe.
        """
        return 0
    
    
    def degre_sommet(self, s:int):
        """
        Calul du degré du sommet d'indice donné.
        Paramètres :
            s : indice du sommet considéré.
        Retour : 
            degré du sommet s.
        """
        return 0


    def degres_sommets(self):
        """
        Calcul de la liste des degrés des sommets du graphe.
        Retour : 
            liste des degrés des sommets du graphe.
        """
        return []
    
    
    def est_connexe(self):
        """
        Test de la connexité du graphe courant.
        Retour : 
            vrai si le graphe est connexe ; faux sinon.
        """
        return True
        
        
    def plus_grosse_cc(self):
        """
        Calcule les composantes connexes du graphe et retourne le sous-graphe 
        correspondant à la plus grosse d'entre elles (en termes de nombre de sommets).
        Retour :
            le sous-graphe correspondant à la plus grosse composante connexe 
            (la numérotation des sommets n'est plus la même que dans le graphe de départ).
        """
        return GrapheValue(np.array([]))



if __name__ == "__main__":
    m = np.array([[math.inf, math.inf,math.inf],
                  [math.inf,math.inf,math.inf],
                  [math.inf,math.inf,math.inf],
                 ])
    g = GrapheValue(m)
    print(g)
    
    m2 = np.array([[math.inf, 3, 2.5, 8],
                   [3,math.inf,math.inf,7],
                   [2.5,math.inf,math.inf,1.5],
                   [8,7,1.5,math.inf],
                  ])
    g2 = GrapheValue(m2)
    print(g2)
    print("degré(0) :", g2.degre_sommet(0))
    print("degré(1) :", g2.degre_sommet(1))
    print(g2.degres_sommets())
    print("Nb sommets :", g2.nb_sommets())
    print("Nb arêtes :", g2.nb_aretes())