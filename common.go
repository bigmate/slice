package slice

import "golang.org/x/exp/constraints"

func amp[T constraints.Ordered](assign func(e, m T) bool, els ...T) T {
	if len(els) < 1 {
		panic("invalid arguments number")
	}

	m := els[0]

	for _, e := range els[1:] {
		if assign(e, m) {
			m = e
		}
	}

	return m
}

func Min[T constraints.Ordered](els ...T) T {
	return amp(func(e, m T) bool {
		return e < m
	}, els...)
}

func Max[T constraints.Ordered](els ...T) T {
	return amp(func(e, m T) bool {
		return e > m
	}, els...)
}

func Sum[T constraints.Ordered](els ...T) T {
	s := *new(T)

	for _, el := range els {
		s += el
	}

	return s
}

func Any[T comparable](target T, els ...T) bool {
	for i := range els {
		el := els[i]

		if el == target {
			return true
		}
	}

	return false
}

func AnyFunc[T any](lambda func(el T) bool, els ...T) bool {
	for _, el := range els {
		if lambda(el) {
			return true
		}
	}

	return false
}

func All[T comparable](target T, els ...T) bool {
	for i := range els {
		el := els[i]

		if el != target {
			return false
		}
	}

	return true
}

func AllFunc[T any](lambda func(el T) bool, els ...T) bool {
	for _, el := range els {
		if !lambda(el) {
			return false
		}
	}

	return true
}

func Unique[T comparable](els ...T) []T {
	m := make(map[T]struct{}, len(els))

	for _, el := range els {
		m[el] = struct{}{}
	}

	u := make([]T, 0, len(m))

	for t := range m {
		u = append(u, t)
	}

	return u
}
