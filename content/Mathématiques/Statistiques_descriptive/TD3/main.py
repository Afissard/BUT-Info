import pandas as pd, numpy as np, matplotlib, matplotlib.pyplot as plt
from scipy.stats import linregress


nbLineShowed = 10
def showDF(df, nbLine=nbLineShowed):
    if nbLine <= 0 :
        print(df,'\n')
    else :
        print(df.head(nbLine),"\n\n")

def showAllDF():
    dPlayer = pd.read_csv('data_players.csv', sep=',')
    dScores = pd.read_csv('match_scores.csv', sep=',')
    dStats = pd.read_csv('match_stats.csv', sep=',')
    dTournaments = pd.read_csv('tournaments.csv', sep=',')
    dRanking1 = pd.read_csv('rankings.csv', sep=',')
    dRanking2 = pd.read_csv('rankings_2.csv', sep=',')
    dRanking3 = pd.read_csv('rankings_3.csv', sep=',')
    
    print("Player\n")
    showDF(dPlayer)
    print("Scores\n")
    showDF(dScores)
    print("Stats\n")
    showDF(dStats)
    print("Tournament\n")
    showDF(dTournaments)
    print("Ranking\n")
    showDF(dRanking1)
    showDF(dRanking2)
    showDF(dRanking3)


def g1():
    dPlayer = pd.read_csv('data_players.csv', sep=',')
    
    grandJoueurs = dPlayer[(dPlayer.height_cm > 6.5) & (dPlayer.birth_year > 1960.0)]['height_cm']
    print(grandJoueurs.describe())
    
    bmin = grandJoueurs.min()
    bmax = grandJoueurs.max()
    liBo = [i for i in range(int(bmin), int(bmax), 5)]
    print(liBo)
    plt.hist(x=grandJoueurs, bins=liBo,ec='black')
    plt.title("Distribution des tailles des joueur nées à partir de 1960")
    plt.show()

def g2():
    """
    Présenter un graphique montrant l’ ́evolution du classement d’un
    joueur dont on aura saisi et enregistré l’identifiant dans une variable
    """
    pID = 's424'
    
    # Panda
    dRanking1 = pd.read_csv('rankings.csv', sep=',')
    dRanking2 = pd.read_csv('rankings_2.csv', sep=',')
    dRanking3 = pd.read_csv('rankings_3.csv', sep=',')
    # showDF(dRanking1)
    # showDF(dRanking2)
    # showDF(dRanking3)
    
    pRank1 = dRanking1[(dRanking1.player_id == pID)][['week_title', 'rank_number']]
    pRank2 = dRanking2[(dRanking2.player_id == pID)][['week_title', 'rank_number']]
    pRank3 = dRanking3[(dRanking3.player_id == pID)][['week_title', 'rank_number']]
    pRank = pd.concat([pRank1, pRank2, pRank3])
    
    pRank.sort_values(by='week_title', inplace=True, ascending=True)
    showDF(pRank, 0)
    
    # Pyplot
    plt.gca().invert_yaxis()
    plt.plot(pRank.week_title.astype('datetime64[ns]'), pRank.rank_number)
    plt.show()

def g3():
    """
    On veut réaliser une  ́etude quant à une corr ́elation potentielle entre la taille
    d’un joueur et le taux de points gagnés suite à un service sur première balle réussi.
    """
    

if __name__ == "__main__":
    debug = False
    if debug :
        showAllDF()

    # g1()
    g2()