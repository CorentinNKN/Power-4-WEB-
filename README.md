#  Puissance 4 - Version Go + HTML

Un jeu de Puissance 4 développé entièrement en Go avec une interface web esthétique.

##  Lancement rapide

1. **Ouvrez un terminal** dans ce dossier
2. **Lancez le serveur** :
   ```bash
   go run main.go
   ```
3. **Ouvrez votre navigateur** sur : `http://localhost:8080`

##  Fonctionnalités

-  Jeu de Puissance 4 classique (2 joueurs)
-  Interface web moderne avec animations
-  Backend Go performant
-  Détection automatique de victoire et match nul
-  Design responsive (mobile + desktop)
-  Pas de JavaScript requis côté client

##  Comment jouer

1. **Joueur 1** (🔴) commence
2. **Cliquez** sur les boutons ↓ pour faire tomber votre jeton
3. **Objectif** : Alignez 4 jetons (horizontal, vertical ou diagonal)
4. **Nouvelle partie** : Cliquez sur " Nouvelle Partie"

##  Technologie

- **Backend** : Go pur (serveur HTTP natif)
- **Frontend** : HTML5 + CSS3 intégré
- **Base de données** : Aucune (état en mémoire)

##  Structure

```
Power-4-WEB-/
├── main.go     # Serveur Go complet avec logique de jeu
├── go.mod      # Module Go
└── README.md   # Documentation
```

##  Personnalisation

Le CSS est intégré dans le fichier `main.go`. Vous pouvez modifier les couleurs et animations dans la section `<style>` du template HTML.

---

**Créé par une equipe de pro 🤡🤡 en Go**