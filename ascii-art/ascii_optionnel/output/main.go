package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	Args1 := os.Args[1]                    // Premier argument donné dans la commande
	content := strings.Split(Args1, "\\n") // Eclatement de l'argument pour chaque occurrence des caractères \ et n
	var s []string
	arg3 := os.Args[3]

	arg2 := os.Args[2]
	// fmt.Println(arg2)
	// Librar

	if arg2 == "shadow" {
		arg2 = "shadow.txt"
	} else if arg2 == "standard" {
		arg2 = "standard.txt"
	} else if arg2 == "thinkertoy" {
		arg2 = "thinkertoy.txt"
	}

	file, err := os.Open(arg2)
	if err != nil {
		fmt.Println("Le fichier n'existe pas.")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text()) // scanner ligne par ligne
		// fmt.Print(s)
	}

	outpute := flag.String("output", arg3, "")

	*outpute = strings.ReplaceAll(arg3, "--output=", "")
	*outpute = strings.ReplaceAll(arg3, "-output=", "")

	flag.Parse()

	fichier, _ := os.Create(*outpute)

	defer fichier.Close()
	// Lecture du tableau de lignes données dans la commande
	for _, element := range content {
		var adress int
		var group int

		// Dans le cas où la cellule lue n'est pas vide
		if len(element) > 0 {
			line := []rune(element) // Eclatement du contenu de la cellule en tableau de rune
			// Boucle pour les 8 lignes de l'ascii art
			for a := 0; a < 8; a++ {
				// LEcture du tableau de rune

				for i := 0; i < len(line); i++ { // commande du terminal

					group = (int(line[i]) - 32) * 9 // Définition de la première ligne dédiée à l'ascii art correspondant à la rune
					adress = group + a + 1          // Définition de l'adresse de la ligne de l'ascii art correspondant à la rune, à la ligne imprimé + décalage
					// De 1 pour prendre en compte la ligne de séparation entre les groupes d'ascii art

					fichier.WriteString((s[adress])) // Impression de la ligne récupétrée

				}
				fichier.WriteString(string(rune('\n'))) // Impression d'un retour à la ligne pour passer aux lignes suivantes
			}
		} else {
			fichier.WriteString(string(rune('\n'))) // Impression d'une ligne seule si la cellule ne contient pas de texte
		}

	}
}
