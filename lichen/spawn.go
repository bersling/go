package main

import "math/rand"

func Spawn(spawnProbability float64, lattice *[]int, interactionMatrix *[][]bool, eatingProbability float64) {
	latticeSize := len(*lattice)
	if spawnProbability > rand.Float64() {
		newSpeciesNumber := len(*interactionMatrix)
		randomCoordinate := rand.Intn(latticeSize - 1)
		(*lattice)[randomCoordinate] = newSpeciesNumber
		AddSpeciesToInteractionMatrix(interactionMatrix, eatingProbability)
		// TODO: Update active links
	}
}

func AddSpeciesToInteractionMatrix(interactionMatrix *[][]bool, eatingProbability float64) {
	var newSpeciesEats []bool
	var newSpeciesNumber = len(*interactionMatrix)
	for i := 0; i < newSpeciesNumber; i++ {
		newSpeciesEats = append(newSpeciesEats, eatingProbability > rand.Float64())
		newSpeciesGetsEatenByI := eatingProbability > rand.Float64()
		(*interactionMatrix)[i] = append((*interactionMatrix)[i], newSpeciesGetsEatenByI)
	}
	newSpeciesEats = append(newSpeciesEats, false) // cannot eat itself
	*interactionMatrix = append(*interactionMatrix, newSpeciesEats)
}
