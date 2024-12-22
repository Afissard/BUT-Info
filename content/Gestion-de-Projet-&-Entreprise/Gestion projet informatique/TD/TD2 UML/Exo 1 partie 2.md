# Traiter passage en caisse
| Catégorie                   | Réponse   |
| ------------------- | --- |
| Responsabilité      | Caissier    |
| Acteur              | Client    |
| Déclancheur         | Client en présance d'article    |
| pré-condition       | Le caissier s'est authentifié    |
| post-condiction     | le client est servi, les stock sont à jours    |
| Scénario nominal    | ?     |
| Scénario alternatif | ?    |
| Scénario exeptionnel                    | Le client ne peut pas payer et doit partir sans ses article pour tout les cas possible selon la loi de Murphi    |
# Effectuer paiement
| Catégorie            | Réponse                                |
| -------------------- | -------------------------------------- |
| Responsabilité       | Caissier                               |
| Acteur               | Client                                 |
| Déclancheur          | Demande de paiement                    |
| pré-condition        | ?                                      |
| post-condiction      | le paiement est validé                 |
| Scénario nominal     | le client choisi sont mode de paiement |
| Scénario alternatif  | le client demande une réduction                                       |
| Scénario exeptionnel | paiement refusé pour tout les cas possible selon la loi de Murphi                                       |
# Annuler paiement
| Catégorie | Réponse |
| ---- | ---- |
| Responsabilité | Manageur |
| Acteur | Manageur |
| Déclancheur | Erreur du caissier |
| pré-condition | passage en caisse en cours |
| post-condiction | erreur annulé et reprise normal du passage en caisse |
| Scénario nominal | le manageur s'authentifi et annule l'erreur |
| Scénario alternatif |  |
| Scénario exceptionnel | le manageur ne peut annulé l'erreur pour tout les cas possible selon la loi de Murphi |
...