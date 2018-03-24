package main

import (
	"fmt"
)

func main() {

	const latticeSize = 10
	const dimensions = 2
	// totalLatticeSize := Pow(latticeSize, dimensions)
	const spawnProbability = 0.05
	const eatingProbability = 0.5 // "Fresswahrscheinlichkeit" - probability that species can invade neighbour
	const simulationEndTime = 100 // number of simulation steps

	// Lattice of the simulation.
	// One dimension array representing multiple dimensions.
	// Int Values represent species, e.g. 3 represents species "3".
	// starting condition: 1 species with id "0". Golang initializes with zeroes, which suites this case perfectly.
	var lattice = make([]int, latticeSize)

	var interactionMatrix = [][]bool{[]bool{false}} // starting condition: 1 species, can't eat itself
	fmt.Print(interactionMatrix)
	var activeLinks [][2]int

	for t := 0; t < simulationEndTime; t++ {
		Battle(&lattice, interactionMatrix, activeLinks)
		Spawn(spawnProbability, &lattice, &interactionMatrix, eatingProbability)
	}

}

// A link joins two cells, e.g. 2 <-> 3. => a representation of a link is an array of length 2.
func getActiveLinks(lattice []int, latticeSize int, dimensions int, interactionMatrix [][]bool) [][2]int {
	var activeLinks [][2]int
	for cell := 0; cell < latticeSize; cell++ {
		for dim := 0; dim < dimensions; dim++ {
			upperNeighbor := UpperNeighbor(cell, dim, latticeSize)
			if CanEat(upperNeighbor, cell, interactionMatrix) || CanEat(cell, upperNeighbor, interactionMatrix) {
				activeLinks = append(activeLinks, [2]int{lattice[cell], lattice[upperNeighbor]})
			}
		}
	}
	return activeLinks
}

func CanEat(speciesA int, speciesB int, interactionMatrix [][]bool) bool {
	return interactionMatrix[speciesA][speciesB]
}
