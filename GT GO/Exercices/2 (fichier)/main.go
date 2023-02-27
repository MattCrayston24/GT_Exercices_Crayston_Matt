package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	for {
		fmt.Println("Pour accéder au menu fichier.txt, tapez 1")
		fmt.Println("Pour quitter, tapez 2")

		var choiceaccessmenu int
		fmt.Scan(&choiceaccessmenu)

		switch choiceaccessmenu {
		case 1:
			Menu()
		case 2:
			Menuu()
		default:
			fmt.Println("Veuillez choisir un chiffre entre 1 et 2.")
		}
	}
}

func Menuu() {
	for {
		fmt.Println("Vous avez quitter...")
		fmt.Println("Pour revenir à la précédente page, tapez 1")

		var choiceaccecsorleave int
		fmt.Scan(&choiceaccecsorleave)

		switch choiceaccecsorleave {
		case 1:
			Menu()
		case 2:
			leave()
		default:
			fmt.Println("Veuillez choisir le chiffre 1 pour accéder à la page précédente.")
		}
	}
}

func leave() {
	fmt.Println("Vous avez quitter.")
}

func Menu() {
	for {
		fmt.Println("M E N U :")
		fmt.Println("Tapez 1 pour : Dans un fichier (.txt), récupérer tout son contenu.")
		fmt.Println("Tapez 2 pour : Dans un fichier (.txt), ajouter du contenu")
		fmt.Println("Tapez 3 pour : Dans un fichier (.txt), supprimer tout le contenu")
		fmt.Println("Tapez 4 pour : Dans un fichier (.txt), remplacer tout le contenu")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			Recuperation()
		case 2:
			Ajouter()
		case 3:
			Suppression()
		case 4:
			Replace()
		default:
			fmt.Println("Veuillez choisir un chiffre entre 1 et 4 (1 et 4 compris) pour accéder à une fonction du menu.")
		}
	}
}

func Recuperation() {
	fmt.Print("Quel est le nom du Fichier ?")
	var Fichier string
	fmt.Scanln(&Fichier)

	texte, err := ioutil.ReadFile(Fichier)
	if err != nil {
		fmt.Println("Un problème est survenu lors de la lecture du fichier :", err)
	} else {
		fmt.Println("Le fichier contient : ", string(texte))
	}
}

func Ajouter() {
	fmt.Print("Quel est le nom du Fichier ? :")
	var Fichier string
	fmt.Scanln(&Fichier)

	fichier, err := os.OpenFile(Fichier, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Un problème est survenu lors de l'ouverture du fichier :", err)
		return
	}
	defer fichier.Close()

	fmt.Print("Que faut-il ajouter ? (entrez le texte) :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texte := scanner.Text()

	_, err = fichier.WriteString(texte)
	if err != nil {
		fmt.Println("Un problème est survenu lors de l'écriture dans le fichier :", err)
	} else {
		fmt.Println("le texte a été ajouté sans problème !")
	}
}

func Suppression() {
	fmt.Print("Quel est le nom du Fichier ? :")
	var Fichier string
	fmt.Scanln(&Fichier)

	err := ioutil.WriteFile(Fichier, []byte(""), 0644)
	if err != nil {
		fmt.Println("Erreur lors de la suppression du contenu du fichier :", err)
	} else {
		fmt.Println("Contenu du fichier supprimé avec succès")
	}
}

func Replace() {
	fmt.Print("Quel est le nom du Fichier ? :")
	var Fichier string
	fmt.Scanln(&Fichier)

	fmt.Print("Quel est le nouveau contenu du fichier ? :")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	texte := scanner.Text()

	err := ioutil.WriteFile(Fichier, []byte(texte), 0644)
	if err != nil {
		fmt.Println("Un problème est survenu lors du remplacement, pour le contenu du fichier :", err)
	} else {
		fmt.Println("Le contenu du fichier a été remplacé sans aucun problème !!")
	}
}
