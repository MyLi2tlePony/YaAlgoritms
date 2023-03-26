package main

import "fmt"

func main() {
	var studentsNumber, variantsNumber, rowP, placeP int

	_, _ = fmt.Scanln(&studentsNumber)
	_, _ = fmt.Scanln(&variantsNumber)
	_, _ = fmt.Scanln(&rowP)
	_, _ = fmt.Scanln(&placeP)

	indexPrevV := (rowP-1)*2 + placeP - variantsNumber
	indexNextV := (rowP-1)*2 + placeP + variantsNumber

	if indexPrevV > 0 && indexNextV <= studentsNumber {
		prevRow, nextRow := indexPrevV/2, indexNextV/2

		if indexPrevV%2 != 0 {
			prevRow++
		}

		if indexNextV%2 != 0 {
			nextRow++
		}

		if rowP-prevRow < nextRow-rowP {
			printPlace(indexPrevV)
		} else {
			printPlace(indexNextV)
		}
	} else if indexPrevV > 0 {
		printPlace(indexPrevV)
	} else if indexNextV <= studentsNumber {
		printPlace(indexNextV)
	} else {
		fmt.Println(-1)
	}
}

func printPlace(i int) {
	if i%2 == 0 {
		fmt.Printf("%d %d", i/2, 2)
	} else {
		fmt.Printf("%d %d", i/2+1, 1)
	}
}
