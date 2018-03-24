package main

import (
	"fmt"
	"testing"
)

func TestBattle0Wins(t *testing.T) {
	lattice := []int{0, 1, 1, 1}
	interactionMatrix := [][]bool{[]bool{false, false}, []bool{true, false}}
	activeLink := [2]int{0, 1}
	activeLinks := [][2]int{activeLink}
	Battle(&lattice, interactionMatrix, activeLinks)
	if lattice[0] != 1 {
		t.Errorf("Expected first cell to be 1")
	}
}

func TestBattle1Wins(t *testing.T) {
	lattice := []int{0, 1, 1, 1}
	interactionMatrix := [][]bool{[]bool{false, true}, []bool{false, false}}
	activeLink := [2]int{0, 1}
	activeLinks := [][2]int{activeLink}
	Battle(&lattice, interactionMatrix, activeLinks)
	if lattice[1] != 0 {
		t.Errorf("Expected second cell to be 0")
	}
}

func TestUncertainWin(t *testing.T) {
	aWonNTimes := 0
	for i := 0; i < 1000; i++ {
		lattice, interactionMatrix, activeLinks := setupInitialConditions()
		Battle(&lattice, interactionMatrix, activeLinks)
		if lattice[0] == 1 {
			aWonNTimes += 1
		}
	}
	fmt.Println(aWonNTimes, "a won") // TODO: always 508, do I need better RNG?
	if aWonNTimes < 10 || aWonNTimes >= 990 {
		t.Errorf("Doesnt seem to be balanced")
	}
}

func setupInitialConditions() ([]int, [][]bool, [][2]int) {
	lattice := []int{0, 1, 1, 1}
	interactionMatrix := [][]bool{[]bool{false, true}, []bool{true, false}}
	activeLink := [2]int{0, 1}
	activeLinks := [][2]int{activeLink}
	return lattice, interactionMatrix, activeLinks
}

func TestEat(t *testing.T) {
	lattice := []int{0, 1, 1, 1}
	Eat(1, 0, &lattice)
	if lattice[0] != 1 {
		t.Errorf("Expected first cell to be 1")
	}
}
