package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	const latticeSize = 10
	const dimensions = 2
	const spawnProbability = 0.05
	const p = 0.5                 // "Fresswahrscheinlichkeit" - probability that species can invade neighbour
	const simulationEndTime = 100 // number of simulation steps

	// Lattice of the simulation.
	// One dimension array representing multiple dimensions.
	// Int Values represent species, e.g. 3 represents species "3".
	var lattice = make([]int, latticeSize)

	var G = make([][]int, 1)
	var currentNumberOfSpecies = 1 // Start simulation with one species
	var activeLinks []int

	for t := 0; t < simulationEndTime; t++ {
		battle(lattice, G, activeLinks)
		spawn(spawnProbability, lattice, currentNumberOfSpecies)
	}

}

func battle(lattice []int, G [][]int, activeLinks []int) {
	for _, link := range activeLinks {
		fmt.Println(link, "battling!")
	}
}

func getActiveLinks(lattice []int, latticeSize int, dimension int) []int {
	var activeLinks []int
	//var linkCandidates []int
	// e.g. next neighbors 8 are 9, 18 and 108 with lattice size 10 and 3 dims
	// var neighbours []int
	for i := 0; i < dimension; i++ {
		//neighbours[i] = i + math.Pow10(dimension)
	}
	return activeLinks
}

func spawn(spawnProbability float64, lattice []int, currentNumberOfSpecies int) {
	latticeSize := len(lattice)
	if spawnProbability > rand.Float64() {
		newSpeciesNumber := currentNumberOfSpecies
		randomCoordinate := rand.Intn(latticeSize - 1)
		lattice[randomCoordinate] = newSpeciesNumber
	}
}

func getNeighbours(cell int, dimensions int, latticeSize int) []int {
	var neighbours []int
	// for i := 0; i < dimension; i++ {
	// 	upperCell := i + math.Pow10(dimension) // Pow 10 ... bad idea :)
	// 	lowerCell := i - math.Pow10(dimension)

	// 	if !isLowerBoundary {

	// 	}
	// 	if !isUpperBoundary {

	// 	}
	// 	append(neighbours)
	// 	neighbours[i] = i + math.Pow10(dimension)

	// }
	return neighbours
}

// e.g. in a 10x10 lattice:
// 00 10 20 ... 90 are lower boundary cells of dim 1
// 00 01 02 ... 09 of dim 2
func CellIsLowerBoundary(cell int, dimension int, latticeSize int) bool {
	dimensionAdjustment := int(math.Pow(float64(latticeSize), float64(dimension-1)))
	return cell/dimensionAdjustment%latticeSize == 0
}
