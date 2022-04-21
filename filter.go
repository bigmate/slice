package slice

func Filter[T any](arr []T, lambda func(T) bool) []T {
	cp := make([]T, 0)
	for _, t := range arr {
		if lambda(t) {
			cp = append(cp, t)
		}
	}
	return cp
}
