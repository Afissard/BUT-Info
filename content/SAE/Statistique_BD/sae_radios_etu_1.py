import numpy as np
import pandas as pd
import scipy.stats as st

path=''
###PARTIE 1###
reponses=pd.read_csv(path+'reponses_radio_sauv.csv',sep=';')

def calcul_eff_obs():
    return #un dataframe

res_list_electro=[12,11,39,42,18,19]
res_list_radio3=[39, 46, 14, 19, 11, 31, 33, 28]
eff_obs=calcul_eff_obs()
print('test calcul_eff_obs : ',len(res_list_electro)-list(np.isclose(eff_obs.loc['Electro'],res_list_electro)).count(True)==0)
print('test calcul_eff_obs : ',len(res_list_radio3)-list(np.isclose(eff_obs['Radio 3'],res_list_radio3)).count(True)==0)

def calcul_eff_theo():
    return #un dataframe

res_list_indie=[5.187, 9.804, 12.597, 7.296, 9.69, 12.426]
res_list_radio6=[30.738, 35.534, 12.426, 18.094, 12.208, 40.112, 35.97, 32.918]
eff_theo=calcul_eff_theo()
print('test calcul_eff_theo : ',len(res_list_indie)-list(np.isclose(eff_theo.loc['Indie'].to_list(),res_list_indie)).count(True)==0)
print('test calcul_eff_theo : ',len(res_list_radio6)-list(np.isclose(eff_theo['Radio 6'].to_list(),res_list_radio6)).count(True)==0)

def calcul_contrib():
    return #un dataframe

res_list_rock=[1.6550266400266396, 1.1129105003523614, 0.329253393665159, 1.7734090909090914, 0.8735294117647057, 0.9908507089241031]
res_list_radio1=[0.05381973345803127, 7.911675925301693, 0.2716346635820321, 11.815941877399705, 62.902907378335954, 8.237072145246058, 1.6550266400266396, 4.360896659631758]
contrib=calcul_contrib()
print('test calcul_contrib : ',len(res_list_rock)-list(np.isclose(contrib.loc['Rock'].to_list(),res_list_rock)).count(True)==0)
print('test calcul_contrib : ',len(res_list_radio1)-list(np.isclose(contrib['Radio 1'].to_list(),res_list_radio1)).count(True)==0)

def analyse_contrib(n):
    return #une liste de tuples

ntest=5
ana_contrib=analyse_contrib(ntest)

res_list=[('Musique classique', 'Radio 1', '+', 62.902907378335954), ('Electro', 'Radio 4', '+', 31.78736170212767), ('Variété', 'Radio 2', '+', 30.24675743107962), ('Pop', 'Radio 6', '+', 18.023647387315513), ('Hip-Hop & RnB', 'Radio 2', '-', 17.320063347125128)]
test_contrib=(list(np.isclose([res_list[i][3] for i in range(ntest)],[ana_contrib[i][3] for i in range(ntest)])).count(True)==ntest)
test_sens_dep=([res_list[i][3]==ana_contrib[i][3] for i in range(ntest)].count(True)==ntest)
test_radios=([res_list[i][1]==ana_contrib[i][1] for i in range(ntest)].count(True)==ntest)
print('test analyse_contrib : ',test_contrib and test_sens_dep and test_radios)

###PARTIE 2###
data_audience = pd.read_csv(path+'audience_cumulee_radio.csv',sep=';')
def decomp_serie():
    return #serie_mm,serie_coef_sais

res_moy_mob=[np.nan, np.nan, 80.0, 79.9375, 79.9, 79.7, 79.4375, 79.3375, 79.2125, 78.9875, 78.7, 78.35, 77.9125, 77.575, 77.4375, np.nan, np.nan, np.nan, np.nan, np.nan, 73.35, 73.26249999999999, 73.225, 73.1375, 72.7375, 71.925, 71.35, 71.125, 70.7375, 70.2375, np.nan, np.nan]
coef_sais={}
coef_sais['T2']=0.992730
coef_sais['T4']=1.006176
serie_mm,serie_coef_sais=decomp_serie()
print('test decomp_serie moy mob : ',list(np.isclose(np.nan_to_num(serie_mm.values),np.nan_to_num(np.array(res_moy_mob)))).count(True)==len(res_moy_mob) )
print('test decomp_serie coef sais : ',np.isclose(serie_coef_sais.loc[2],coef_sais['T2']) and np.isclose(serie_coef_sais.loc[4],coef_sais['T4']))

def prevision(annee_prev,trim_prev,serie_coef_sais):
    return #une valeur

prev=prevision(2024,4,serie_coef_sais)
res=68.7763620276032
print('test prévision : ',np.isclose(prev,res))