package main

import (
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
