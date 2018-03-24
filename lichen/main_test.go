package main

import (
	"testing"
)

func TestCanEat(t *testing.T) {
	interactionMatrix := [][]bool{[]bool{false, true}, []bool{false, false}}
	if CanEat(0, 0, interactionMatrix) || !CanEat(0, 1, interactionMatrix) ||
		CanEat(1, 0, interactionMatrix) || CanEat(1, 1, interactionMatrix) {
		t.Errorf("There's a problem with CanEat")
	}
}

func TestPow(t *testing.T) {
	if Pow(2, 3) != 8 || Pow(5, 2) != 25 || Pow(3, 3) != 27 || Pow(2, 5) != 32 {
		t.Error("Incorrect")
	}
}
