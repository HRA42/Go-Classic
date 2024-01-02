package unknown

func Unknown(n uint8) uint8 {
	if n <= 1 {
		return n
	}
	return Unknown(n-1) + Unknown(n-2)
}

func UnknownSequence(n uint8) []uint8 {
	var sequence []uint8
	for i := uint8(0); i <= n; i++ {
		sequence = append(sequence, Unknown(i))
	}
	return sequence
}
