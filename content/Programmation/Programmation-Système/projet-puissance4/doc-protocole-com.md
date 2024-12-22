---
title: doc-protocole-com
draft: 
description: 
tags:
  - Programmation
  - SAE
---
# Protocole communication
Chaque message est précédé par un byte de distinction : 
```go
OctetPlayerColor   byte = 0x01 // Player color selection  
OctetOpponentMove  byte = 0x02 // Opponent's move  
OctetOpponentState byte = 0x03 // Game state updates
```
Le contenu "utile" du message est envoyé dans le byte suivant, un message entier est donc un array de byte.
Par exemple, quand le joueur 1 envoi la position 4 qu'il joue, il envoi le message suivant : `[0x02, 0x04]`.
# Diagramme déroulement du jeu
```mermaid
sequenceDiagram
    actor C1 as Client-1
    actor C2 as Client-2
    participant S as Server

	%% connection
	Note left of C1 : Connection
	C1->>S: connexion
	C2->>S: connexion
	
	create participant H1 as Handle-client-1
	S->>H1: donne connection client-1
    
	create participant H2 as Handle-client-2
	S->>H2: donne connection client-2
	
	S->>H1: pret
	S->>H2: pret
	H1->>C1: pret
	H2->>C2: pret

	%% couleur
	Note left of C1 : Choix couleur
	loop tant que les deux joueur n'ont pas choisis leurs couleurs
		C1->>H1: couleur
		H1->>S: couleur
		S->>H2: couleur
		H2->>C2: couleur
	end

	%% boucle de jeu
	Note left of C1 : Boucle du jeu
	loop tant qu'un des deux joueur n'a pas perdu
		C1->>H1: position joué
		H1->>H2: position joué
		H2->>C2: position joué
		C2->>H2: position joué
		H2->>H1: position joué
		H1->>C1: position joué
	end

```

# Notes
[doc mermaid séquence-diagram](http://mermaid.js.org/syntax/sequenceDiagram.html)
