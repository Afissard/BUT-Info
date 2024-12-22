package queue

import "testing"

func TestQueue(t *testing.T) {
	prems := personne{nom: "Torvalds", prenom: "Linus", age: 52}
	p1 := place{
		occupant:  prems,
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
	p4 := place{
		occupant:  personne{nom: "Nakamura", prenom: "Aya", age: 27},
		precedant: &p3,
	}
	q := queue{fin: &p4, taille: 4}
	premier := q.premier()
	if premier != prems {
		t.Error("Vous dites que la première personne dans la queue est", premier, "alors que c'est", prems)
	}
}

// ajouté après le test machine

func TestIntegrity(t *testing.T) {
	prems := personne{nom: "Torvalds", prenom: "Linus", age: 52}
	p1 := place{
		occupant:  prems,
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
	p4 := place{
		occupant:  personne{nom: "Nakamura", prenom: "Aya", age: 27},
		precedant: &p3,
	}
	p5 := place{
		occupant:  personne{nom: "Nakamura", prenom: "Aya2", age: 27},
		precedant: &p4,
	}
	q := queue{fin: &p5, taille: 5}
	q.premier()
	if q.fin != &p5 {
		t.Fail()
	}
	if p5.precedant != &p4 {
		t.Fail()
	}
	if p4.precedant != &p3 {
		t.Fail()
	}
	if p3.precedant != &p2 {
		t.Fail()
	}
	if p2.precedant != &p1 {
		t.Fail()
	}
	if p1.precedant != nil {
		t.Fail()
	}
}
