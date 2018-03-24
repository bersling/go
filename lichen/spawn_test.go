package main

import "testing"

func TestSpawn(t *testing.T) {
	interactionMatrix := [][]bool{[]bool{false, true}, []bool{false, false}}
	// lattice:
	// 1 1 1
	// 0 0 1
	// 1 1 1 (2 species, species 1 is dominating species 0)
	lattice := []int{1, 1, 1, 0, 0, 1, 1, 1, 1}
	initialLattice := make([]int, len(lattice))
	copy(initialLattice, lattice)
	spawnProbability := 1.0
	eatingProbability := 1.0
	// test no spawn
	spawnProbability = 0.0
	Spawn(spawnProbability, &lattice, &interactionMatrix, eatingProbability)
	if !SliceEquals(initialLattice, lattice) {
		t.Errorf("Expected no change")
	}

	// test spawn
	spawnProbability = 1.0
	Spawn(spawnProbability, &lattice, &interactionMatrix, eatingProbability)
	if SliceEquals(lattice, initialLattice) {
		t.Errorf("Expected update")
	}
	containsTwo := false
	for _, v := range lattice {
		if v == 2 {
			containsTwo = true
		}
	}
	if containsTwo == false {
		t.Errorf("Expected lattice to contain '2'")
	}
}

func TestAddSpeciesToInteractionMatrix(t *testing.T) {
	interactionMatrix := [][]bool{[]bool{false, true}, []bool{false, false}}
	eatingProbability := 1.0
	AddSpeciesToInteractionMatrix(&interactionMatrix, eatingProbability)
	if len(interactionMatrix) != 3 || len(interactionMatrix[2]) != 3 {
		t.Errorf("Matrix has wrong dimensions")
	} else if interactionMatrix[1][2] != true || interactionMatrix[2][1] != true || interactionMatrix[2][2] != false {
		t.Errorf("There is a wrong entry")
	}
}
