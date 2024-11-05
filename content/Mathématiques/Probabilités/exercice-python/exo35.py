import numpy as np, pandas as pd, scipy.stats as scs, matplotlib.pyplot as plt

# init
df = pd.ExcelFile('video_etu.xlsx')
dFilms = pd.read_excel(df, 'films')
dureeFilms = dFilms['FILM_DUREE']

def showGeneralInfo():
    print(dFilms)
    # print(dureeFilms)
    # print("moyenne\t\t", dureeFilms.mean())
    # print("écart type\t", dureeFilms.std())
    print(dureeFilms.describe())
    
def q1(nbEchantillons=1000):
    rChoices = [np.random.choice(dureeFilms, size=50, replace=False) for _ in range(nbEchantillons)]
    # print("mean\t\tstd")
    # for i in rChoices :
    #     print(i.mean(),"  \t", i.std()) 
    
    means = sorted([rChoices[i].mean() for i in range(nbEchantillons)], key = lambda x:float(x))
    # plt.bar(height=[i for i in range(nbEchantillons)], x=means)
    
    plt.style.use('dark_background')
    fig, ax = plt.subplots()
    ax.hist(means, bins=50)
    ax.plot()
    plt.show()

def q2(nbEchantillons=1000):
    rChoices = [np.random.choice(dureeFilms, size=50, replace=False) for _ in range(nbEchantillons)]
    means = sorted([rChoices[i].mean() for i in range(nbEchantillons)], key = lambda x:float(x))
    res = [scs.norm.ppf(means[i]) for i in range(nbEchantillons)]
    print(res)
    # plt.hist(res, bins=50)
    

if __name__ == "__main__":
    # showGeneralInfo()
    # q1()
    q1(10000)
    # q2(10000)