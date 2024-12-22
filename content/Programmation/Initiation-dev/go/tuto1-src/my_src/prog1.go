// Madoc : R101 initiation au dev

package main
import (
	"fmt"
	"math"
)

func sayhello(){
	fmt.Println("Hello!")
}

func hello_var(){
	var text string = "hello with a variable"
	fmt.Println(text)
}

func build_str(){
	var hello, world string = "hello ", "world!"
	var text string = hello + world
	fmt.Println(text)
}

func math_1(x int) int{
	return (6*x + 4)/2
}

func test_sup(x int) bool{
	return x>5
}

func time_from_sec(sec int){
	var min, hour int

	hour = sec / 3600
	sec -= hour * 3600

	min = sec / 60
	sec -= min * 60

	fmt.Println("time :", hour, "h", min, "m", sec, "s.")
}

func sayhello_from_tab(){
	var tab[] string = []string{"hello", "from", "tab"}
	for i := 0; i < len(tab); i ++ {
		fmt.Println(tab[i])
	}
}

func tab3(){
	var tint []int = []int{1, 2, 3, 1000, 27}
	fmt.Println(tint[0])
	fmt.Println(tint[3])
	fmt.Println(tint)
	tint[3] = 0
	fmt.Println(tint)
	fmt.Println("last cell :", tint[len(tint)-1])
	//fmt.Println("outside of the array", tint[50])
}

func table_de_5(){
	for i := 10; i<= 100; i+=5 {
		fmt.Println(i)
	}
}

func visit_array(){
	var tab[] int = []int{0,1,2,3,4,5,6,7,8,9}
	for i := 0; i < len(tab); i ++ {
		fmt.Println(tab[i])
	}
}

func test_for2(){
	var i int
	for i = 0; i < 10; i = i + 2 {
		fmt.Println(i)
	}
	fmt.Println(i)
}

func add_array(){
	var tab[] int = []int{0,1,2,3,4,5,6,7,8,9}
	var result int
	for i := 0; i < len(tab); i ++ {
		result += tab[i]
	}
	fmt.Println(result)
}

func default_float(){
	var virgule float64
	fmt.Println(virgule)
}

func clac_moy(){
	var tab[] int = []int{0,1,2,3,4,5,6,7,8,9}
	var result int
	for i := 0; i < len(tab); i ++ {
		result += tab[i]
	}
	fmt.Println("moyenne :", result/len(tab))
}

func prod(x, y int) int{
	var result int
	for x>0 {
		result += y
		x--
	}
	return result
}

func money(age int) float64{
	var money float64 = 100
	for i := 1; i <= age; i++{
		money *= 1.035 			// taux d'interet
		money += float64(i*10) 	// argent de papy
	}
	return money
}

func remise(prix float64){
	if prix > 300 {
		prix *= 0.95
	}
	fmt.Println("Prix payé :", math.Floor(prix))
}

func search_highest(){
	var liste[] int = []int{0,1,2,3,4,5,6,7,8,9}
	var result int
	for i:=0; i<= len(liste)-1; i++ {
		if result < liste[i]{
			result = liste[i]
		}
	}
	fmt.Println(result, "est le plus haut dans la liste.")
}

func test_signe(x int){
	switch{
	case x < 0: fmt.Println("négatif")
	case x == 0: fmt.Println("nul")
	case x > 0: fmt.Println("positif")
	default : fmt.Println("erreur")
	}
}

func main(){
	//sayhello()
	//hello_var()
	//build_str()
	//fmt.Println("input 5, return :", math_1(5))
	//fmt.Println("input 7, return :", math_1(7))
	//fmt.Println("input 2 return:", test_sup(2), "\ninput 8 return:", test_sup(8))
	//time_from_sec(154645)
	//sayhello_from_tab()
	//tab3()
	//table_de_5()
	//visit_array()
	//test_for2()
	//add_array()
	//default_float()
	//clac_moy()
	//fmt.Println(prod(2,8))
	//fmt.Println("Age = 1, money =", money(1), "\nAge = 4, money =", money(4))
	//remise(305)
	//search_highest()
	test_signe(-1); test_signe(0); test_signe(1);
	
	// Question 1.5.4
}