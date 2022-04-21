package slice

func Map[A, B any](arr []A, lambda func(A) B) []B {
	refArr := make([]B, len(arr))

	for i, t := range arr {
		refArr[i] = lambda(t)
	}

	return refArr
}
