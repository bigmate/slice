package slice

func Map[A, B any](arr []A, lambda func(A) B) []B {
	refArr := make([]B, len(arr))

	for i := range arr {
		refArr[i] = lambda(arr[i])
	}

	return refArr
}

func MapPtr[A, B any](arr []A, lambda func(*A) B) []B {
	refArr := make([]B, len(arr))

	for i := range arr {
		refArr[i] = lambda(&arr[i])
	}

	return refArr
}

func MapWithError[A, B any](arr []A, lambda func(A) (B, error)) ([]B, error) {
	refArr := make([]B, len(arr))

	for i, t := range arr {
		res, err := lambda(t)
		if err != nil {
			return nil, err
		}

		refArr[i] = res
	}

	return refArr, nil
}

func Apply[T any](arr []T, lambda func(T) T) {
	for i := range arr {
		arr[i] = lambda(arr[i])
	}
}

func ApplyPtr[T any](arr []T, lambda func(*T)) {
	for i := range arr {
		lambda(&arr[i])
	}
}
