package aoc

func Map[A, B any](as []A, fn func(A) B) []B {
	out := []B{}
	for _, a := range as {
		out = append(out, fn(a))
	}
	return out
}

func Filter[A any](as []A, fn func(A) bool) []A {
	out := []A{}
	for _, a := range as {
		if fn(a) {
			out = append(out, a)
		}
	}
	return out
}

func Reduce[A, B any](as []A, fn func(B, A) B) B {
	var out B
	for _, a := range as {
		out = fn(out, a)
	}
	return out
}
