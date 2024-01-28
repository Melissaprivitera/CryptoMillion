package main

func packNumbers(numbers [5]uint8) uint64 {
	var result uint64 = 0

	for	i := 4; i >= 0; i-- {
		result = result << 8
		result = result | uint64(numbers[i])
	}
	return result
}

func unpackNumbers(number uint64) [5]uint8 {
	var result [5]uint8

	for i := 0; i < 5; i++ {
		result[i] = uint8(number & 0xFF)
		number = number >> 8
	}
	return result
}
