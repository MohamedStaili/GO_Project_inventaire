
# Projet de Gestion d'Inventaire IT

## Description
Ce projet est une application de gestion d'inventaire d'équipements informatiques pour une entreprise. Il permet de gérer les informations relatives aux équipements, tels que les numéros de série, les références d'achat, les fournisseurs, et plus encore. L'application est construite avec **Go** pour le backend et **React** pour le frontend.

### Fonctionnalités principales :
- Gestion des équipements (ajout, modification, suppression)
- Suivi des informations liées aux achats (référence d'achat, numéro de facture, prix, fournisseur)
- Gestion des utilisateurs (admin et employés)
- Système d'authentification sécurisé pour les administrateurs
- Interface utilisateur dynamique et responsive

## Prérequis

Avant de commencer, assurez-vous d’avoir les éléments suivants installés :

- [Go](https://golang.org/dl/) (version 1.18 ou supérieure)
- [Node.js](https://nodejs.org/) (version 14 ou supérieure)
- [npm](https://www.npmjs.com/)
- [Git](https://git-scm.com/)

## Installation

### Backend (Go)

1. Clonez le dépôt du backend :

```bash
git clone <url_du_dépôt_backend>
cd <dossier_backend>
```

2. Installez les dépendances Go :

```bash
go mod tidy
```

3. Démarrez le serveur backend :

```bash
go run cmd/web/main.go
```

Le backend sera disponible sur `http://localhost:8080`.

### Frontend (React)

1. Clonez le dépôt du frontend :

```bash
git clone <url_du_dépôt_frontend>
cd <dossier_frontend>
```

2. Installez les dépendances Node.js :

```bash
npm install
```

3. Démarrez l'application React :

```bash
npm start
```

L'application frontend sera disponible sur `http://localhost:3000`.

## Architecture

### Backend

Le backend est développé en Go et suit une architecture **MVC** (Modèle-Vue-Contrôleur).

- **Models** : Contient les structures et interactions avec la base de données.
- **Controllers** : Gère la logique de l'application et les requêtes HTTP.
- **Routes** : Définit les points d'entrée de l'API et les routes associées.
- **Middleware** : Implémente les vérifications d'authentification et autres fonctionnalités de sécurité.

### Frontend

Le frontend est une application React qui utilise des composants pour afficher dynamiquement les données de l'inventaire. Il communique avec l'API backend pour récupérer et afficher les données.

- **Pages** : Contient les différentes vues (Tableau de bord, Gestion des équipements, etc.).
- **Composants** : Composants réutilisables pour afficher les informations.
- **Services** : Gère les appels API et l'état des données.

## Authentification

L'application implémente un système d'authentification basé sur **JWT** pour sécuriser l'accès aux pages réservées aux administrateurs. Les utilisateurs doivent se connecter via le formulaire de connexion pour obtenir un token d'authentification.

## Développement

Si vous souhaitez contribuer au développement, voici les étapes de base pour configurer un environnement de développement local :

1. Clonez le projet.

2. Installez les dépendances du backend et du frontend.

3. Configurez la base de données pour le backend (par exemple, en utilisant MySQL ou PostgreSQL).

4. Lancez le backend et le frontend séparément.

5. Commencez à développer et à tester les nouvelles fonctionnalités.

## Tests

Des tests unitaires sont disponibles pour certaines parties du code, principalement pour la logique backend. Vous pouvez les exécuter en utilisant la commande suivante dans le dossier du backend :

```bash
go test ./...
```

## Contribuer

Les contributions sont les bienvenues ! Si vous avez des suggestions ou souhaitez résoudre un bug, veuillez suivre ces étapes :

1. Fork ce dépôt.
2. Créez une nouvelle branche (`git checkout -b feature-nouvelle-fonctionnalité`).
3. Committez vos modifications (`git commit -am 'Ajout de la nouvelle fonctionnalité'`).
4. Poussez votre branche (`git push origin feature-nouvelle-fonctionnalité`).
5. Créez une nouvelle pull request.

## License

Ce projet est sous la licence MIT - voir le fichier [LICENSE](LICENSE) pour plus de détails.

## Contact

Si vous avez des questions ou des suggestions, n’hésitez pas à me contacter à l’adresse suivante : [mohamedstaili962@gmail.com].
