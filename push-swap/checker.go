package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var count int

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

func initA(s string) ([]int, []int) {
	var spl []string
	var i int

	var a []int
	var b []int
	// b = []int{0}
	spl = strings.Split(s, " ")
	for i = range spl {

		in, err := strconv.Atoi(spl[i])
		if err != nil {
			fmt.Println("error(seulement des chiffres) !")
			os.Exit(1)

		}

		// fmt.Print(in)
		a = append(a, in)

	}
	// fmt.Print(len(a))
	if sort.IntsAreSorted(a) && len(a) > 1 {
		fmt.Println(" C'est deja trie")
		os.Exit(1)
	}
	return a, b
}

func initOnlyA(s string) []int {
	var spl []string
	var i int
	var in int
	var a []int

	// b = []int{0}
	spl = strings.Split(s, " ")
	for i = range spl {

		in, _ = strconv.Atoi(spl[i])

		// fmt.Print(in)
		a = append(a, in)

	}
	return a
}

func pb(a *[]int, b *[]int) {
	var bb []int
	var i int
	bb = append(bb, (*a)[0])
	*a = RemoveIndex(*a, 0)
	for i = range *b {
		bb = append(bb, (*b)[i])
	}
	*b = bb
}

func pa(a *[]int, b *[]int) {
	var aa []int
	var i int
	aa = append(aa, (*b)[0])
	*b = RemoveIndex(*b, 0)
	for i = range *a {
		aa = append(aa, (*a)[i])
	}
	*a = aa
}

func sa(a *[]int, b *[]int) {
	(*a)[0], (*a)[1] = (*a)[1], (*a)[0]
}

func sb(a *[]int, b *[]int) {
	(*b)[0], (*b)[1] = (*b)[1], (*b)[0]
}

func ss(a *[]int, b *[]int) {
	sa(a, b)
	sb(a, b)
}

func ra(a *[]int, b *[]int) {
	var aa []int
	var aaa []int
	aa = append(aa, (*a)[0])
	aaa = RemoveIndex(*a, 0)

	aaa = append(aaa, aa[0])
	*a = aaa
}

func rb(a *[]int, b *[]int) {
	var bb []int
	var bbb []int

	bb = append(bb, (*b)[0])
	bbb = RemoveIndex(*b, 0)
	bbb = append(bbb, bb[0])
	*b = bbb
}

func rr(a *[]int, b *[]int) {
	ra(a, b)
	rb(a, b)
}

func rra(a *[]int, b *[]int) {
	var aa []int
	var aaa []int
	aa = append(aa, (*a)[len(*a)-1])
	aaa = RemoveIndex(*a, len(*a)-1)
	for i := 0; i < len(aaa); i++ {
		aa = append(aa, aaa[i])
	}
	*a = aa
}

func rrb(a *[]int, b *[]int) {
	var bb []int
	var bbb []int
	bb = append(bb, (*b)[len(*b)-1])
	bbb = RemoveIndex(*b, len(*b)-1)
	for i := 0; i < len(bbb); i++ {
		bb = append(bb, bbb[i])
	}
	*b = bb
}

func rrr(a *[]int, b *[]int) {
	rra(a, b)
	rrb(a, b)
}

// func Middle(a []int) int {
// 	var middle int
// 	sort.Ints(a)
// 	index := len(a) / 2
// 	middle = a[index]
// 	return middle
// }

func CopyInts(arr []int) []int {
	cop := make([]int, len(arr))
	copy(cop, arr)
	sort.Ints(cop)
	return cop
}

