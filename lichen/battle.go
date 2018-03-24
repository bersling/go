package main

import (
	"math/rand"
)

func Battle(lattice *[]int, interactionMatrix [][]bool, activeLinks [][2]int) {
	for _, link := range activeLinks {
		cellA := link[0]
		speciesA := (*lattice)[cellA]
		cellB := link[1]
		speciesB := (*lattice)[cellB]
		aCanEatB := CanEat(speciesA, speciesB, interactionMatrix)
		bCanEatA := CanEat(speciesB, speciesA, interactionMatrix)
		if aCanEatB && bCanEatA {
			aEatsB := rand.Float64() > 0.5
			if aEatsB {
				Eat(cellA, cellB, lattice)
			} else {
				Eat(cellB, cellA, lattice)
			}
		} else if aCanEatB {
			Eat(cellA, cellB, lattice)
		} else if bCanEatA {
			Eat(cellB, cellA, lattice)
		} else {
			// program error
		}
	}
}

func Eat(winnerCell int, loserCell int, lattice *[]int) {
	winnerSpecies := (*lattice)[winnerCell]
	(*lattice)[loserCell] = winnerSpecies
	// TODO: update active links!!!
}
