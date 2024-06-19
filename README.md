# Elephant Carpacio


## Lien et énoncé 

[Extreme carpaccio](https://diegolemos.net/2016/01/07/extreme-carpaccio/)

## Backlog

Fonctionnalités à ajouter :

- Donner une liste d'états
- Proposer un nombre d'articles disponibles
Un seul article :
    - nombre disponibles
    - prix
- Ajouter une taxe de 10%
- Les articles diffèrent selon l'état


But final : 
- Choisir un Etat dans une liste
- Lister les articles d'un Etat

## Attendu
./elephant state

["UTAH"]
 
./elephant state --name=UTAH articles

[ { "name": "banana", "unitPrice": "270.99", "unit": "dollar", "availableItems": "1000", "taxe": "0.1" } ]
 
 
./elephant state --name=UTAH buy --name="banana" --numberOfItems=978 

291531,04 ((270.99*978)*1.1)
 
 
./elephant state --name=UTAH buy --name="banana" --numberOfItems=1001

NOT POSSIBLE
 
Où tu penses qu'on arriver en 5 itérations de 10min ?  
