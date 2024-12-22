/*
Étant donné un tableau t contenant des caractères et un entier k<=len(t),
donner toutes les chaînes de k caractères que l'on peut construire en utilisant
au plus une fois chaque caractère de t.
*/
package main

import "fmt"

// Fonction auxiliaire utilisée dans la version récursive et la version
// itérative pour construire un tableau qui contient tous les éléments
// de t à part celui d'indice i, donc pour "enlever" l'élèment d'indice i
// du tableau t
func remove(i int, t []string) []string {
	result := make([]string, len(t)-1)
	// On recopie tous les éléments de t qui sont avant la position i
	for j := 0; j < i; j++ {
		result[j] = t[j]
	}
	// On recopie en les décalant d'une case vers la gauche tous les éléments
	// de t qui sont après la position i
	for j := i; j < len(result); j++ {
		result[j] = t[j+1]
	}
	return result
}

// Fonction récursive pour résoudre le problème
func recur(t []string, k int) []string {
	// Initialisation : quand k = 0 on retourne directement un résultat
	if k == 0 {
		return []string{""}
	}
	// Récurence : quand k > 0
	result := make([]string, 0)
	// Pour chaque élément de t, qu'on appelle char
	for charPos, char := range t {
		// On résout le problème pour le tableau t privé de cet élément et k-1
		// cela nous donne toutes les chaînes de k-1 lettres qui ne contiennent
		// pas char
		prefixes := recur(remove(charPos, t), k-1)
		// On ajoute char au début de chacune de ces chaînes
		for _, prefix := range prefixes {
			result = append(result, prefix+char)
		}
	}
	return result
}

// Fonction itérative pour résoudre le problème
func iter(t []string, k int) []string {
	// On initialise un tableau ne contenant que la chaîne vide (c'est ce qu'on
	// va retourner si k=0)
	result := []string{""}
	// On initialise un tableau contenant t, l'idée est que chaque élément de ce
	// tableau est associé à l'élément du tableau result (une chaîne de caractères)
	// qui a le même indice et indique les lettres qui peuvent être ajoutés à cette
	// chaîne (celle qui n'y sont pas déjà)
	associatedt := [][]string{t}
	// On va construire toutes les chaînes de 1 lettre, puis les étendre en
	// chaînes de 2 lettres, puis 3 et ainsi de suite jusqu'à k lettres
	for charNum := 0; charNum < k; charNum++ {
		newResult := make([]string, 0)
		newAssociatedt := make([][]string, 0)
		// Pour chacune des chaînes construites au tour de boucle précédent
		for prefixPos, prefix := range result {
			// Pour chaque lettre disponible dans le tableau associé à cette chaîne
			for charPos, char := range associatedt[prefixPos] {
				// On ajoute le charactère à la chaîne et on construit le nouveau tableau
				// associé à cette chaîne. On stocke ces éléments temporairement.
				newResult = append(newResult, prefix+char)
				newAssociatedt = append(newAssociatedt, remove(charPos, associatedt[prefixPos]))
			}
		}
		// Les éléments qui ont été stockés temporairement remplacent les contenus
		// de result et associatedt : on remplace les chaînes de longueur 0 par des
		// chaînes de longueur 1 au premier tour de boucle, puis longueur 1 par
		// longueur 2 au deuxième tour, etc.
		result = newResult
		associatedt = newAssociatedt
	}
	return result
}

func main() {

	t := []string{"a", "b", "c", "d"}

	fmt.Println("Fonction récursive")
	fmt.Println("k=2 : ", recur(t, 2))
	fmt.Println("k=3 : ", recur(t, 3))

	fmt.Println("Fonction itérative")
	fmt.Println("k=2 : ", iter(t, 2))
	fmt.Println("k=3 : ", iter(t, 3))

}
