package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	action := flag.String("action", "", "Action à effectuer")
	nom := flag.String("nom", "", "Nom du contact")
	tel := flag.String("tel", "", "Numéro de téléphone")
	flag.Parse()

	annuaire, err := chargerAnnuaire()
	if err != nil {
		fmt.Println("Erreur lors du chargement de l'annuaire :", err)
		os.Exit(1)
	}

	switch *action {
	case "ajouter":
		if *nom == "" || *tel == "" {
			fmt.Println("Veuillez fournir un nom et un numéro de téléphone.")
			return
		}
		err := annuaire.Ajouter(Contact{Nom: *nom, Tel: *tel})
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Contact ajouté avec succès.")
		}
	case "rechercher":
		if *nom == "" {
			fmt.Println("Veuillez fournir un nom.")
			return
		}
		contact := annuaire.Rechercher(*nom)
		if contact != nil {
			fmt.Printf("Contact trouvé : %s - %s\n", contact.Nom, contact.Tel)
		} else {
			fmt.Println("Aucun contact trouvé.")
		}
	case "supprimer":
		if *nom == "" {
			fmt.Println("Veuillez fournir un nom.")
			return
		}
		err := annuaire.Supprimer(*nom)
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Contact supprimé.")
		}
	case "modifier":
		if *nom == "" || *tel == "" {
			fmt.Println("Veuillez fournir le nom du contact et le nouveau téléphone.")
			return
		}
		err := annuaire.Modifier(*nom, Contact{Nom: *nom, Tel: *tel})
		if err != nil {
			fmt.Println("Erreur :", err)
		} else {
			fmt.Println("Contact modifié.")
		}
	case "lister":
		fmt.Println("Liste des contacts :")
		for _, c := range annuaire.Contacts {
			fmt.Printf("- %s : %s\n", c.Nom, c.Tel)
		}
	default:
		fmt.Println("Action inconnue. Actions possibles : ajouter, rechercher, supprimer, modifier, lister")
	}
}
