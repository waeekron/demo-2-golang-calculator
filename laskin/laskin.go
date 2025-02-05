package laskin

import (
	"fmt"

	"golang.org/x/exp/maps"
)

type Laskija interface {
	AsetaLuvut(float64, float64)
	Laske(operaatio string) error
	Tulos() float64
	TuetutOperaatiot() []string
}
type Laskin struct {
	luku1, luku2 float64
	tulos        float64
	operaatiot   map[string]func(...float64) (float64, error)
	virhe        error
}

func (l *Laskin) Tulos() float64 {
	return l.tulos
}
func (l *Laskin) AsetaLuvut(l1, l2 float64) {
	l.luku1 = l1
	l.luku2 = l2
}
func (l *Laskin) TuetutOperaatiot() []string {
	return maps.Keys(l.operaatiot)
}

func NewLaskin(op map[string]func(...float64) (float64, error)) *Laskin {
	return &Laskin{
		operaatiot: op,
	}
}
func (l *Laskin) Laske(operaatio string) error {
	l.suoritaOperaatio(operaatio)
	if l.virhe != nil {
		return l.virhe
	}
	return nil
}
func (l *Laskin) suoritaOperaatio(operaatio string) {
	op, ok := l.operaatiot[operaatio]
	if !ok {
		l.virhe = fmt.Errorf("laskin ei tue operaatiota")
		return
	}
	tulos, err := op(l.luku1, l.luku2)
	if err != nil {
		l.virhe = err
	}
	l.tulos = tulos

}
