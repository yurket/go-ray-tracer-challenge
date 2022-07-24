package main

type Color struct {
	r float64
	g float64
	b float64
}

func (a Color) Equal(b Color) bool {
	return equal_fp(a.r, b.r) && equal_fp(a.g, b.g) && equal_fp(a.b, b.b)
}

func (a Color) Add(b Color) Color {
	return Color{a.r + b.r, a.g + b.g, a.b + b.b}
}

func (a Color) Sub(b Color) Color {
	return Color{a.r - b.r, a.g - b.g, a.b - b.b}
}

func (a Color) MultScalar(n float64) Color {
	return Color{a.r * n, a.g * n, a.b * n}
}

func (a Color) MultHadamar(b Color) Color {
	return Color{a.r * b.r, a.g * b.g, a.b * b.b}
}
