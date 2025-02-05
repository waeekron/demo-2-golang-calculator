package operaatiot

import (
	"errors"
	"math"
	"testing"
)

type Result struct {
	value float64
	err   error
}

func TestSumma(t *testing.T) {
	testit := []struct {
		nimi  string
		syöte []float64
		want  Result
	}{
		{"kahden luvun summa", []float64{1, 2}, Result{3, nil}},
		{"kolmen luvun summa", []float64{1, 2, 3.5}, Result{6.5, nil}},
		{"Virhe kun summa menee yli tuetun numero tyypin", []float64{math.MaxFloat64, math.MaxFloat64}, Result{0, ErrLiianIsoLuku}},
	}

	for _, tp := range testit {
		got, err := summa(tp.syöte...)
		if got != tp.want.value || !errors.Is(err, tp.want.err) {
			t.Fatalf(`got: %v - %v, want: %v - %v`, got, err, tp.want.value, tp.want.err)
		}
	}
}
