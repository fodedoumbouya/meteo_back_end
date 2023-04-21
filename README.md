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

# API Endpoints

Voici les endpoints disponibles dans l'API de ce projet :

    /api/station : retourne la liste de toutes les stations météorologiques sous forme d'un JSON contenant l'identifiant de chaque station, son numéro de série, son emplacement, son modèle et ses coordonnées géographiques (latitude et longitude).

    /api/widget?id=id_widget : retourne le code HTML d'un widget pour afficher les données météorologiques d'une station spécifiée par son identifiant.

    /api/weather/measurements?id=id_widget : retourne le temps qu'il fait dans la station spécifiée par son identifiant sous forme d'un JSON contenant un message de succès, un code de statut et une chaîne de caractères décrivant les conditions météorologiques actuelles (par exemple "ensoleillé", "pluvieux", "orageux", etc.).

Bonne utilisation !