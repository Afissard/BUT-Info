---
title: Untitled
draft: 
description: 
tags:
---
# graphe de flux (invalide)
```mermaid
graph TD
    A[Début] --> B[Planification des examens]
    B --> C[Réservation des salles]
    B --> D[Sollicitation des enseignants pour les sujets]
    D --> E[Réception et vérification des sujets]
    E --> F[Archivage des sujets]
    F --> G[Convocation des étudiants]
    G --> H[Convocation des surveillants]
    H --> I[Distribution des sujets aux surveillants]
    I --> J[Jour de l'examen]
    J --> K[Distribution des sujets aux étudiants]
    K --> L[Récupération des copies par les surveillants]
    L --> M[Transmission des copies à la scolarité]
    M --> N[Transmission des copies aux enseignants]
    N --> O[Correction par les enseignants]
    O --> P[Renvoi des notes corrigées à la scolarité]
    P --> Q[Saisie des notes]
    Q --> R[Affichage des notes]
    R --> S[Consultation des copies par les étudiants]
    S --> T{Réclamation ?}
    T -- Oui --> U[Révision des copies par les enseignants]
    T -- Non --> V[Clôture du processus]
    U --> Q
    V --> Fin[Processus terminé]

```
# modèle de flux actuel
```mermaid
graph TB
	%% Styles
	classDef external fill:#d55,stroke:#333,stroke-width:2px
	classDef internal fill:#55d,stroke:#000,stroke-width:1px

	%% Entités externes
	%%Enseignants[Enseignants Vacataires]:::external
	%%Surveillants[Surveillants]:::external
	Enseignants([Enseignants Vacataires]):::external
	Surveillants([Surveillants]):::external

	%% Sous-systèmes du système interne
	Scolarite[Service Scolarité]:::internal
	Correction[Correction des copies]:::internal
	Notes[Gestion des notes]:::internal
	%%Jury[Jury et Clôture]:::internal
	Etudiants[Étudiants]:::internal

	%% Flux de données
	Etudiants -->|Convocations| Scolarite
	Surveillants -->|Disponibilités| Scolarite
	Enseignants -->|Sujets d'examen| Scolarite
	Scolarite -->|Sujets corrigés| Enseignants
	Scolarite -->|Copies d'examen| Correction
	Correction -->|Notes corrigées| Notes
	Notes -->|Notes finales| Etudiants
	Etudiants -->|Réclamations| Notes
	Notes -->|Décisions sur réclamations| Enseignants
	%%Notes -->|Synthèse des résultats| Jury
	%%Jury -->|Validation finale| Scolarite
```
**Interne :** Bleu
**Externe :** Rouge
# modèle de flux actuel V2
```mermaid
graph LR
	%% Styles
	classDef external fill:#d55,stroke:#333,stroke-width:2px
	classDef internal fill:#55d,stroke:#000,stroke-width:1px

	%% Entités externes
	%%Enseignants[Enseignants Vacataires]:::external
	%%Surveillants[Surveillants]:::external
	Enseignants([Enseignants Vacataires]):::external
	Surveillants([Surveillants]):::external

	%% Sous-systèmes du système interne
	Scolarite[Service Scolarité]:::internal
	Evaluation[Processus d'évaluation]:::internal
	%%Correction[Correction des copies]:::internal
	%%Notes[Gestion des notes]:::internal
	Etudiants[Étudiants]:::internal

	%% Flux de données
	Scolarite -->|Convocations| Etudiants
	Surveillants -->|Disponibilités| Scolarite
	Enseignants -->|Sujets d'examen| Scolarite
	%%Scolarite -->|Sujets corrigés| Enseignants
	Scolarite -->|Copies d'examen| Evaluation
	Etudiants -->|Rend les copies| Scolarite
	Scolarite -->|Envois des copies| Enseignants
	Enseignants -->|Copies corrigées| Evaluation
	Enseignants -->|Modification des notes| Evaluation
	Evaluation -->|Notes finales| Etudiants
	Etudiants -->|Réclamations| Evaluation
	Evaluation -->|Décisions sur réclamations| Enseignants
	%%Notes -->|Synthèse des résultats| Jury
	%%Jury -->|Validation finale| Scolarite
```
**Interne :** Bleu
**Externe :** Rouge

# modèle de flux révisé

```mermaid
graph TB
	%% Styles
	classDef external fill:#d55,stroke:#333,stroke-width:2px
	classDef internal fill:#55d,stroke:#000,stroke-width:1px

	%% Entités externes
	Enseignants([Enseignants Vacataires]):::external
	Surveillants([Surveillants]):::external

	%% Entités internes
	Scolarite[Scolarité]:::internal
	Etudiant[Etudiant]:::internal

	subgraph Plateforme_d'examen_numérique
        DSujets[Gestion des sujets]:::internal
        DCorr[Système de correction]:::internal
        DNotes[Saisie des notes]:::internal
        DCopies[Copies numériques]:::internal
    end

	%% Flux
    Enseignants -->|Soumission du sujet| DSujets
    DSujets -->|Validation automatique| Scolarite

    Scolarite -->|Convocation numérique| Etudiant
    Etudiant -->|Confirmation via plateforme| Scolarite

    Scolarite -->|Affectation automatisée| Surveillants
    Surveillants -->|Confirmation de disponibilité| Scolarite

    Etudiant -->|Rendu des copies physique ou numérique| DCopies
    DCopies -->|Transmission numérique des copies| Enseignants
    Enseignants -->|Annotation des copies| DCorr
    DCorr -->|Transmission des notes| DNotes

    DNotes -->|Accès aux notes en ligne| Etudiant
    Etudiant -->|Demande de révision| Enseignants
    Enseignants -->|Modification des notes| DNotes
```
# modèle de flux révisé V2
```mermaid
graph TB
	%% Styles
	classDef external fill:#d55,stroke:#333,stroke-width:2px
	classDef internal fill:#55d,stroke:#000,stroke-width:1px

	%% Entités externes
	Enseignants([Enseignants Vacataires]):::external
	Surveillants([Surveillants]):::external

	%% Entités internes
	subgraph Plateforme_d'examen_numérique
		Scolarite[Scolarité]:::internal
		Etudiant[Etudiant]:::internal
        DSujets[Gestion des sujets]:::internal
        DCorr[Système de correction]:::internal
        DNotes[Saisie des notes]:::internal
        DCopies[Copies numériques]:::internal
    end

	%% Flux
    Enseignants -->|Soumission du sujet| DSujets
    DSujets -->|Validation automatique| Scolarite

    Scolarite -->|Convocation numérique| Etudiant
    Etudiant -->|Confirmation via plateforme| Scolarite

    Scolarite -->|Affectation automatisée| Surveillants
    Surveillants -->|Confirmation de disponibilité| Scolarite

    Etudiant -->|Rendu des copies physique ou numérique| DCopies
    DCopies -->|Transmission numérique des copies| Enseignants
    Enseignants -->|Annotation des copies| DCorr
    DCorr -->|Transmission des notes| DNotes

    DNotes -->|Accès aux notes en ligne| Etudiant
    Etudiant -->|Demande de révision| Enseignants
    Enseignants -->|Modification des notes| DNotes
```
