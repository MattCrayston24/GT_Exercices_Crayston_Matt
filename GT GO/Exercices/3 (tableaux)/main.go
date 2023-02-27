package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(5) + 5
	tab := make([]interface{}, 1)

	for i := 0; i < n; i++ {
		tab = []interface{}{tab}
		taille := rand.Intn(5) + 1
		for j := 0; j < taille; j++ {
			tab = append(tab, rand.Intn(100))
		}
	}

	html := "<html><body>" + affichertab(tab) + "</body></html>"
	fmt.Println(html)
}

func affichertab(tab interface{}) string {
	html := "<table>"

	switch t := tab.(type) {
	case []interface{}:
		for _, v := range t {
			html += "<tr><td>" + affichertab(v) + "</td></tr>"
		}
	default:
		html += fmt.Sprint(t)
	}

	html += "</table>"
	return html
}
