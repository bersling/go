package main

import (
	"strconv"
	"testing"
)

// how a 3x3x3 cube looks like

// 00 01 02
// 03 04 05
// 06 07 08

// 09 10 11
// 12 13 14
// 15 16 17

// 18 19 20
// 21 22 23
// 24 25 26

func TestCellIsLowerBoundaryDim1(t *testing.T) {

	x := true
	f := false
	latticeSize := 3
	dimension := 1

	lowerBoundaries := []bool{
		x, f, f,
		x, f, f,
		x, f, f,

		x, f, f,
		x, f, f,
		x, f, f,

		x, f, f,
		x, f, f,
		x, f, f}

	for i := 0; i < Pow(latticeSize, 3); i++ {
		if CellIsLowerBoundary(i, dimension, latticeSize) != lowerBoundaries[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i))
		}
	}
}

func TestCellIsLowerBoundaryDim2(t *testing.T) {

	x := true
	f := false
	latticeSize := 3
	dimension := 2

	lowerBoundaries := []bool{
		x, x, x,
		f, f, f,
		f, f, f,

		x, x, x,
		f, f, f,
		f, f, f,

		x, x, x,
		f, f, f,
		f, f, f}

	for i := 0; i < Pow(latticeSize, 3); i++ {
		if CellIsLowerBoundary(i, dimension, latticeSize) != lowerBoundaries[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i))
		}
	}
}

func TestCellIsLowerBoundaryDim3(t *testing.T) {

	x := true
	f := false
	latticeSize := 3
	dimension := 3

	lowerBoundaries := []bool{
		x, x, x,
		x, x, x,
		x, x, x,

		f, f, f,
		f, f, f,
		f, f, f,

		f, f, f,
		f, f, f,
		f, f, f}

	for i := 0; i < Pow(latticeSize, 3); i++ {
		if CellIsLowerBoundary(i, dimension, latticeSize) != lowerBoundaries[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i))
		}
	}
}

// Upper Boundaries
func TestCellIsUpperBoundaryDim1(t *testing.T) {

	x := true
	f := false
	latticeSize := 3
	dimension := 1

	upperBoundaries := []bool{
		f, f, x,
		f, f, x,
		f, f, x,

		f, f, x,
		f, f, x,
		f, f, x,

		f, f, x,
		f, f, x,
		f, f, x}

	for i := 0; i < Pow(latticeSize, 3); i++ {
		if CellIsUpperBoundary(i, dimension, latticeSize) != upperBoundaries[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i))
		}
	}
}

func TestCellIsUpperBoundaryDim2(t *testing.T) {

	x := true
	f := false
	latticeSize := 3
	dimension := 2

	upperBoundaries := []bool{
		f, f, f,
		f, f, f,
		x, x, x,

		f, f, f,
		f, f, f,
		x, x, x,

		f, f, f,
		f, f, f,
		x, x, x}

	for i := 0; i < Pow(latticeSize, 3); i++ {
		if CellIsUpperBoundary(i, dimension, latticeSize) != upperBoundaries[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i))
		}
	}
}

func TestCellIsUpperBoundaryDim3(t *testing.T) {

	x := true
	f := false
	latticeSize := 3
	dimension := 3

	upperBoundaries := []bool{
		f, f, f,
		f, f, f,
		f, f, f,

		f, f, f,
		f, f, f,
		f, f, f,

		x, x, x,
		x, x, x,
		x, x, x}

	for i := 0; i < Pow(latticeSize, 3); i++ {
		if CellIsUpperBoundary(i, dimension, latticeSize) != upperBoundaries[i] {
			t.Errorf("ERROR at cell " + strconv.Itoa(i))
		}
	}
}
