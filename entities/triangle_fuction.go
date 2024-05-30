package entities

// TriangleMembership calculates the membership value for a given x
func TriangleMembership(x, a, b, c float64) float64 {
	if x <= a || x >= c {
		return 0.0
	} else if x == b {
		return 1.0
	} else if x > a && x < b {
		return (x - a) / (b - a)
	} else { // x > b && x < c
		return (c - x) / (c - b)
	}
}
