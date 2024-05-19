package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	Black  = "\033[1;30m"
	Red    = "\033[1;31m"
	Green  = "\033[1;32m"
	Yellow = "\033[1;33m"
	Purple = "\033[1;34m"
	Orange = "\033[0;33m"
	Blue   = "\033[1;36m"
	White  = "\033[1;37m"
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func main() {
	var str string
	var adress int
	var group int
	var line []rune
	var colored bool
	// var i int

	// var adress1 int
	// var group1 int
	var a int
	Args1 := os.Args[1]
	content := strings.Split(Args1, "\\n")
	var s []string
	// arg2 := os.Args[3]
	// arg4 := os.Args[4]
	arg2 := os.Args[2]
	/*kolor := flag.String("color", arg2, "")


	flag.Parse()*/

	/*if arg2 == "shadow" {
		arg2 = "shadow.txt"
	} else if arg2 == "standard" {
		arg2 = "standard.txt"
	} else if arg2 == "thinkertoy" {
		arg2 = "thinkertoy.txt"
	}
	*/
	file, err := os.Open("standard.txt")
	if err != nil {
		fmt.Println("Le fichier n'existe pas.")
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s = append(s, scanner.Text()) // scanner ligne par ligne
		// fmt.Print(s)
	}

	for _, element := range content {
		if !strings.Contains(arg2, "--color=") {
			fmt.Print("Usage: go run . [STRING] [OPTION]\n\nEX: go run . something --color=<color>\n")

			break
		}
		if len(element) > 0 {
			line = []rune(element)
			colored = false

			for a = 0; a < 8; a++ {

				for i := 0; i < len(line); i++ {

					if line[i] == '¤' && strings.Contains(arg2, "black") {
						colored = !colored
						if colored {
							str += Black // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] == '¤' && strings.Contains(arg2, "red") {
						colored = !colored
						if colored {
							str += Red // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] == '¤' && strings.Contains(arg2, "green") {
						colored = !colored
						if colored {
							str += Green // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] == '¤' && strings.Contains(arg2, "yellow") {
						colored = !colored
						if colored {
							str += Yellow // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] == '¤' && strings.Contains(arg2, "purple") {
						colored = !colored
						if colored {
							str += Purple // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] == '¤' && strings.Contains(arg2, "orange") {
						colored = !colored
						if colored {
							str += Orange // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] == '¤' && strings.Contains(arg2, "blue") {
						colored = !colored
						if colored {
							str += Blue // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] == '¤' && strings.Contains(arg2, "white") {
						colored = !colored
						if colored {
							str += White // lettre qu'on colore
						} else {
							str += "" // le reste
						}
						continue
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "black") {
						str += Black
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "red") {
						str += Red
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "green") {
						str += Green
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "yellow") {
						str += Yellow
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "purple") {
						str += Purple
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "orange") {
						str += Orange
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "blue") {
						str += Blue
					}
					if line[i] != '¤' && !colored && strings.Contains(arg2, "white") {
						str += White
					}

					group = (int(line[i]) - 32) * 9
					adress = group + a + 1

					str += (s[adress])
				}
				str += (string(rune('\n')))
			}
		} else {
			str += (string(rune('\n')))
		}

		fmt.Print(str)

	}
}
