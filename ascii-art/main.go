package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	Args1 := os.Args[1] // Premier argument donné dans la commande

	content := strings.Split(Args1, "\\n") // Eclatement de l'argument pour chaque occurrence des caractères \ et n

	var s []string
	str := ""
	// arg2 := os.Args[2]
	// fmt.Println(arg2)
	// Library
	file, err := os.Open("shadow.txt")
	if err != nil {
		fmt.Println("Le fichier n'existe pas.")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text())
	}

	for _, element := range content {
		if len(element) > 0 {
			line := []rune(element)

			for a := 0; a < 8; a++ {

				for i := 0; i < len(line); i++ {
					group := (int(line[i]) - 32) * 9
					adress := group + a + 1

					str += (s[adress])
				}
				str += "\n"
			}
		} else {
			str += "\n"
		}
	}
	fmt.Print(str)
}
