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

	Justify := flag.String("align", arg3, "")

	flag.Parse()

	/*fichier, _ := os.Create(*outpute)

	defer fichier.Close()
	*/
	// Lecture du tableau de lignes données dans la commande
	for _, element := range content {
		var adress int
		var group int

		// var str string
		// Dans le cas où la cellule lue n'est pas vide
		if len(element) > 0 {
			line := []rune(element) // Eclatement du contenu de la cellule en tableau de rune
			if len(line) > 14 {
				fmt.Println("choisissez un mot moins grand ou agrandissez le terminal!")
			}

			// Boucle pour les 8 lignes de l'ascii art
			for a := 0; a < 8; a++ {

				for i := 0; i < len(line); i++ { // commande du terminal

					group = (int(line[i]) - 32) * 9
					adress = group + a + 1

					if s[adress] != "" {
						if strings.Contains(*Justify, "center") {
							fmt.Printf("%50v", (s[adress]))
							break
						}
						if strings.Contains(*Justify, "left") {
							fmt.Printf("%1v", (s[adress]))
							break
						}
						if strings.Contains(*Justify, "right") {
							fmt.Printf("%70v", (s[adress]))
							break
						}
					}
				}

				for j := 1; j < len(line); j++ {
					group = (int(line[j]) - 32) * 9

					adress = group + a + 1
					if strings.Contains(*Justify, "center") {
						fmt.Printf(s[adress])
					}
					if strings.Contains(*Justify, "left") {
						fmt.Printf(s[adress])
					}
					if strings.Contains(*Justify, "right") {
						fmt.Printf(s[adress])
					}
				}
				if strings.Contains(*Justify, "center") {
					fmt.Printf("%50v", string(rune('\n')))
				}
				if strings.Contains(*Justify, "left") {
					fmt.Printf("%1v", string(rune('\n')))
				}
				if strings.Contains(*Justify, "right") {
					fmt.Printf("%70v", string(rune('\n')))
				}
			}

		} else {
			if strings.Contains(*Justify, "center") {
				fmt.Printf("%50v", string(rune('\n')))
			}
			if strings.Contains(*Justify, "left") {
				fmt.Printf("%1v", string(rune('\n')))
			}
			if strings.Contains(*Justify, "right") {
				fmt.Printf("%70v", string(rune('\n')))
			}

		}

	}
}
