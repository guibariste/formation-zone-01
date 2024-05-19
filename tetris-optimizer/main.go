package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func change_letter(s []string) string {
	var strconv string
	for i := 0; i < len(s); i++ {
		switch {
		case i == 0:
			strconv = strings.ReplaceAll(s[i], "#", "A")

		case i == 1:
			strconv += strings.ReplaceAll(s[i], "#", "B")
		case i == 2 && len(s) > 2:
			strconv += strings.ReplaceAll(s[i], "#", "C")
		case i == 3 && len(s) > 3:
			strconv += strings.ReplaceAll(s[i], "#", "D")
		case i == 4 && len(s) > 4:
			strconv += strings.ReplaceAll(s[i], "#", "E")
		case i == 5 && len(s) > 5:
			strconv += strings.ReplaceAll(s[i], "#", "F")
		case i == 6 && len(s) > 6:
			strconv += strings.ReplaceAll(s[i], "#", "G")
		case i == 7 && len(s) > 7:
			strconv += strings.ReplaceAll(s[i], "#", "H")
		case i == 8 && len(s) > 8:
			strconv += strings.ReplaceAll(s[i], "#", "I")
		case i == 9 && len(s) > 9:
			strconv += strings.ReplaceAll(s[i], "#", "J")
		case i == 10 && len(s) > 10:
			strconv += strings.ReplaceAll(s[i], "#", "K")
		case i == 11 && len(s) > 11:
			strconv += strings.ReplaceAll(s[i], "#", "L")
		case i == 12 && len(s) > 12:
			strconv += strings.ReplaceAll(s[i], "#", "M")
		case i == 13 && len(s) > 13:
			strconv += strings.ReplaceAll(s[i], "#", "N")
		case i == 14 && len(s) > 14:
			strconv += strings.ReplaceAll(s[i], "#", "0")
		case i == 15 && len(s) > 15:
			strconv += strings.ReplaceAll(s[i], "#", "P")
		}
	}
	return strconv
}

func letter_n(s []string, str string) int {
	alphabet := []string{"A", "B", "C"}
	var z int
	var n int
	for z = range s {
		for _, str = range alphabet {
			if strings.Contains(s[z], str) {
				n = z - 1
			}
		}
	}
	return n
}

func deplace(sliceunsortedC []int, pointC []int, n int) []int {
	var formeC []int

	// var i int
	for range sliceunsortedC {
		if sliceunsortedC[0] == 1 && sliceunsortedC[1] == 4 && sliceunsortedC[2] == 1 {
			formeC = append(formeC, pointC[0])
			formeC = append(formeC, (pointC[0] + sliceunsortedC[0]))

			formeC = append(formeC, (pointC[0] + sliceunsortedC[0] + sliceunsortedC[1] + n))

			formeC = append(formeC, (pointC[0]+sliceunsortedC[0]+sliceunsortedC[1]+sliceunsortedC[2])+n)
		} else {
			formeC = append(formeC, pointC[0])
			if sliceunsortedC[0] == 5 {
				formeC = append(formeC, (pointC[0]+sliceunsortedC[0])+n)
			}
			if sliceunsortedC[1] == 5 {
				formeC = append(formeC, (pointC[0] + sliceunsortedC[0] + sliceunsortedC[1] + 2*n))
			}
			if sliceunsortedC[2] == 1 {
				formeC = append(formeC, (pointC[0]+sliceunsortedC[0]+sliceunsortedC[1]+sliceunsortedC[2])+2*n)
			}
		}
		/*for i = range sliceunsortedC {

			if (sliceunsortedC[i] == 5 || sliceunsortedC[i] == 4) && i == 0 { // si on agrandit le carre les ecarts sont plus importants

				formeC = append(formeC, pointC[0])
				formeC = append(formeC, (pointC[0]+sliceunsortedC[i])+n) // rectifier si il y aun ecart de 1

			} else if (sliceunsortedC[i] == 1) && i == 0 { // si on agrandit le carre les ecarts sont plus importants

				formeC = append(formeC, pointC[0])
				formeC = append(formeC, (pointC[0] + sliceunsortedC[i])) // rectifier si il y aun ecart de 1

			}

			if (sliceunsortedC[i] == 5) && i == 1 {
				formeC = append(formeC, (pointC[0] + sliceunsortedC[i-1] + sliceunsortedC[i] + (2 * n)))
			} else if (sliceunsortedC[i] == 4) && i == 1 {
				formeC = append(formeC, (pointC[0] + sliceunsortedC[i-1] + sliceunsortedC[i] + n))
			} else if (sliceunsortedC[i] == 1) && i == 1 {
				formeC = append(formeC, (pointC[0]+sliceunsortedC[i-1]+sliceunsortedC[i])+n)
			}
			if sliceunsortedC[i] == 5 && i == 2 {
				formeC = append(formeC, (pointC[0]+sliceunsortedC[i-2]+sliceunsortedC[i-1]+sliceunsortedC[i])+2*n)
			} else if sliceunsortedC[i] == 4 && i == 2 {
				formeC = append(formeC, (pointC[0]+sliceunsortedC[i-2]+sliceunsortedC[i-1]+sliceunsortedC[i])+2*n)
			} else if (sliceunsortedC[i] == 1) && i == 2 {
				formeC = append(formeC, (pointC[0]+sliceunsortedC[i-2]+sliceunsortedC[i-1]+sliceunsortedC[i])+2*n)
			}
		}*/
	}
	return formeC
}

