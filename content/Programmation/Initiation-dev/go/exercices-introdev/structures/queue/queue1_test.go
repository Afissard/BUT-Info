package queue

import "testing"

func TestQueue(t *testing.T) {
	p1 := place{
		occupant:  personne{nom: "Torvalds", prenom: "Linus", age: 52},
		precedant: nil,
	}
	p2 := place{
		occupant:  personne{nom: "Sacquet", prenom: "Bilbon", age: 50},
		precedant: &p1,
	}
	p3 := place{
		occupant:  personne{nom: "Djuna", prenom: "Gandhi", age: 36},
		precedant: &p2,
	}
	q := queue{fin: &p3, taille: 3}
	nouveau := personne{
		nom:    "Nakamura",
		prenom: "Aya",
		age:    27,
	}
	q.ajout(nouveau)
	if q.taille != 4 {
		t.Error("Après ajout d'une personne la taille de la queue devrait être 4 mais elle vaut", q.taille)
		return
	}

	if q.fin == nil {
		t.Error("La place ajoutée dans la queue n'est pas correcte, elle vaut nil")
		return
	}

	if q.fin.occupant != nouveau {
		t.Error("La dernière personne dans la queue devrait être", nouveau, "mais c'est", q.fin.occupant)
	}
}

// ajouté après le test

func TestQueue2(t *testing.T) {
	q := queue{fin: nil, taille: 0}
	var record []*place
	for i := 0; i < 100; i++ {
		q.ajout(personne{nom: "N", prenom: "P", age: 0})
		if q.fin == nil {
			t.Fail()
		}
		record = append(record, q.fin)
		for i := len(record) - 1; i > 0; i-- {
			if record[i].precedant != record[i-1] {
				t.Fail()
			}
		}
	}
}
