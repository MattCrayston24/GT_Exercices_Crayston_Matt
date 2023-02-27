package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	// Initialisation du générateur de nombres aléatoires
	rand.Seed(time.Now().UnixNano())

	// Génération d'un tableau à n dimensions aléatoires
	n := rand.Intn(3) + 3
	tableau := genererTableau(n)

	// Affichage du tableau en HTML
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		afficherTableauHTML(w, tableau)
	})
	http.ListenAndServe(":8080", nil)
}

func genererTableau(n int) interface{} {
	if n == 1 {
		tableau := make([]int, rand.Intn(5)+5)
		for i := range tableau {
			tableau[i] = rand.Intn(100)
		}
		return tableau
	} else {
		tableau := make([]interface{}, rand.Intn(5)+5)
		for i := range tableau {
			tableau[i] = genererTableau(n - 1)
		}
		return tableau
	}
}

func afficherTableauHTML(w http.ResponseWriter, tableau interface{}) {
	fmt.Fprintln(w, "<html><body><table border='1'>")
	afficherTableauHTMLRec(w, tableau)
	fmt.Fprintln(w, "</table></body></html>")
}

func afficherTableauHTMLRec(w http.ResponseWriter, tableau interface{}) {
	switch t := tableau.(type) {
	case []int:
		fmt.Fprintln(w, "<tr>")
		for _, valeur := range t {
			fmt.Fprintf(w, "<td>%d</td>", valeur)
		}
		fmt.Fprintln(w, "</tr>")
	case []interface{}:
		for _, sousTableau := range t {
			afficherTableauHTMLRec(w, sousTableau)
		}
	}
}
