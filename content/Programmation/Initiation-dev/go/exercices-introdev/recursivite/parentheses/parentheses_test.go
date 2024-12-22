package parentheses

import "testing"

func Test0(t *testing.T) {
	res, ok := parentheses(0)
	if !ok {
		t.Error("parentheses(0) devrait retourner true mais retourne false")
	}
	if len(res) != 1 || res[0] != "" {
		t.Error("parentheses(0) devrait retourner une seule chaîne (la chaîne vide), mais retourne", res)
	}
}

func Test1(t *testing.T) {
	_, ok := parentheses(1)
	if ok {
		t.Error("parentheses(1) devrait retourner false mais retourne true")
	}
}

func Test2(t *testing.T) {
	res, ok := parentheses(2)
	if !ok {
		t.Error("parentheses(2) devrait retourner true mais retourne false")
	}
	if len(res) != 1 || res[0] != "()" {
		t.Error("parentheses(2) devrait retourner une seule chaîne : (), mais retourne", res)
	}
}

func Test3(t *testing.T) {
	res, ok := parentheses(6)
	if !ok {
		t.Error("parentheses(6) devrait retourner true mais retourne false")
	}
	if different(res, []string{"()(())", "(())()", "(()())", "((()))", "()()()"}) {
		t.Error("parentheses(6) devrait retourner 5 chaînes : ()()(), ()(()), (())(), (()()) et ((())), mais retourne", res)
	}
}

// fonction auxiliaire pour vérifier les égalités d'ensembles de chaînes

func different(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return true
	}

outer:
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				continue outer
			}
		}
		return true
	}
	return false
}