func preTri(a *[]int, A *[]int, B *[]int) {
	var i int

	// var AA []int
	sort.Ints(*a)
	IndexRef := (*a)[len(*a)/2]

	Milieu := len(*A) / 2
	// fmt.Print(A[Milieu])

	for i = range *A {
		if (*A)[i] < IndexRef {
			// fmt.Println(IndexRef, "index ref")
			if i < Milieu {
				Moitie1 := (*A)[i]

				if Moitie1 == (*A)[1] {
					sa(A, B)
					count++
					fmt.Println("sa")

				}
				// if min == Moitie1 {
				for Moitie1 != (*A)[0] {
					ra(A, B)
					count++
					fmt.Println("ra")

				}
				if (*A)[0] == Moitie1 && Moitie1 < IndexRef { // index ref doit etre fixe
					pb(A, B)
					count++
					fmt.Println("pb")
					break
				}

				// }
				// pour 7296031 il doit y avoir dans B les nombres inferieurs a 3(0,1,2)
			}
			if i > Milieu || i == Milieu {
				Moitie2 := (*A)[i]

				for Moitie2 != (*A)[0] && Moitie2 < IndexRef {
					rra(A, B)
					count++
					fmt.Println("rra")
				}
				if (*A)[0] == Moitie2 && Moitie2 < IndexRef {
					pb(A, B)
					fmt.Println("pb")
					count++
					break

				}
			}

		}
	}
}

func preTri2(a *[]int, A *[]int, B *[]int) {
	var i int

	var AA []int
	sort.Ints(*a)
	IndexRef := (*a)[len(*a)/2]
	AA = CopyInts(*A)
	sort.Ints(AA)

	// min := (AA)[0]
	// fmt.Print(min, "min")

	// fmt.Println(AA, "AA")

	Milieu := len(*A) / 2
	// fmt.Print(A[Milieu])

	for i = range *A {
		if (*A)[i] < IndexRef {
			// fmt.Println(IndexRef, "index ref")
			if i < Milieu {
				Moitie1 := (*A)[i]

				if Moitie1 == (*A)[1] {
					sa(A, B)
					count++
					fmt.Println("sa")

				}
				// if min == Moitie1 {
				for Moitie1 != (*A)[0] {
					ra(A, B)
					count++
					fmt.Println("ra")

				}
				if (*A)[0] == Moitie1 && Moitie1 < IndexRef { // index ref doit etre fixe
					pb(A, B)
					count++
					fmt.Println("pb")
					break
				}

				// }
				// pour 7296031 il doit y avoir dans B les nombres inferieurs a 3(0,1,2)
			}
			if i > Milieu || i == Milieu {
				Moitie2 := (*A)[i]

				for Moitie2 != (*A)[0] && Moitie2 < IndexRef {
					rra(A, B)
					count++
					fmt.Println("rra")
				}
				if (*A)[0] == Moitie2 && Moitie2 < IndexRef {
					pb(A, B)
					fmt.Println("pb")
					count++
					break

				}
			}

		}
	}
}

func rendreB(A *[]int, B *[]int) {
	var Milieu int
	if len(*B) == 2 {
		Milieu = (*B)[0]
	} else {
		Milieu = len(*B) / 2
	}

	BB := CopyInts(*B)
	// sort.Ints(BB)
	var max int
	if max < len(BB) {
		// fmt.Print(len(BB), "BB")
		max = (BB)[len(BB)-1]
	}
	for i := range *B {
		// fmt.Print(BB, "BB")
		// fmt.Print(max, "max")
		if i < len(*B) {

			if i <= Milieu {
				for (*B)[0] < (*B)[i] {
					rb(A, B)
					count++
					fmt.Println("rb")
					// fmt.Print(A, B, "rbbbbb")
				}
			}
			if i > Milieu {
				for (*B)[0] < (*B)[i] {
					rrb(A, B)
					count++
					// fmt.Print(A, B)
					// fmt.Print(Milieu, "mil")
					fmt.Println("rrb")

				}
			}

			// if len(*B) == 2 && (*B)[0] < (*B)[1] {
			// 	sb(A, B)
			// 	fmt.Print("sb")
			// }
			if len(*B) == 1 {
				pa(A, B)
				fmt.Println("pa")
				count++
			} else if (*B)[0] >= (*B)[1] {

				pa(A, B)
				// fmt.Print(A, B)
				count++
				fmt.Println("pa") // faire une moitie aussi pour dire soit rb soit rrb
			}
			if len(*B) == 2 && (*B)[0] == (*B)[1] {
				pa(A, B)
				fmt.Println("pa")
				count++
			}

		}
	}

	// fmt.Print(Milieu, "milieu")
	// fmt.Print(A, B)
}

