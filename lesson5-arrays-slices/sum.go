package main

func Sum(numArray []int) int {
	sum := 0

	for _, elem := range numArray {
		sum += elem
	}
	return sum
}

func SumAll(slicesToSum ...[]int) (result []int) {
	for _, slice := range slicesToSum {
		result = append(result, Sum(slice))
	}
	return result
}

func SumAllTails(slicesToSum ...[]int) (result []int) {
	for _, slice := range slicesToSum {
		if len(slice) > 0 {
			sliceWithoutFirst := slice[1:]
			result = append(result, Sum(sliceWithoutFirst))
		} else {
			result = append(result, 0)
		}
	}
	return result
}
