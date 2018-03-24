package main

// e.g. in a 10x10 lattice:
// 00 10 20 ... 90 are lower boundary cells of dim 1
// 00 01 02 ... 09 of dim 2
func CellIsLowerBoundary(cell int, dimension int, latticeSize int) bool {
	dimensionAdjustment := Pow(latticeSize, dimension-1)
	return cell/dimensionAdjustment%latticeSize == 0
}

// e.g. in a 10x10 lattice:
// 09 19 29 ... 99 are upper boundary cells of dim 1
// 90 91 92 ... 99 of dim 2
func CellIsUpperBoundary(cell int, dimension int, latticeSize int) bool {
	return CellIsLowerBoundary(getCellOnOppositeSite(cell, dimension, latticeSize), dimension, latticeSize)
}

func getCellOnOppositeSite(cell int, dimension int, latticeSize int) int {
	dimensionAdjustment := Pow(latticeSize, dimension-1)
	distanceBetweenOppositeSides := (latticeSize - 1) * dimensionAdjustment
	totalLatticeSize := Pow(latticeSize, dimension)
	return ((cell - distanceBetweenOppositeSides) + totalLatticeSize) % totalLatticeSize
}
