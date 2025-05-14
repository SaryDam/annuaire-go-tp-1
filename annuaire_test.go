package main

import (
	"os"
	"testing"
)

func TestAjouterEtRechercher(t *testing.T) {
	os.Remove(dataFile)

	a := Annuaire{}
	contact := Contact{"Alice", "0123456789"}
	err := a.Ajouter(contact)
	if err != nil {
		t.Fatalf("Erreur à l'ajout : %v", err)
	}

	trouve := a.Rechercher("Alice")
	if trouve == nil || trouve.Tel != "0123456789" {
		t.Fatalf("Contact non trouvé ou incorrect")
	}
}

func TestDoublon(t *testing.T) {
	a := Annuaire{}
	a.Ajouter(Contact{"Bob", "0987654321"})
	err := a.Ajouter(Contact{"Bob", "1231231234"})
	if err == nil {
		t.Fatalf("Erreur attendue pour un doublon")
	}
}

func TestModifier(t *testing.T) {
	a := Annuaire{}
	a.Ajouter(Contact{"Eve", "0000"})
	err := a.Modifier("Eve", Contact{"Eve", "1111"})
	if err != nil {
		t.Fatalf("Erreur à la modification : %v", err)
	}
	contact := a.Rechercher("Eve")
	if contact.Tel != "1111" {
		t.Fatalf("Modification non effectuée")
	}
}

func TestSupprimer(t *testing.T) {
	a := Annuaire{}
	a.Ajouter(Contact{"Zack", "9999"})
	err := a.Supprimer("Zack")
	if err != nil {
		t.Fatalf("Erreur à la suppression : %v", err)
	}
	if a.Rechercher("Zack") != nil {
		t.Fatalf("Contact non supprimé")
	}
}
