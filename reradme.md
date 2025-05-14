# Annuaire CLI en Go

Ce projet est un petit annuaire en ligne de commande développé en **Golang**, permettant de gérer des contacts (nom + numéro de téléphone) depuis un terminal.

## Fonctionnalités

-  Ajouter un contact
-  Supprimer un contact
-  Modifier un contact
-  Rechercher un contact
-  Lister tous les contacts
-  Empêche les doublons (même nom)
-  Tests unitaires inclus
-  Persistance locale dans un fichier `contacts.json`

---

## Installation

### 1. Cloner le dépôt

```bash
git clone https://github.com/SaryDam/annuaire-go-tp-1
```
### 2. Initialiser Go

```bash
go mod tidy
```

## Utilisation

### Ajouter un contact

```bash
go run . --action ajouter --nom "Charlie" --tel "0811223344"
```

### Rechercher un contact

```bash
go run . --action rechercher --nom "Charlie"
```

### Supprimer un contact

```bash
go run . --action supprimer --nom "Charlie"
```

### Modifier un contact

```bash
go run . --action modifier --nom "Charlie" --tel "0999888777"
```

### Lister les contacts

```bash
go run . --action lister
```

### Exécution des tests

```bash
go test
```

Damien Creux