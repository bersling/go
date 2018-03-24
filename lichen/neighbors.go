package main

func GetUpperNeighbors(cell int, dimensions int, latticeSize int) []int {
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
