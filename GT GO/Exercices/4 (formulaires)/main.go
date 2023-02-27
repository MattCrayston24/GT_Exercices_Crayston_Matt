package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			name := r.FormValue("name")
			message := r.FormValue("message")
			fmt.Printf("name: %s\nMessage: %s\n", name, message)
		}
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("KLe serveur est démarré, tapez (dans l'URL) : http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