func main() {
	if len(os.Args) > 1 {

		var recu []string

		Arg1 := os.Args[1]
		A, B := initA(Arg1)

		console := bufio.NewScanner(os.Stdin)
		for console.Scan() {
			if console.Text() != "" {
				recu = append(recu, console.Text())
			} else {
				break
			}
		}
		// fmt.Print(recu)
		for i := range recu {
			if recu[i] == "pa" {
				pa(&A, &B)
				// fmt.Print("pa")
			}
			if recu[i] == "pb" {
				pb(&A, &B)
				// fmt.Print("pb")
			}
			if recu[i] == "sa" {
				sa(&A, &B)
				// fmt.Print("sa")
			}
			if recu[i] == "sb" {
				sb(&A, &B)
				// fmt.Print("sb")
			}
			if recu[i] == "ss" {
				ss(&A, &B)
				// fmt.Print("ss")
			}
			if recu[i] == "ra" {
				ra(&A, &B)
				// fmt.Print("ra")
			}
			if recu[i] == "rb" {
				rb(&A, &B)
				// fmt.Print("rb")
			}
			if recu[i] == "rr" {
				rr(&A, &B)
				// fmt.Print("rr")
			}
			if recu[i] == "rra" {
				rra(&A, &B)
				// fmt.Print("rra")
			}
			if recu[i] == "rrr" {
				rrr(&A, &B)
				// fmt.Print("rrr")
			}
		}
		if sort.IntsAreSorted(A) {
			fmt.Print("OK")
		} else {
			fmt.Print("K0")
		}
		// fmt.Print(A, B)
	} else {
		fmt.Println("entrez une serie de chiffres")
		os.Exit(1)
	}
}

// ss ra rb rr rra rrr
// fmt.Print(count)
// fmt.Print(str, "recu")
// if len(os.Args) > 1 {
// Arg1 := os.Args[1]

// 	if len(Arg1) <= 3 {
// 		fmt.Println("entrez au moins 4 chiffres")
// 		os.Exit(1)
// 	}
// 	if len(Arg1) <= 12 {

// 		A, B := initA(Arg1)
// 		a := initOnlyA(Arg1)

// 		for i := 0; i < len(A); i++ {
// 			preTri(&a, &A, &B)
// 		}

// 		// fmt.Println(A, B, "evo")
// 		for i := 0; i < len(B); i++ {
// 			rendreB(&A, &B)
// 		}

// 		fmt.Println("Resultat trie : ", A, B)
// 		fmt.Println("  Total Operations : ", count)

// 	}
// 	if len(Arg1) > 12 {

// 		A, B := initA(Arg1)
// 		a := initOnlyA(Arg1)

// 		for i := 0; i < len(A); i++ {
// 			preTri(&a, &A, &B)
// 		}
// 		new := A
// 		for i := 0; i < len(B); i++ {
// 			preTri2(&new, &A, &B)
// 		}

// 		// fmt.Print(A, B)
// 		if sort.IntsAreSorted(A) {
// 			for i := 0; i <= len(Arg1); i++ {
// 				rendreB(&A, &B)
// 			}
// 		}

// 		fmt.Println("Resultat trie : ", A, B)
// 		fmt.Println("  Total Operations : ", count)

// 	}
// 	// fmt.Print(len(Arg1))
// } else {
// 	fmt.Print("veuillez entrer un argument")
// }
