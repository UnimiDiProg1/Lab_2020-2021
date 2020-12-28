package main

import "testing"

func TestNuovaFrazione(t *testing.T) {
	frazione := NuovaFrazione(2, 5)
	if frazione == nil {
		t.Error("Frazione non creata")
		return
	}

	if frazione.numeratore != 2 {
		t.Error("Numeratore non inizializzato correttamente")
	}

	if frazione.denominatore != 5 {
		t.Error("Denominatore non inizializzato correttamente")
	}
}

func TestString(t *testing.T) {
	frazione := NuovaFrazione(2, 5)
	if frazione == nil {
		t.Error("Frazione non creata")
		return
	}

	s := String(*frazione)
	if s != "2/5" {
		t.Errorf("Stampa non corretta: atteso \"2/5\", ottenuto \"%s\"", s)
	}
}
