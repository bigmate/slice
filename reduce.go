package slice

func Reduce[A, B any](acc A, arr []B, lambda func(a A, b B) A) A {
	for i := range arr {
		acc = lambda(acc, arr[i])
	}

	return acc
}

func ReducePtr[A, B any](acc A, arr []B, lambda func(a A, b *B) A) A {
	for i := range arr {
		acc = lambda(acc, &arr[i])
	}

	return acc
}
