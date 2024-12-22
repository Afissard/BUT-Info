#!/bin/bash

if [ $# -eq 4 ] && printf '%u' $1 >&/dev/null && printf '%u' $2 >&/dev/null && printf '%u' $3 >&/dev/null
then
#nombres de dossiers et de fichiers et profondeur récupérés des arguments
    numdos=$1;
    numfil=$2;
    deptre=$3;
    arbo=$4;

#creation du repertoire dans lequel sera générée l'arborescence
    mkdir $arbo;
    cd $arbo;

#creation de deux fichiers .mp3 dans le repertoire arbo (ceci pour assurer que find -name *.mp3 ne fonctionne pas)
    touch bidon1.mp3
    touch bidon2.mp3

#initialisation de la profondeur courante, du nombre de dossiers déjà créés et du nombre de fichiers déjà créés
    curdep=0;
    curdos=0;
    curfil=0;
    
#noms possibles pour les fichiers et dossiers
    names=(untruc montruc letruc chose machin foo bar foobar divers something athing anything various stuff esohc);
    numnames=${#names[*]};
    
#contenus possibles pour les fichiers
    texts=(bonjour bonsoir salut hey hello hi allo hola yo ohayo);
    numtexts=${#texts[*]};
    
#extensions possibles pour les fichiers
    exts=(.txt .java .html .jpg .mp3);
    numexts=${#exts[*]};
    
#tant que pas assez dossiers/fichiers
    while [ $curfil -lt $numfil ] || [ $curdos -lt $numdos ]
    do
#si on est déjà à la profondeur maximale
	name=${names[$RANDOM % $numnames]};
	text=${texts[$RANDOM % $numtexts]};
	ext=${exts[$RANDOM % $numexts]};
#si on est à la profondeur maximale on ne peut créer que des fichiers (on vérifie quand même qu'on n'efface rien et qu'il reste des dossiers à créer)
	if [ $deptre -le $curdep ]
	then
	    if [ ! -e $name$ext ] && [ $curfil -lt $numfil ]
	    then
		echo $text > $name$ext;
		bourrage=$(( $RANDOM % 1000 ))
		for i in $(seq 1 $bourrage) ;  do
				echo $i >> $name$ext;
		done
		((curfil=curfil+1));
	    fi
#sinon on créé aléatoirement un fichier ou un dossier (sauf si le dossier à créer ou le fichier à créer existent déjà ou qu'on a déjà créé assez de fichiers ou de dossiers)
	else
	    if [ -e $name$ext ] || [ $numfil -le $curfil ]
	    then
		if [ ! -e $name ] && [ $curdos -lt $numdos ]
		then
		    mkdir $name;
		    ((curdos=curdos+1));
		fi
	    elif [ -e $name ] || [ $numdos -le $curdos ]
	    then
		echo $text > $name$ext;
		((curfil=curfil+1));
	    else
		if [ $[ $RANDOM % 2 ] -eq 0 ]
		then
		    echo $text > $name$ext;
		    ((curfil=curfil+1));
		else
		    mkdir $name;
		    ((curdos=curdos+1));
		fi
	    fi
	fi
#aléatoirement on se déplace parmi les dossiers
	if [ $[ $RANDOM % 2 ] -eq 0 ]
	then
	    numindos=`find . -mindepth 1 -maxdepth 1 -type d | wc -l`;
	    if [ $deptre -le $curdep ] || [ $numindos -le 0 ] 
	    then
		if [ $curdep -gt 0 ]
		then
		    cd ..;
		    ((curdep=curdep-1));
		fi
	    elif [ $curdep -le 0 ]
	    then
		    nextdos=`ls -d */ | shuf -n1`;
		    cd $nextdos;
		    ((curdep=curdep+1));
	    elif [ $[ $RANDOM % 4 ] -eq 0 ]
	    then
		cd ..;
		((curdep=curdep-1));
	    else
		nextdos=`ls -d */ | shuf -n1`;
		cd $nextdos;
		((curdep=curdep+1));
	    fi
	fi
    done
else
    echo Arguments manquants ou incorrects; 
    echo gen.sh numdos numfil deptre arbo;
    echo Crée le dossier arbo, puis génère dans celui ci une arborescence de fichiers contenant numdos dossiers, numfil fichiers et une profondeur maximale de deptre.;
fi
