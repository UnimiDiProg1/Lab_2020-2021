package main

import "testing"

func TestMoltiplica(t *testing.T) {
	frazione1 := NuovaFrazione(4, 6)
	frazione2 := NuovaFrazione(3, 6)

	prodotto := Moltiplica(*frazione1, *frazione2)
	if s := String(*prodotto); s != "12/36" {
		t.Errorf("Il prodotto non è corretto: atteso \"12/36\", ottenuto \"%s\"", s)
	}

	frazione1 = NuovaFrazione(0, 6)
	frazione2 = NuovaFrazione(3, 3)

	prodotto = Moltiplica(*frazione1, *frazione2)
	if s := String(*prodotto); s != "0/18" {
		t.Errorf("Il prodotto non è corretto: atteso \"0/18\", ottenuto \"%s\"", s)
	}
}

func TestMoltiplicaN(t *testing.T) {
	frazioni := []Frazione{
		*NuovaFrazione(4, 6),
		*NuovaFrazione(1, 3),
		*NuovaFrazione(5, 2),
	}

	prodotto := MoltiplicaN(frazioni)
	if s := String(*prodotto); s != "20/36" {
		t.Errorf("Il prodotto non è corretto: atteso \"20/36\", ottenuto \"%s\"", s)
	}
}
