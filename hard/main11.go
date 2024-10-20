package main

import (
	"fmt"
	"strings"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}
func Permutations(input string) []string {
	mass := []string{}
	iters := factorial(len(input))
	//z := strings.Split(input[1:], "")
	//z[2] = z[1]
	//fmt.Println(z)
	for i := 0; i <= len(input)-1; i++ {
		for j := 0; j <= iters/len(input); j++ {
			z := strings.Split(input[1:], "")
			for k := len(z); k > 1; k-- {
				z[k-1], z[k-2] = z[k-2], z[k-1]
				mass = append(mass, string([]rune(input)[i])+strings.Join(z, ""))
			}
		}
	}
	return mass
}

func main() {
	fmt.Println(Permutations("abc"))
}
