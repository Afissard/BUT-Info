import pandas as pd
import numpy as np
import math as mt
import matplotlib.pyplot as plt
from scipy.stats import *
import geopandas as gpd


nbLineShowed = 10
def showDF(df, nbLine=nbLineShowed):
    if nbLine <= 0 :
        print(df,'\n')
    else :
        print(df.head(nbLine),"\n\n")

print("geojson loaded")
path='./'
france = gpd.read_file(path+"departements.geojson")
world = gpd.read_file(path+"countries.geojson")

print("csv loaded")
data_communes = pd.read_csv(path+'RGC_2013.csv',sep=';', dtype=str)
data_emplois=pd.read_csv(path+'Emplois_salaries.csv',sep=';', dtype=str)
#Traitement préalable des données à effectuer
data_communes.DEP.astype(str)
data_communes.COM.astype(str)
data_emplois.astype(str)
data_communes.insert(loc=2, column="COG", value=data_communes.DEP+data_communes.COM.str.zfill(3))

def associate_arr_mun(num_dep, cog, num_arr_debut, num_arr_fin):
    list_arr_com=[str(num_dep)+str(num_arr).zfill(3) for num_arr in range(num_arr_debut,num_arr_fin+1)]
    mask=data_emplois['CODGEO'].str.zfill(5).isin(list_arr_com)
    data_emplois_commune=data_emplois[mask].sum()
    data_emplois_commune['CODGEO']=str(cog)
    data_emplois.loc[len(data_emplois)]=data_emplois_commune.values

def q3_2():
    """
    moyenne = df.mean()
    médiane = df.median()
    écart_type = df.std()
    
    On s'intéresse au ratio de l’effectif salarié par habitant sur l’ensemble des communes 
    (on supprimera pour cette seule question les communes aux populations nulles et les communes 
    pour lesquelles le ratio calculé renvoie la valeur nan). Réaliser le graphique ci-dessous 
    (plt.boxplot). Pouvez-vous expliquer la différence significative entre la moyenne et la médiane ?
    """
    temp_commune = data_communes[data_communes['POPU']!= 0][['COG', 'POPU', 'SURFACE']]
    temp_emplois = data_emplois[['CODGEO','EFF_TOT']]
    showDF(temp_commune)
    showDF(temp_emplois)
    ratio = pd.merge(left=temp_commune, right=temp_emplois, how='inner', left_on=['COG'], right_on=['CODGEO'])[['COG', 'POPU', 'EFF_TOT', 'SURFACE']]
    ratio.insert(loc=len(ratio.columns), column="RATIO", value=ratio.EFF_TOT.astype(float)/ratio.POPU.astype(float))
    ratio.insert(loc=len(ratio.columns), column='DENS', value=(ratio.POPU.astype(float)*100)/(ratio.SURFACE.astype(float)/100))
    showDF(ratio)
    plt.boxplot(x=ratio['RATIO'], showmeans=True, meanline=True, showfliers=False)
    plt.show()
    
    
# * DENS (Hab/Km²) = (POPU*100)/(SURFACE/100) 

def main():
    print("setup done, starting")
    # showDF(data_emplois)
    # showDF(data_communes)
    # print(data_communes.COG)
    
    associate_arr_mun(75,75056,101,120)
    associate_arr_mun(13,13055,201,216)
    associate_arr_mun(69,69123,381,389)
    
    q3_2()
    
if __name__=="__main__":
    main()