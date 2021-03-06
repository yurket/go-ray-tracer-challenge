package main

import "math"

// w represents whether it's a vector or a point
type tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func Tuple() tuple {
	return tuple{0, 0, 0, 0}
}

func Point(x float64, y float64, z float64) tuple {
	return tuple{x, y, z, 1.0}
}

func Vector(x float64, y float64, z float64) tuple {
	return tuple{x, y, z, 0.0}
}

func IsPoint(t tuple) bool {
	return equal_fp(t.w, 1.0)
}

func IsVector(t tuple) bool {
	return equal_fp(t.w, 0.0)
}

const EPSILON = 1e-5

func equal_fp(a float64, b float64) bool {
	diff := a - b
	return math.Abs(diff) < EPSILON
}

func equal(a tuple, b tuple) bool {
	return equal_fp(a.x, b.x) && (equal_fp(a.y, b.y) && equal_fp(a.z, b.z) && equal_fp(a.w, b.w))
}

func Add(a tuple, b tuple) tuple {
	if IsPoint(a) && IsPoint(b) {
		panic("Can't add two points!")
	}

	return tuple{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func Sub(a tuple, b tuple) tuple {
	w := a.w - b.w
	if w < 0 {
		panic("Can't subtract point from vector!")
	}
	return tuple{a.x - b.x, a.y - b.y, a.z - b.z, w}
}

func Negate(t tuple) tuple {
	return tuple{-t.x, -t.y, -t.z, -t.w}
}

func Mul(t tuple, c float64) tuple {
	return tuple{t.x * c, t.y * c, t.z * c, t.w * c}
}

func Div(t tuple, c float64) tuple {
	if equal_fp(c, 0) {
		panic("Can't divide by zero!")
	}
	return tuple{t.x / c, t.y / c, t.z / c, t.w / c}
}

func Magnitude(t tuple) float64 {
	if !IsVector(t) {
		panic("Magnitude isn't applicable to non-Vectors!")
	}

	// the book proposes also to add up w*w term (error?)
	return math.Sqrt(t.x*t.x + t.y*t.y + t.z*t.z)
}

func Normalize(t tuple) tuple {
	m := Magnitude(t)
	return tuple{t.x / m, t.y / m, t.z / m, t.w / m}
}

func Dot(a tuple, b tuple) float64 {
	if !IsVector(a) || !IsVector(b) {
		panic("Dot product isn't applicable to non-Vectors!")
	}

	return a.x*b.x + a.y*b.y + a.z*b.z
}

func Cross(a tuple, b tuple) tuple {
	if !IsVector(a) || !IsVector(b) {
		panic("Cross product isn't applicable to non-Vectors!")
	}

	return Vector(a.y*b.z-a.z*b.y, a.z*b.x-a.x*b.z, a.x*b.y-a.y*b.x)
}
