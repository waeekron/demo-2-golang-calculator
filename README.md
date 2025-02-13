# Demo 2
Golang kielen alkeiden harjoittelua varten:)

### laskukone
- main-haarasta löytyy simppeli versio

### laskukone aliohjelmilla
- subriitinifoitu-haarasta löytyy versio, joka on jaettu pää- ja aliohjelmiin (tai itseasiassa menty vähän pidemmälle jakamalla ohjelma tyyppeihin, metodeihin sekä paketteihin)
  - `Käly`-tietue laskee laskutoimitukset `Laskija`-rajapinnan avulla ja huolehtii syöte/tulostus operaatioista
  - laskin tietue toteuttaa `Laskija`-rajapinnan
  - `operaatiot`-paketti sisältää laskutoimitukset, jotka asetetaan laskimelle
    - paketti exportaa vain funktion, joka palauttaa laskutoimitusfunktiot sisältävän map-tietorakenteen (tämä tehty lähinnä siksi, että halusin kokeilla näkyvyyden rajoittamista:)
