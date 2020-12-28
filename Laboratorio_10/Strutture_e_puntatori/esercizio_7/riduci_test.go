package main

import "testing"

func TestRiduci(t *testing.T) {
	frazione := NuovaFrazione(4, 6)
	Riduci(frazione)
	if s := String(*frazione); s != "2/3" {
		t.Errorf("Riduzione errata: atteso \"2/3\", ottenuto \"%s\"", s)
	}

	frazione = NuovaFrazione(1, 6)
	Riduci(frazione)
	if s := String(*frazione); s != "1/6" {
		t.Errorf("Riduzione errata: atteso \"1/6\", ottenuto \"%s\"", s)
	}

	frazione = NuovaFrazione(6, 6)
	Riduci(frazione)
	if s := String(*frazione); s != "1/1" {
		t.Errorf("Riduzione errata: atteso \"1/1\", ottenuto \"%s\"", s)
	}

	frazione = NuovaFrazione(18, 6)
	Riduci(frazione)
	if s := String(*frazione); s != "3/1" {
		t.Errorf("Riduzione errata: atteso \"3/1\", ottenuto \"%s\"", s)
	}

	frazione = NuovaFrazione(18, 1)
	Riduci(frazione)
	if s := String(*frazione); s != "18/1" {
		t.Errorf("Riduzione errata: atteso \"18/1\", ottenuto \"%s\"", s)
	}
}
