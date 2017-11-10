package main

import "fmt"

func main() {
	sort := [3]string{"Q", "A", "2"}
	var sorted []string
	cards := [13]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "K", "Q"}

	for i := 0; i <= 12; i++ {
		f := cards[i]
		for x := 0; x<=2; x++{
			g := sort[x]
			if f == g{
				sorted = append(sorted, f)
			}
		}
	}

	fmt.Println(sorted)
}
