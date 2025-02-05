package laskin

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/waeekron/demo-2-golang-calculator/operaatiot"
)

type Käly struct {
	laskin Laskija
}

func NewKäly() *Käly {
	return &Käly{laskin: NewLaskin(operaatiot.Operaatiot())}
}
func (k Käly) Aloita() {
	k.kysyLuvut(k.laskin.AsetaLuvut)
	k.kysyOperaatio(k.laskin.TuetutOperaatiot(), k.laskin.Laske)
	k.tulos(k.laskin.Tulos)
}

func (k Käly) tulos(tulos func() float64) {
	fmt.Printf("Tulos: %f\n", tulos())
}
func (k Käly) kysyLuvut(asetaLuvutCB func(float64, float64)) {
	luku1 := k.kysyLuku("Anna ensimmäinen luku")
	luku2 := k.kysyLuku("Anna toinen luku")

	asetaLuvutCB(luku1, luku2)
}

func (k Käly) kysyLuku(kehote string) float64 {
	var luku float64
	for {
		println(kehote)
		lukija := bufio.NewReader(os.Stdin)
		mj, _ := lukija.ReadString('\n')
		luvut := strings.ReplaceAll(mj, " ", "")
		luvut = strings.TrimSpace(luvut)
		l, err := strconv.ParseFloat(luvut, 64)
		if err == nil {
			luku = l
			break
		}
		fmt.Println(err)
	}
	return luku
}

func (k Käly) kysyOperaatio(tuetutOperaatiot []string, laskeOperaatioCB func(string) error) {
	for {
		fmt.Printf("Anna operaatio: %s", tuetutOperaatiot)
		var lt string
		fmt.Scanf("%s", &lt)
		err := laskeOperaatioCB(lt)
		if err == nil {
			break
		}
		fmt.Printf("Virhe: %s\n", err.Error())

	}
}
