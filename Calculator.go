package main

import (
	"fmt"
)

func main() {
	fmt.Print("Berikan input string: ")
	var kalimat string
	fmt.Scanf("%s", &kalimat)

	mapSTRING := make(map[string]int)
	for _, kata := range kalimat {
		if kata == ' ' {
			continue
		}
		huruf := string(kata)
		mapSTRING[huruf]++
	}

	for huruf, jumlah := range mapSTRING {
		fmt.Printf("%s: %d ", huruf, jumlah)
	}
}
