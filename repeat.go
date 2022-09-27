package slice

func Repeat[T any](el T, count int) []T {
	res := make([]T, 0, count)

	for i := 0; i < count; i++ {
		res = append(res, el)
	}

	return res
}
