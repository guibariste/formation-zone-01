package main

import (
	"os"
	"strconv"
	"strings"
	"unicode"
)

func conv_cap(s string) string {
	return strings.Title(s)
}

func conv_up(s string) string {
	return strings.ToUpper(s)
}

func conv_low(s string) string {
	return strings.ToLower(s)
}

/*

cont, err := ioutil.ReadFile(fileName)    // peut etre utiliser au lieu du nom du fichier
if err != nil {
	fmt.Println("Error")
}*/
func main() {
	//	arg := os.Args[1:]
	//	err := os.Truncate(os.Args[2], 0)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	// fileName := arg[1]
	content, _ := os.ReadFile("sample.txt") // lire le fichier
	// if err != nil {
	//	fmt.Println(err)
	// return
	//	}

	cont := string(content)
	// cont = strings.ReplaceAll(cont, "‘", "/")

	str := strings.Split(cont, " ")

	/*	for i, words := range split {
		if strings.Contains(words, "(cap") {
		    split[i-1] = strings.Title(split[i-1])
		    findnvalue := regexp.MustCompile("[0-9*?]")
		    n := findnvalue.FindAllString(split[i+1], -1)
	*/

	for i, w := range str {

		if w == "(cap" {
			str[i-1] = conv_cap(str[i-1])
			// findnvalue := regexp.MustCompile("[0-9*?]")
			// n := findnvalue.FindAllString(str[i+1], -1)
		}
		if w == "(up" {
			str[i-1] = conv_up(str[i-1])
		}
		if w == "(low" {
			str[i-1] = conv_low(str[i-1])
		}
		if w == "(hex" {
			f, _ := strconv.ParseInt(str[i-1], 16, 64)

			str[i-1] = strconv.Itoa(int(f))
		}
		if w == "(bin" {
			f, _ := strconv.ParseInt(str[i-1], 2, 64)

			str[i-1] = strconv.Itoa(int(f))
			// fmt.Print(str)
		}

		if (i + 1) < len(str) {
			str[i+1] = strings.ReplaceAll(str[i+1], ")", "")
		}
		//if (i - 1) > len(str) {
		//str[i-1] = strings.ReplaceAll(str[i-1], "(", "")
		//}

		if w == "(cap," {

			ncap, _ := strconv.Atoi(str[i+1])

			{
				for z := (i - 1); z != (i-1)-ncap; z-- {
					str[z] = strings.Title(str[z]) // (cap,n)
					str[i+1] = ""
				}
			}
		}
		if w == "(low," {
			nlow, _ := strconv.Atoi(str[i+1])

			for z := (i - 1); z != (i-1)-nlow; z-- { // a revoir tout

				str[z] = strings.ToLower(str[z])
				str[i+1] = ""
			}

		}
		if w == "(up," {
			nup, _ := strconv.Atoi(str[i+1])
			for z := (i - 1); z != (i-1)-nup; z-- { // a revoir tout

				str[z] = strings.ToUpper(str[z])
				str[i+1] = ""
			}
		}
		if (i + 1) != len(str) {
			if w == "a" && (str[i+1][0] == 'u' || str[i+1][0] == 'e' || str[i+1][0] == 'i' || str[i+1][0] == 'o' || str[i+1][0] == 'a' || str[i+1][0] == 'y' || str[i+1][0] == 'h') {
				str[i] = "an"
			}
		}

		/*compteur := 0
		for i := 0; i != len(str); i++ {
			if str[i] == "’" {
				compteur++
				if i+1 < len(str) {
					if compteur%2 != 0 {
						str[i+1] = "%"
					}
				}                                        //marche pas

				//  pour l'autre apostrophe
				if i+1 < len(str) {
					if compteur%2 == 0 {
						str[i-1] = "%"
					}
				}
			}
		}*/
	}
	str1 := strings.Join(str, " ")

	str3 := []byte(str1)

	for i, z := range str3 {
		/*count := 0
		for i := 0; i != len(str3); i++ {
			if (str3[i]) == '/' {
				count++

				if i+1 < len(str3) {
					if count%2 != 0 {
						str3[i+1] = '%'
					}
					if i-2 != len(str3) {
						if count%2 == 0 {
							// str3[i-3] = str3[i-4]
							str3[i-1] = '%'
							str3[i-2] = '%'
							str3[i-3] = '.'

						}
					}
				}
			}
		}*/
		if z == ',' || z == '.' || z == ':' || z == ';' || z == '!' || z == '?' { // ponctuation on passe caractere par caractere(byte)
			if i+1 >= len(str3) {
				break
			}
			if unicode.IsLetter(rune(str3[i-1])) || unicode.IsNumber(rune(str3[i-1])) {
				continue
			}

			str3[i-1], str3[i] = str3[i], str3[i-1]
		}
	}
	str4 := string(str3)
	// str4 = strings.ReplaceAll(str4, "/", "‘")
	// str4 = strings.ReplaceAll(str4, "%", "")
	str4 = strings.ReplaceAll(str4, "  ", " ")
	str4 = strings.ReplaceAll(str4, "cap", "")
	str4 = strings.ReplaceAll(str4, "( ", "")

	str4 = strings.ReplaceAll(str4, " (, ", "")
	str4 = strings.ReplaceAll(str4, " (up, ", "")
	str4 = strings.ReplaceAll(str4, "(low, ", "")
	str4 = strings.ReplaceAll(str4, " (hex", "")
	str4 = strings.ReplaceAll(str4, " (bin", "")
	str4 = strings.ReplaceAll(str4, "‘ ", " ‘")
	str4 = strings.ReplaceAll(str4, " ‘", "‘")
	// fmt.Print(str4)
	str5 := []byte(str4)

	// fmt.Print(str4)

	file, err := os.OpenFile("result.txt", os.O_CREATE|os.O_WRONLY, 1024)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(str5) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
}
