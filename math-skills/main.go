package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Average(conve []float64) float64 {
	moconveenne := 0.0
	sum := 0.0
	m := len(conve)
	for _, i := range conve {
		sum += i
	}
	moconveenne = sum / float64(m)
	return math.Round(moconveenne)
}

func Mediane(conve []float64) float64 {
	mediane := 0.0
	sort.Slice(conve, func(i, j int) bool {
		return conve[i] < conve[j]
	})

	if len(conve)%2 == 0 {
		mediane = (conve[(len(conve)-1)/2] + conve[len(conve)/2]) / 2
	} else {
		mediane = conve[len(conve)/2]
	}
	return mediane
}

func Variance(conve []float64) float64 {
	variance := 0.0

	for _, k := range conve {

		s := k - Average(conve)
		variance += math.Pow(s, 2) / float64(len(conve))

	}
	return math.Round(variance)
}

func standard_deviation(conve []float64) float64 {
	ecarT := math.Sqrt(Variance(conve))
	return math.Round(ecarT)
}

func main() {
	var conve []float64
	content, err := os.ReadFile("data.txt") // lire le fichier
	if err != nil {
		fmt.Println(err)
		return
	}

	cont := string(content)
	// fmt.Print(cont)
	str := strings.Split(cont, "\n")
	// fmt.Print(str)
	for i := 0; i < len(str)-1; i++ {

		conv, _ := strconv.ParseFloat(str[i], i)

		conve = append(conve, conv)

	}
	fmt.Println("average :", math.Round(Average(conve)))
	fmt.Println("mediane : ", math.Round((Mediane(conve))))
	fmt.Println("variance : ", int((Variance(conve))))
	fmt.Println("Standard deviation : ", math.Round((standard_deviation(conve))))
}
