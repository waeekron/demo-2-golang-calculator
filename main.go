package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Laskettavat luvut
	var luku1, luku2 float64

	// Laskutoimitus
	var lt string

	// Kysy ensimmäistä lukua
	for {
		fmt.Println("Anna ensimmäinen luku:")
		_, err := fmt.Fscan(os.Stdin, &luku1)
		if err == nil {
			break
		} else {
			fmt.Println("Syöte ei kelpaa!")
			fmt.Printf("Virhe: %s\n\n", err)
		}
	}

	// Kysy toista lukua
	for {
		println("Anna toinen luku")
		lukija := bufio.NewReader(os.Stdin)
		mj, _ := lukija.ReadString('\n')
		luvut := strings.ReplaceAll(mj, " ", "")
		luvut = strings.TrimSpace(luvut)
		println(luvut)
		l, err := strconv.ParseFloat(luvut, 64)
		if err == nil {
			luku2 = l
			break
		}
		fmt.Println(err)
	}

	// Kysy laskutoimitus.
	fmt.Println("Anna laskutoimitus (+, -, /, *):")
	fmt.Scanf("%s", &lt)

	// Tulostetaan, mitä luettiin.
	fmt.Println("Lasken: ", luku1, lt, luku2)

	// Tulos sisältää laskutoimituksen tuloksen.
	var tulos float64

	// Virhe sisältää kuvauksen virheestä, mutta on tyhjä, jos virhettä ei tapahtunut.
	var virhe string

	// Suoritetaan laskutoimituksen mukainen laskutoimitus.
	switch lt {
	case "+":
		tulos = luku1 + luku2
	case "-":
		tulos = luku1 - luku2
	case "/":
		if luku2 != 0 {
			tulos = luku1 / luku2
		} else {
			virhe = "Luvulla 0 ei jaa jakaa!"
		}
	case "*":
		tulos = luku1 * luku2

	default:
		tulos = luku1 + rand.Float64()*luku2
		virhe = "En tunnistanut laskutoimitusta '" + lt + "'. " +
			"Tässä sinulle satunnaisluku: " +
			strconv.FormatFloat(tulos, 'G', 3, 64)
	}

	// Tulostetaan laskutoimituksen tulos tai virhe.
	if virhe != "" {
		fmt.Println("Virhe:", virhe)
	} else {
		fmt.Println("Tulos on", strconv.FormatFloat(tulos, 'f', -1, 64))
	}
}
