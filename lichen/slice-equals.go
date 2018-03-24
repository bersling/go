package main

func SliceEquals(sliceA []int, sliceB []int) bool {
	isEqual := true
	if len(sliceA) != len(sliceB) {
		isEqual = false
	} else {
		for i, v := range sliceA {
			if sliceB[i] != v {
				isEqual = false
			}
		}
	}
	return isEqual
}
