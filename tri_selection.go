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
	pivot := tab[n-1]

	tabp := []int{}
	tabg := []int{}
	if n > 1 {
		for i := 0; i < n-1; i++ {
			if tab[i] < pivot {
				tabp = append(tabp, tab[i])
			} else {
				tabg = append(tabg, tab[i])
			}
		}
		tabg = append(tabg, tab[n-1])
		tabg = tri_rapide(tabg)
		tabp = tri_rapide(tabp)
	}
	for k := 0; k < len(tabg); k++ {
		tabp = append(tabp, tabg[k])
	}

	return (tabp)
}

func fusion(tab1 []int, tab2 []int) []int {

	// Voir : https://yourbasic.org/golang/append-explained/

	return ([]int{})
}

func tri_fusion(tab []int) []int {

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
