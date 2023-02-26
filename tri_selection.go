package main

import "fmt"

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
		for i := 0; i < n-1; i++ {
			if tab[i] < pivot {
				tabp = append(tabp, tab[i])
			} else {
				tabg = append(tabg, tab[i])
			}
		}
		tabg = tri_rapide(tabg)
		tabp = tri_rapide(tabp)
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

	// Voir : https://yourbasic.org/golang/append-explained/
	check := (len(tab1) * 2) - 1
	indextab1 := 0
	indextab2 := 0 
	for i := 0; i <= check; i++ {
		if tab1[i] < tab2[i] {
			tab = append(tab, tab1[i])
			indextab1 += 1
		} else {
			tab = append(tab, tab2[i])
			indextab2 += 1
		}

	return (tab)
	//return ([]int{})
}

func tri_fusion(tab []int) []int {

	n := len(tab)
	tabgauche := []int{}
	tabdroite := []int{}
	if n > 1 {

		coupure := n / 2

		tabgauche = append(tabgauche, tab[0:coupure]...)
		tabdroite = append(tabdroite, tab[coupure:n]...)
		if len(tabgauche) > 1 || len(tabdroite) > 1 {
			tri_fusion(tabgauche)
			tri_fusion(tabdroite)
		}
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

}
