package main

import (
	"fmt"
	"math"
	"math/rand"
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
	var activeLinks []int

	for t := 0; t < simulationEndTime; t++ {
		battle(lattice, interactionMatrix, activeLinks)
		Spawn(spawnProbability, lattice, interactionMatrix, eatingProbability)
	}

}

func battle(lattice []int, interactionMatrix [][]bool, activeLinks []int) {
	for _, link := range activeLinks {
		fmt.Println(link, "battling!")
	}
}

// A link joins two cells, e.g. 2 <-> 3. => a representation of a link is an array of length 2.
func getActiveLinks(lattice []int, latticeSize int, dimensions int, interactionMatrix [][]bool) [][2]int {
	var activeLinks [][2]int
	for cell := 0; cell < latticeSize; cell++ {
		for dim := 0; dim < dimensions; dim++ {
			upperNeighbor := UpperNeighbor(cell, dim, latticeSize)
			if canEat(upperNeighbor, cell, interactionMatrix) || canEat(cell, upperNeighbor, interactionMatrix) {
				activeLinks = append(activeLinks, [2]int{lattice[cell], lattice[upperNeighbor]})
			}
		}
	}
	return activeLinks
}

func canEat(speciesA int, speciesB int, interactionMatrix [][]bool) bool {
	return interactionMatrix[speciesA][speciesB]
}

func AddSpeciesToInteractionMatrix(interactionMatrix *[][]bool, eatingProbability float64) {
	var newSpeciesEats []bool
	interactionMatrixPtr := *interactionMatrix
	var newSpeciesNumber = len(interactionMatrixPtr)
	for i := 0; i < newSpeciesNumber; i++ {
		newSpeciesEats = append(newSpeciesEats, eatingProbability > rand.Float64())
		newSpeciesGetsEatenByI := eatingProbability > rand.Float64()
		interactionMatrixPtr[i] = append(interactionMatrixPtr[i], newSpeciesGetsEatenByI)
	}
	newSpeciesEats = append(newSpeciesEats, false) // cannot eat itself
	interactionMatrixPtr = append(interactionMatrixPtr, newSpeciesEats)
}

func Spawn(spawnProbability float64, lattice []int, interactionMatrix [][]bool, eatingProbability float64) ([]int, [][]bool) {
	latticeSize := len(lattice)
	latticeCopy := make([]int, len(lattice))
	copy(latticeCopy, lattice)
	if spawnProbability > rand.Float64() {
		newSpeciesNumber := len(interactionMatrix)
		randomCoordinate := rand.Intn(latticeSize - 1)
		latticeCopy[randomCoordinate] = newSpeciesNumber
		AddSpeciesToInteractionMatrix(&interactionMatrix, eatingProbability)
		return lattice, interactionMatrix
	}
	return nil, nil
}

func getUpperNeighbors(cell int, dimensions int, latticeSize int) []int {
	var upperNeighbors []int
	for i := 0; i < dimensions; i++ {
		upperNeighbors = append(upperNeighbors, UpperNeighbor(cell, i, latticeSize))
	}
	return upperNeighbors
}

func UpperNeighbor(cell int, dimension int, latticeSize int) int {
	neighbor := Pow(latticeSize, dimension-1)
	rangelength := Pow(latticeSize, dimension)
	rangestart := cell / rangelength * rangelength
	return (cell+neighbor)%rangelength + rangestart
}

// That's a bit lame that I have to write stuff like that?
// Also the concept of "just make a capital letter to export it" I don't find the best
// You'll have to modify more places than by just using a private / public modifier...
func Pow(base int, exp int) int {
	return int(math.Pow(float64(base), float64(exp)))
}

func abs(num int) int {
	return int(math.Abs(float64(num)))
}
