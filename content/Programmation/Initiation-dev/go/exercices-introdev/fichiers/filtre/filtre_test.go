package filtre

import "testing"

func TestPasLa(t *testing.T) {
	_, err := filtrePairs("test")
	if err != errImpossible {
		t.Error("Le fichier test n'existe pas, il faut retourner errImpossible")
	}
}

func TestVide(t *testing.T) {
	pairs, err := filtrePairs("../fichiers-tests/filtre-test2")
	if err != nil {
		t.Error("Le fichier test2 existe, il ne faut pas retourner d'erreur")
		return
	}
	if pairs != "" {
		t.Error("Vous retournez \"" + pairs + "\" au lieu de \"\"")
	}
}

func TestOk(t *testing.T) {
	pairs, err := filtrePairs("../fichiers-tests/filtre-test3")
	if err != nil {
		t.Error("Le fichier test3 existe, il ne faut pas retourner d'erreur")
		return
	}
	if pairs != "hola\n\nyo\ncoucou" {
		t.Error(
			"Le fichier test3 filtré doit donner :\n"+
				"hola\n\nyo\ncoucou\n",
			"mais vous retournez :\n"+pairs)
	}
}
