package main

import (
	"encoding/json"
	"errors"
	"os"
)

type Contact struct {
	Nom string `json:"nom"`
	Tel string `json:"tel"`
}

type Annuaire struct {
	Contacts []Contact `json:"contacts"`
}

const dataFile = "contacts.json"

func chargerAnnuaire() (Annuaire, error) {
	var a Annuaire
	file, err := os.ReadFile(dataFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return a, nil
		}
		return a, err
	}
	err = json.Unmarshal(file, &a)
	return a, err
}

func sauvegarderAnnuaire(a Annuaire) error {
	data, err := json.MarshalIndent(a, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(dataFile, data, 0644)
}

func (a *Annuaire) Ajouter(contact Contact) error {
	for _, c := range a.Contacts {
		if c.Nom == contact.Nom {
			return errors.New("un contact avec ce nom existe déjà")
		}
	}
	a.Contacts = append(a.Contacts, contact)
	return sauvegarderAnnuaire(*a)
}

func (a *Annuaire) Rechercher(nom string) *Contact {
	for _, c := range a.Contacts {
		if c.Nom == nom {
			return &c
		}
	}
	return nil
}

func (a *Annuaire) Supprimer(nom string) error {
	for i, c := range a.Contacts {
		if c.Nom == nom {
			a.Contacts = append(a.Contacts[:i], a.Contacts[i+1:]...)
			return sauvegarderAnnuaire(*a)
		}
	}
	return errors.New("contact introuvable")
}

func (a *Annuaire) Modifier(nom string, nouveau Contact) error {
	for i, c := range a.Contacts {
		if c.Nom == nom {
			a.Contacts[i] = nouveau
			return sauvegarderAnnuaire(*a)
		}
	}
	return errors.New("contact à modifier introuvable")
}
