package main

import "testing"

func TestCellIsLowerBoundaryDim1(t *testing.T) {
	// Testing for lattice 3x3
	t0 := CellIsLowerBoundary(0, 1, 3)
	if !t0 {
		t.Errorf("Incorrect")
	}
	t1 := CellIsLowerBoundary(1, 1, 3)
	if t1 {
		t.Errorf("Incorrect")
	}
	t2 := CellIsLowerBoundary(2, 1, 3)
	if t2 {
		t.Errorf("Incorrect")
	}
	t3 := CellIsLowerBoundary(3, 1, 3)
	if !t3 {
		t.Errorf("Incorrect")
	}
}

func TestCellIsLowerBoundaryDim2(t *testing.T) {
	// Testing for lattice 3x3
	t0 := CellIsLowerBoundary(0, 2, 3)
	if !t0 {
		t.Errorf("Incorrect")
	}
	t1 := CellIsLowerBoundary(1, 2, 3)
	if !t1 {
		t.Errorf("Incorrect")
	}
	t2 := CellIsLowerBoundary(2, 2, 3)
	if !t2 {
		t.Errorf("Incorrect")
	}
	t3 := CellIsLowerBoundary(3, 2, 3)
	if t3 {
		t.Errorf("Incorrect")
	}
	t4 := CellIsLowerBoundary(4, 2, 3)
	if t4 {
		t.Errorf("Incorrect")
	}
}

func TestCellIsLowerBoundaryDim3(t *testing.T) {
	// Testing for lattice 3x3
	t0 := CellIsLowerBoundary(0, 3, 3)
	if !t0 {
		t.Errorf("Incorrect")
	}
	t1 := CellIsLowerBoundary(6, 3, 3)
	if !t1 {
		t.Errorf("Incorrect")
	}
	t2 := CellIsLowerBoundary(8, 3, 3)
	if !t2 {
		t.Errorf("Incorrect")
	}
	t3 := CellIsLowerBoundary(9, 3, 3)
	if t3 {
		t.Errorf("Incorrect")
	}
	t4 := CellIsLowerBoundary(16, 3, 3)
	if t4 {
		t.Errorf("Incorrect")
	}
}
