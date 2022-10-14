package slice

func Map[A, B any](arr []A, lambda func(A) B) []B {
	refArr := make([]B, len(arr))

	for i, t := range arr {
		refArr[i] = lambda(t)
	}

	return refArr
}

func Apply[T any](arr []T, lambda func(T) T) {
	for i := range arr {
		arr[i] = lambda(arr[i])
	}
}
