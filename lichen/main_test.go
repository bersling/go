package main

import (
	"fmt"
	"strconv"
	"testing"
)

// For testing, we mainly use a 3x3x3 example:
// 0, 1, 2,
// 3, 4, 5,
// 6, 7, 8,

// 9,  10, 11,
// 12, 13, 14,
// 15, 16, 17,

// 18, 19, 20,
// 21, 22, 23,
// 24, 25, 26

func TestCellIsUpperNeighborDim1(t *testing.T) {
	latticeSize := 3
	dimension := 1
	dimensions := 3

	upperNeighbors := []int{
		1, 2, 0,
		4, 5, 3,
		7, 8, 6,

		10, 11, 9,
		13, 14, 12,
		16, 17, 15,

		19, 20, 18,
		22, 23, 21,
		25, 26, 24}

	for i := 0; i < Pow(latticeSize, dimensions); i++ {
		if UpperNeighbor(i, dimension, latticeSize) != upperNeighbors[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i) + ". Expected: " + strconv.Itoa(upperNeighbors[i]) + " but found " + strconv.Itoa(UpperNeighbor(i, dimension, latticeSize)))
		}
	}
}

func TestCellIsUpperNeighborDim2(t *testing.T) {
	latticeSize := 3
	dimension := 2
	dimensions := 3

	upperNeighbors := []int{
		3, 4, 5,
		6, 7, 8,
		0, 1, 2,

		12, 13, 14,
		15, 16, 17,
		9, 10, 11,

		21, 22, 23,
		24, 25, 26,
		18, 19, 20}

	for i := 0; i < Pow(latticeSize, dimensions); i++ {
		if UpperNeighbor(i, dimension, latticeSize) != upperNeighbors[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i) + ". Expected: " + strconv.Itoa(upperNeighbors[i]) + " but found " + strconv.Itoa(UpperNeighbor(i, dimension, latticeSize)))
		}
	}
}

func TestCellIsUpperNeighborDim3(t *testing.T) {
	latticeSize := 3
	dimension := 3
	dimensions := 3

	upperNeighbors := []int{
		9, 10, 11,
		12, 13, 14,
		15, 16, 17,

		18, 19, 20,
		21, 22, 23,
		24, 25, 26,

		0, 1, 2,
		3, 4, 5,
		6, 7, 8}

	for i := 0; i < Pow(latticeSize, dimensions); i++ {
		if UpperNeighbor(i, dimension, latticeSize) != upperNeighbors[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i) + ". Expected: " + strconv.Itoa(upperNeighbors[i]) + " but found " + strconv.Itoa(UpperNeighbor(i, dimension, latticeSize)))
		}
	}
}

func TestImmutability(t *testing.T) {
	bla := []bool{false}
	myFun(bla)
	fmt.Println(bla[0])
}

func myFun(bla []bool) {
	bla = append(bla, true) // Does this create a function scoped variable bla?
	bla[0] = true
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

func TestSpawn(t *testing.T) {
	interactionMatrix := [][]bool{[]bool{false, true}, []bool{false, false}}
	// lattice:
	// 1 1 1
	// 0 0 1
	// 1 1 1 (2 species, species 1 is dominating species 0)
	lattice := []int{1, 1, 1, 0, 0, 1, 1, 1, 1}
	spawnProbability := 1.0
	eatingProbability := 1.0
	// test immutability
	tmp := make([]int, len(lattice))
	copy(tmp, lattice)
	Spawn(spawnProbability, lattice, interactionMatrix, eatingProbability)
	if !sliceEquals(tmp, lattice) {
		t.Errorf("Expected no mutation")
	}
	// test no spawn
	spawnProbability = 0.0
	updatedLattice, updatedInteractionMatrix := Spawn(spawnProbability, lattice, interactionMatrix, eatingProbability)
	if updatedLattice != nil || updatedInteractionMatrix != nil {
		t.Errorf("Expected no change")
	}
	// test spawn
	spawnProbability = 1.0
	updatedLattice, updatedInteractionMatrix = Spawn(spawnProbability, lattice, interactionMatrix, eatingProbability)
	if sliceEquals(lattice, updatedLattice) {
		t.Errorf("Expected update")
	}
	containsTwo := false
	for _, v := range updatedLattice {
		if v == 2 {
			containsTwo = true
		}
	}
	if containsTwo == false {
		t.Errorf("Expected lattice to contain '2'")
	}
}

func sliceEquals(sliceA []int, sliceB []int) bool {
	isEqual := true
	fmt.Println("===")
	fmt.Println("len", len(sliceA), len(sliceB))
	if len(sliceA) != len(sliceB) {
		isEqual = false
	} else {
		for i, v := range sliceA {
			fmt.Println("val", v, sliceB[i])
			if sliceB[i] != v {
				isEqual = false
			}
		}
	}
	return isEqual
}

func TestPow(t *testing.T) {
	if Pow(2, 3) != 8 || Pow(5, 2) != 25 || Pow(3, 3) != 27 || Pow(2, 5) != 32 {
		t.Error("Incorrect")
	}
}
