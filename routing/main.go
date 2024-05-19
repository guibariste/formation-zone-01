package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Serveur de fichiers statiques
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", fs)

	// Démarrer le serveur sur le port 8080
	fmt.Println("http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
