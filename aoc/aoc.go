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

type Point struct {
	X, Y int
}

type Grid []string

func (g Grid) InBounds(p Point) bool {
	return p.Y >= 0 && p.Y < len(g) && p.X >= 0 && p.X < len(g[0])
}

func (g Grid) Get(p Point) string {
	if !g.InBounds(p) {
		return ""
	}
	return string(g[p.Y][p.X])
}

func (g Grid) Neighbours(p Point) []Point {
	ns := []Point{}
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			n := Point{p.X + i, p.Y + j}
			if g.InBounds(n) && n != p {
				ns = append(ns, n)
			}
		}
	}
	return ns
}

func (g Grid) Points(filter func(Point) bool) []Point {
	ps := []Point{}
	for y := 0; y < len(g); y++ {
		for x := 0; x < len(g[0]); x++ {
			p := Point{x, y}
			if filter(p) {
				ps = append(ps, Point{x, y})
			}
		}
	}
	return ps
}
