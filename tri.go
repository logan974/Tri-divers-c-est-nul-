package main

import (
	"fmt"
)

func tri_selection(tab []int) []int {

	n := len(tab)

	for i := 0; i <= n-2; i++ {
		min := i
		for j := i + 1; j <= n-1; j++ {
			if tab[j] < tab[min] {
				min = j
			}
		}
		if min != i {
			temp := tab[i]
			tab[i] = tab[min]
			tab[min] = temp

		}
	}

	return (tab)
}

func tri_bulles(tab []int) []int {
	n := len(tab)
	for i := n - 1; i >= 1; i-- {
		tableau_trié := true
		for j := 0; j <= i-1; j++ {
			if tab[j+1] < tab[j] {
				temp := tab[j]
				tab[j] = tab[j+1]
				tab[j+1] = temp
				tableau_trié = false
			}
		}
		if tableau_trié {
			break
		}
	}

	return (tab)
}

func tri_insertion(tab []int) []int {

	n := len(tab)
	x := 0

	for i := 1; i <= n-1; i++ {
		j := i
		x = tab[i]
		for j > 0 && tab[j-1] > x {
			tab[j] = tab[j-1]
			j -= 1
		}
		tab[j] = x
	}

	return (tab)
}

func tri_rapide(tab []int) []int {

	// Voir : https://yourbasic.org/golang/append-explained/
	n := len(tab)

	tabp := []int{} //tab petit
	tabg := []int{} //tab grand
	if n > 1 {
		pivot := tab[n-1]
		//Création des 2 tableaux inférieur et supérieur au pivot
		for i := 0; i < n-1; i++ {
			if tab[i] < pivot {
				tabp = append(tabp, tab[i])
			} else {
				tabg = append(tabg, tab[i])
			}
		}
		//Récursivité
		tabg = tri_rapide(tabg)
		tabp = tri_rapide(tabp)
		//Fusion de la récursivité
		tabp = append(tabp, pivot)
		for k := 0; k < len(tabg); k++ {
			tabp = append(tabp, tabg[k])
		}
	} else {
		tabp = tab
	}

	return (tabp)
}

func fusion(tab1 []int, tab2 []int) []int {
	tab := []int{}
	//Check par rapport à la taille du tableau finale
	//check := (len(tab1) * 2) - 1
	check := len(tab1) + len(tab2) - 1
	indextab1 := 0
	indextab2 := 0
	for i := 0; i <= check; i++ {
		//Condition : Les deux tableaux ont des valeurs restante à traiter
		if indextab1 < len(tab1) && indextab2 < len(tab2) {
			//Tri
			if tab1[indextab1] < tab2[indextab2] {
				tab = append(tab, tab1[indextab1])
				indextab1 += 1
			} else {
				tab = append(tab, tab2[indextab2])
				indextab2 += 1
			}
		}
		//Condition : Un des tableau a des valeurs restante à traiter
		if indextab1 >= len(tab1) && indextab2 < len(tab2) {
			tab = append(tab, tab2[indextab2])
			indextab2 += 1
		}
		if indextab2 >= len(tab2) && indextab1 < len(tab1) {
			tab = append(tab, tab1[indextab1])
			indextab1 += 1
		}
		//Si plus de valeur à traiter
		if indextab1 >= len(tab1) && indextab2 >= len(tab2) {
			break
		}

		//return ([]int{})
	}
	return (tab)
}

func tri_fusion(tab []int) []int {

	n := len(tab)
	tabgauche := []int{}
	tabdroite := []int{}
	// On vérifie que la tableau à au moins 2 élément
	if n > 1 {
		coupure := n / 2
		//Coupure
		tabgauche = append(tabgauche, tab[0:coupure]...)
		tabdroite = append(tabdroite, tab[coupure:n]...)
		//Récursivité
		if len(tabgauche) > 1 || len(tabdroite) > 1 {
			tabgauche = tri_fusion(tabgauche)
			tabdroite = tri_fusion(tabdroite)
		}
		//Fusion
		tab = fusion(tabgauche, tabdroite)

	}
	return (tab)
}

func main() {

	tab := []int{408, 9873, 35, 12, 16, 66, 77, 321, 788, 2, 45, 8, 55, 1244, 1987, 1762}

	fmt.Println("Tri selection:", tri_selection(tab))

	tab = []int{408, 9873, 35, 12, 16, 66, 77, 321, 788, 2, 45, 8, 55, 1244, 1987, 1762}

	fmt.Println("Tri bulles:", tri_bulles(tab))

	tab = []int{408, 9873, 35, 12, 16, 66, 77, 321, 788, 2, 45, 8, 55, 1244, 1987, 1762}

	fmt.Println("Tri insertion:", tri_insertion(tab))

	tab = []int{408, 9873, 35, 12, 16, 66, 77, 321, 788, 2, 45, 8, 55, 1244, 1987, 1762}

	fmt.Println("Tri rapide:", tri_rapide(tab))

	tab = []int{408, 9873, 35, 12, 16, 66, 77, 321, 788, 2, 45, 8, 55, 1244, 1987, 1762}

	fmt.Println("Tri fusion:", tri_fusion(tab))

	// Pause a la fin
	//bufio.NewReader(os.Stdin).ReadBytes('\n')

}
