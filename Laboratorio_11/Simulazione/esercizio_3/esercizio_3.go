package main

import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "math"

func main() {
	tragitto := NuovoTragitto()
	lunghezza := Lunghezza(tragitto)
	fmt.Printf("Lunghezza percorso: %.3f\n", lunghezza)
	for indice, punto := range tragitto {
		if Lunghezza(tragitto[:indice+1]) > lunghezza/2 {
			fmt.Println("Punto oltre metà:", String(punto))
			break
		}
	}
}

type Punto struct {
	etichetta string
	x, y      float64
}

func NuovoTragitto() (tragitto []Punto) {
	righe := LeggiTesto()
	for _, riga := range righe {
		tragitto = append(tragitto, RigaInPunto(riga))
	}
	return
}

func LeggiTesto() (righe []string) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		righe = append(righe, scanner.Text())
	}
	return
}

func RigaInPunto(riga string) Punto {

	componenti := strings.Split(riga, ";")
	etichetta := componenti[0]
	x, _ := strconv.ParseFloat(componenti[1], 64)
	y, _ := strconv.ParseFloat(componenti[2], 64)

	return Punto{etichetta, x, y}
}

func String(p Punto) string {
	return fmt.Sprintf("%s = (%.1f, %.1f)", p.etichetta, p.x, p.y)
}

func Distanza(p1, p2 Punto) float64 {
	deltax := p1.x - p2.x
	deltay := p1.y - p2.y
	return math.Sqrt(deltax*deltax + deltay*deltay)
}

func Lunghezza(tragitto []Punto) (lunghezza float64) {
	for indice := 1; indice < len(tragitto); indice++ {
		lunghezza += Distanza(tragitto[indice-1], tragitto[indice])
	}
	return
}
