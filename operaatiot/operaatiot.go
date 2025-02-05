package operaatiot

import (
	"errors"
	"math"
)

var ErrLiianIsoLuku = errors.New("liian iso luku käytä oikeaa laskinta:)")

func tarkistaTulos(t float64) error {
	if math.IsInf(t, 0) {
		return ErrLiianIsoLuku
	}
	return nil
}
func summa(luvut ...float64) (float64, error) {

	var tulos float64
	for _, luku := range luvut {
		tulos += luku
	}
	if err := tarkistaTulos(tulos); err != nil {
		return 0, err
	}
	return tulos, nil
}

func erotus(luvut ...float64) (float64, error) {

	var tulos float64
	for _, luku := range luvut {
		tulos -= luku
	}
	if err := tarkistaTulos(tulos); err != nil {
		return 0, err
	}
	return tulos, nil
}

func tulo(luvut ...float64) (float64, error) {

	var tulos float64 = 1
	for _, luku := range luvut {
		tulos *= luku
	}
	if err := tarkistaTulos(tulos); err != nil {
		return 0, err
	}
	return tulos, nil
}

func osamaara(luvut ...float64) (float64, error) {
	var tulos float64 = luvut[0]
	for idx, luku := range luvut {
		if idx != 0 && luku == 0 {
			return 0, errors.New("nollalla ei saa jakaa")
		}
		if idx == 0 {
			continue
		}
		tulos /= luku
	}
	if err := tarkistaTulos(tulos); err != nil {
		return 0, err
	}
	return tulos, nil
}

func Operaatiot() map[string]func(...float64) (float64, error) {
	return map[string]func(...float64) (float64, error){
		"+": summa,
		"-": erotus,
		"*": tulo,
		"/": osamaara,
	}
}
