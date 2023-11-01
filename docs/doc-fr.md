# Générateur de carnets de tombola

# Objectif
Générez facilement des carnets de tombola personnalisables en utilisant une interface en ligne de commande.
* 3 tickets par page
* 10 tickets par carnet
# Comment l'utiliser ?
## Commande
### Choisissez le binaire approprié en fonction de votre système d'exploitation
* Windows : `bin/generate-raffle-book-amd64-windows`
* Linux : `bin/generate-raffle-book-amd64-linux`
* MacOS : `bin/generate-raffle-book-amd64-darwin`

### Fichier PDF
Le fichier généré se trouve dans le chemin `build/tickets.json`

### Standard (adaptez à votre système d'exploitation)
```bash
bin/generate-raffle-book-amd64-windows
```
Génère 12 carnets, soit 120 tickets

Paramètres supplémentaires pour la numérotation personnalisée
bash
Copy code
```bash
bin/generate-raffle-book-amd64-windows -start=120 -count=100
```
Génère 100 carnets, en commençant à 120.

Le début et le compte sont arrondis au nombre supérieur le plus proche (commence à 121, crée 12 carnets, soit 120 tickets) pour maintenir la continuité de la numérotation.

Comment personnaliser ?
Ce que vous ne pouvez pas personnaliser
Tous les textes sont personnalisables à l'exception des numéros (N° 0XXXX)
![Not customizable](docs/images/not-customizable.png)


Ce que vous pouvez personnaliser
![Customizabe](docs/images/customizable.png)


Comment le faire
Copiez `assets/customs.json.dist` vers `assets/customs.json`

```bash
make customs
```
Maintenant, modifiez le fichier `assets/customs.json` pour qu'il corresponde à vos besoins.

Exemple 1 : mettez à jour les éléments de texte
Respectez le format JSON (utilisez [cet outil](https://jsonformatter.curiousconcept.com/) par exemple)

```json
{
    "left_title": [
        "LYCÉE",
        "LYON BELLECOUR"
    ],
    "right_title": [
        "LYCÉE LYON BELLECOUR",
        "SOUSCRIPTION"
    ],
    "event_description": [
        "Date du tirage",
        "31/10/2023 Lyon Bellecour"
    ],
    "prize_title": [
        "Liste des prix"
    ],
    "prize_list": [
        "* 1er prix",
        "* 2e prix",
        "* 3e prix"
    ],
    "price": "Prix : 1€",
    "acheteur": "Nom de l'acheteur",
    "vendeur": "Nom du vendeur"
}
```