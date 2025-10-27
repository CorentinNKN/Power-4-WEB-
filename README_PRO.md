# Puissance 4 Web 

Une implémentation moderne et sobre du jeu Puissance 4 avec une architecture séparée backend/frontend.

## Architecture

```
Power-4-WEB-/
├── backend/
│   └── server.go           # Serveur Go avec API REST et logique de jeu
├── frontend/
│   ├── index.html          # Page d'accueil
│   └── game.html           # Interface de jeu
├── static/
│   ├── css/
│   │   └── style.css       # Styles modernes et sobres
│   └── js/                 # Répertoire pour futurs scripts
├── go.mod                  # Module Go
└── README.md               # Documentation
```

##  Installation et lancement

### Prérequis
- Go 1.21 ou version supérieure

### Lancement
```bash
# Se placer dans le répertoire du projet
cd Power-4-WEB-

# Lancer le serveur depuis le dossier racine
go run backend/server.go
```

Le serveur sera accessible sur `http://localhost:8080`

##  Fonctionnalités

### Interface utilisateur
- **Page d'accueil** : Présentation du jeu et règles
- **Interface de jeu** : Design moderne et épuré
- **Responsive** : Compatible mobile et desktop
- **Accessibilité** : Navigation claire et intuitive

### Logique de jeu
- Plateau 6x7 classique
- Détection automatique de victoire (horizontal, vertical, diagonal)
- Gestion des tours de jeu
- Détection de match nul
- Reset de partie

### API REST (Bonus)
- `GET /api/game` : État actuel du jeu
- `POST /api/move` : Jouer un coup
- `POST /api/reset` : Nouvelle partie

##  Routes disponibles

| Route | Méthode | Description |
|-------|---------|-------------|
| `/` | GET | Page d'accueil |
| `/game` | GET | Interface de jeu |
| `/move` | POST | Jouer un coup |
| `/reset` | POST | Nouvelle partie |
| `/api/game` | GET | État du jeu (JSON) |
| `/api/move` | POST | Jouer un coup (JSON) |
| `/api/reset` | POST | Nouvelle partie (JSON) |

##  Design

### Palette de couleurs
- **Primaire** : `#2c3e50` (bleu marine)
- **Secondaire** : `#34495e` (gris bleu)
- **Accent** : `#3498db` (bleu)
- **Joueur 1** : `#e74c3c` (rouge)
- **Joueur 2** : `#f39c12` (orange)

### Caractéristiques
- Design sobre et professionnel
- Animations subtiles
- Interface claire et lisible
- Contrastes respectant les standards d'accessibilité

##  Développement

### Structure du code Go
- **Séparation des responsabilités** : Logique métier séparée du serveur
- **API REST** : Routes JSON pour intégration future
- **Templates HTML** : Rendu côté serveur
- **Gestion d'erreurs** : Responses HTTP appropriées

### CSS
- **Variables CSS** : Thème cohérent et maintenable
- **Grid Layout** : Mise en page moderne
- **Mobile First** : Responsive design
- **Animations** : Transitions fluides

##  Utilisation de l'API

### Exemple avec curl
```bash
# Obtenir l'état du jeu
curl http://localhost:8080/api/game

# Jouer un coup (colonne 3)
curl -X POST -d "column=3" http://localhost:8080/api/move

# Nouvelle partie
curl -X POST http://localhost:8080/api/reset
```

##  Statuts HTTP

- `200` : Succès
- `400` : Requête invalide (coup impossible)
- `405` : Méthode non autorisée
- `500` : Erreur serveur

##  Améliorations possibles

- Ajout d'un système de score
- Sauvegarde des parties
- Mode multijoueur en temps réel
- Thèmes personnalisables
- Historique des coups

---

**Développé avec Go et des technologies web modernes**