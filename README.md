# meteo_back_end
# Backend de l'application de visualisation de données météorologiques
# Configuration de l'environnement

Avant de pouvoir utiliser le backend, vous devez vous assurer que votre environnement dispose  des éléments suivants :

    Git : vous pouvez l'installer depuis le lien suivant : https://git-scm.com/downloads
    GoLang version go1.19.1 ou plus récente : vous pouvez l'installer depuis le lien suivant : https://golang.org/dl/
    Un serveur MySQL : vous pouvez utiliser WAMP Server qui contient MySQL. Si vous préférez installer MySQL directement, vous pouvez le télécharger depuis le lien suivant : https://dev.mysql.com/downloads/

Vous aurez également besoin d'un éditeur de texte ou d'un IDE comme Visual Studio Code. Vous pouvez l'installer depuis le lien suivant : https://code.visualstudio.com/download

# Installation

Voici les étapes à suivre pour installer et exécuter le backend :

    Clonez le projet en utilisant la commande suivante : git clone https://github.com/fodedoumbouya/meteo_back_end
    Lancez le serveur MySQL
    Exécutez le script SQL attaché au projet pour créer la base de données en local
    Ouvrez le dossier du projet dans Visual Studio Code
    Exécutez la commande go mod tidy pour installer les dépendances
    Exécutez la commande go run main.go pour lancer le serveur Go

N'oubliez pas de modifier les informations de connexion à la base de données dans le fichier config.go avant de lancer le serveur.

Bonne utilisation !