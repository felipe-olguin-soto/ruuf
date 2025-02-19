package utils

// CalculateSquareArea calculates the area of a square given its width and height.
func CalculateSquareArea(width, height float64) float64 {
	return width * height
}

// CalculateTriangleArea calculates the area of a triangle given its width and height.
func CalculateTriangleArea(width, height float64) float64 {
	return (width * height) / 2
}

// Min returns the minimum of two integers.
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
