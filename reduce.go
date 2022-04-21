package slice

func Reduce[T any](arr []T, lambda func(a, b T) T) T {
	if len(arr) == 0 {
		panic("expected one or more args")
	}

	acc := arr[0]

	for _, t := range arr[1:] {
		acc = lambda(acc, t)
	}

	return acc
}