func unsorted(ap []string, str string) []int {
	var sortedA []int
	var i int
	var w string

	for i, w = range ap {
		if w == str {
			sortedA = append(sortedA, i)
		}
	}
	return sortedA
}

func point(ap []string, s string) []int {
	var pointB []int
	var i int
	var w string
	for i, w = range ap {
		if w == s {
			pointB = append(pointB, i) // regarder les correlations entre les differentes slices et ce qui matche
		}
	}
	return pointB
}

func ecart(unsortedB []int) []int {
	var ecartunsortedB int
	var sliceunsortedB []int
	var i int
	for i = 0; i < len(unsortedB); i++ {
		if i+1 < len(unsortedB) {
			ecartunsortedB = unsortedB[i+1] - unsortedB[i]
			sliceunsortedB = append(sliceunsortedB, ecartunsortedB)
		}
	}
	return sliceunsortedB
}

func depart(new []string, formeC []int, pointC []int, sliceunsortedC []int, str string, n int) {
	var formeC0 int
	var formeC1 int
	var formeC2 int
	var formeC3 int
	var nb int
	var i int
	var index []int
	// var index1 []int
	for nb = range pointC { // refaire une fonction depart avec ca
		if formeC[0] == pointC[0] {
			formeC0 = pointC[0]
			// fmt.Print(formeC0, "formeC0")

			formeC1 = (pointC[0] + sliceunsortedC[0]) + n

			// fmt.Print(formeC1, "formeC1")
			formeC2 = (pointC[0] + sliceunsortedC[0] + sliceunsortedC[1]) + 2*n

			// fmt.Print(formeC2, "formeC2")
			formeC3 = (pointC[0] + sliceunsortedC[0] + sliceunsortedC[1] + sliceunsortedC[2]) + 3*n

			// fmt.Print(formeC3, "formeC3")
			// fmt.Print("\n")
		}
	}
	for nb = range pointC {
		if pointC[nb] == formeC0 /*|| pointC[n] == formeC1 || pointC[n] == formeC2 || pointC[n] == formeC3 */ {
			index = append(index, nb)
		}
		if pointC[nb] == formeC1 {
			index = append(index, nb)
			fmt.Print(formeC1, nb, "ici")
			fmt.Print(pointC, "pointC")
			fmt.Print(index, "index")
		} /*else {
			if pointC[n] == formeC1+1 || pointC[n] == formeC1+2 {
				index = append(index, n)
			}
		}*/
		if pointC[nb] == formeC2 {
			index = append(index, nb)
		}

		if pointC[nb] == formeC3 {
			index = append(index, nb)
		}

	}
	if len(index) == 4 {
		if formeC0 == pointC[index[0]] && formeC1 == pointC[index[1]] && formeC2 == pointC[index[2]] && formeC3 == pointC[index[3]] {
			for i = range index {
				new[pointC[index[i]]] = str // a rectifier
			}
		}
	} /*else {
		index1 = []int{0}
		index1[0] = index[0] + 1

		for nb = range pointC {
			if pointC[index1[0]] == pointC[nb] {
				continue
			}
			if pointC[index1[0]]+(sliceunsortedC[0])+n == pointC[nb] {
				index1 = append(index1, nb)
			}
			if pointC[index1[0]]+(sliceunsortedC[0]+sliceunsortedC[1])+2*n == pointC[nb] {
				index1 = append(index1, nb)
			}
			if pointC[index1[0]]+(sliceunsortedC[0]+sliceunsortedC[1]+sliceunsortedC[2])+3*n == pointC[nb] {
				index1 = append(index1, nb)
			}
		}
	}
	for i = range index1 {
		new[pointC[index1[i]]] = str
	}*/
}

func remove(unsortedB []int, ap []string) {
	var i int

	for i = range unsortedB {
		ap[unsortedB[i]] = "."
	}
}

func expand(forme []int, ap []string) {
	// var j int
	var w string

	var h int
	// var ajout []int
	// var ajoutstr []string
	var formestr []string
	var ligne1 string
	var ligne []string
	file2, _ := os.ReadFile("ligne.txt")
	ligne1 = string(file2)
	ligne = strings.Split(ligne1, "\n")
	// fmt.Print(ligne, "lignevrai")
	for h = range forme {
		formestr = append(formestr, strconv.Itoa(forme[h]))
	}
	// fmt.Print(formestr, "normalement4")
	for _, w = range ligne {
		if w == formestr[0] || w == formestr[1] || w == formestr[2] || w == formestr[3] {
			cible, _ := strconv.Atoi(w)

			ap[cible] += "\n"
		}
	}
}

