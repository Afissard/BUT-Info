# Importer des librairies  spécifiques

Pour éxécuter certains cas de tests, il peut être nécessaire d'ajouter des librairies à la configuration standard de gradle.

Pour cela, on va modifier les dépendances de gradle pour qu'il télécharge/installe les librairies nécessaires.


## Erreurs `"junit-params"` 

Si vous avez des erreurs concernant `"junit-params"`.

![junit params](junit-params.png)

Remplacez dans le fichier `build.gradle.kts` la dépendance :

	dependencies {
	    testImplementation(kotlin("test"))
	}

par 

	dependencies {
	    testImplementation("org.junit.jupiter:junit-jupiter:5.8.2")
	    testImplementation("org.junit.jupiter:junit-jupiter-params:5.8.2")
	}
	
Relancez le build de gradle en apppuyant sur 
![build gradle](gradle.png)


	
	
## Erreurs `"UMLChecker"`

Si vous avez des erreurs concernant `"UMLChecker"`

![uml checker](umlchecker.png)

il faut créer un dosssier `lib/` à la racine du projet ; puis copier-coller le fichier `univ.nantes.umlchecker-2.1.jar` (présent par exemple dans `dev.objets.tp4/exo01/lib/`)

Dans `build.gradle.kts`, ajoutez les dépendances suivantes :


	dependencies {
		...
		testImplementation("org.jetbrains.kotlin:kotlin-reflect:X.X.XX")
    	testImplementation("net.sourceforge.plantuml:plantuml:1.2023.5")
    	testImplementation(files("lib/univ.nantes.umlchecker-2.1.jar"))
	}
	
Les `X.X.XX`doivent correspondre à la version de Kotlin indiquée au début du fichier 
`build.gradle.kts` ; par exemple :

		plugins {
		    kotlin("jvm") version "1.6.10"
		    application
		}
	
	
Puis, relancez le build de gradle en apppuyant sur 
![build gradle](gradle.png)