func ligne(ap []string) string {
	var i int
	var w string
	var ligne []string

	for i, w = range ap {
		if w == "\n" {
			ligne = append(ligne, strconv.Itoa(i))
		}
	}
	return strings.Join(ligne, "\n")
}

func main() {
	var str string

	var i int

	var sortedA []int
	var unsortedB []int
	var unsortedC []int
	var pointB []int
	var ap []string
	var sliceunsortedB []int
	var lignedepart string
	var formeA []int
	var sliceunsortedA []int

	/*lgauche := []int{5, 1, 1}
	ldroite := []int{3, 1, 1}
		Lgauche := []int{5,5 ,1}
		Ldroite := []int{5,4 ,1}
	bitedroite := []int{4, 1, 5}
		bitegauche := []int{5,1,4}
	Linversdroit := []int{1, 4, 5}*/

	var pointC []int
	var pointA []int
	var sliceunsortedC []int
	var formeC []int
	var formeB []int
	var n int

	news := "....\n....\n....\n....\n"

	file, err := os.OpenFile("ligne.txt", os.O_CREATE|os.O_WRONLY, 1024)
	if err != nil {
		panic(err)
	}

	file1, _ := os.ReadFile("tetris.txt")

	str = string(file1)
	strsplit := strings.Split(str, "\n")
	for i = range strsplit {
		if strsplit[i] == "" {
			strsplit[i] = "/"
		}
		if len(strsplit) > 20 || strsplit[i] == "####" {
			news = "......\n......\n......\n......\n"
			n = 2
		}
	}
	new := strings.Split(news, "")
	str1 := strings.Join(strsplit, "\n")

	strsplit1 := strings.Split(str1, "/")
	if len(strsplit1) > 5 {
		strsplit1[0] += "......\n......\n"
	}

	strconve := change_letter(strsplit1)

	ap = strings.Split(strconve, "")

	lignedepart = ligne(ap)
	lignefile := []byte(lignedepart)
	_, err = file.Write(lignefile) // écrire dans le fichier
	if err != nil {
		panic(err)
	}
	sortedA = unsorted(ap, "A")
	fmt.Print(sortedA, "aaaaaaaaaaaaa")
	sliceunsortedA = ecart(sortedA)
	fmt.Print(sliceunsortedA, "sliceA")
	pointA = point(new, ".")
	fmt.Print(pointA, "pointA")
	formeA = deplace(sliceunsortedA, pointA, n)
	// depart(new, formeA, pointA, sliceunsortedA, "A", n)
	for i = range formeA {
		new[formeA[i]] = "A"
	}
	fmt.Print(formeA, "aaaaaaaaaaaa")
	//---------------------------------------------------------------------------------
	pointB = point(new, ".")

	// fmt.Print(pointB, "pointb")

	unsortedB = unsorted(ap, "B")

	unsortedC = unsorted(ap, "C")

	sliceunsortedB = ecart(unsortedB)
	// ap[9] = "."

	formeB = deplace(sliceunsortedB, pointB, n)
	fmt.Print(formeB, "b")
	/*for i = range formeB {
		formeB[i] += 5
	}*/ // cas particulier a mettre(b tout en bas ou bien mettre c en premier)

	// depart(new, formeB, pointB, sliceunsortedB, "B", n)
	// expand(formeB, ap)
	// remove(unsortedB, ap)

	//------------------------------------------------------------------------------
	pointC = point(new, ".")
	fmt.Print(pointC, "pointC")
	fmt.Print(ecart(unsortedC), "ecartC")
	sliceunsortedC = ecart(unsortedC)
	// fmt.Print(sliceunsortedC, "sliceunsortedC")
	// fmt.Print("\n")
	// ap[4] = "."
	formeC = deplace(sliceunsortedC, pointC, n)
	fmt.Print(formeC, "c")
	// depart(new, formeC, pointC, sliceunsortedC, "C", n)
	fmt.Print(n, "n")
	// remove(unsortedC, ap)

	// depart(ap, formeC, pointC, "C")
	// expand(formeC, ap)

	// fmt.Println(ap[21], "ligne vide")

	// pointD := point(ap, ".") // pour rajouter des points fonction a faire(fo rajouiter aussi des indexs a ap[])

	// fmt.Println(pointD, "pointD")

	//	fmt.Print(formeC, "c apres")
	//	fmt.Print("\n")
	//---------------------------------------------------------------------------
	newer := strings.Join(new, "")
	fmt.Print("\n")
	fmt.Print(newer)

	// fmt.Print(ap)

	/* fmt.Println(sorted)
	 fmt.Print(slicepointBB)

	fmt.Println(unsortedC, "unsortedC")

	fmt.Print(sliceunsortedB, "sliceunsortedB")
	fmt.Print(formeB, "formb")
	fmt.Print(pointB, "pointB")
	fmt.Print(sortedA, "aapres")

	 fmt.Print(formeC, "formeC")
	 fmt.Print(formeB, "formeB")*/
}
