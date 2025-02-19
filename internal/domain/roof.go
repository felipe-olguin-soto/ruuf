package domain

import (
	"ruuf/pkg/utils"
	"strings"
)

// Roof represents a roof with a type, width, and height.
type Roof struct {
	Type    ShapeType `json:"type"`
	Size    []Size    `json:"size"`
	Overlap Size      `json:"overlap"`
}

// ShapeType represents a type of shape as a string.
type ShapeType string

const (
	Rectangle ShapeType = "rectangle"
	Triangle  ShapeType = "triangle"
	Overlap   ShapeType = "overlap"
)

// Area returns the area of the shape.
func (s *Roof) Area() float64 {
	switch s.Type.Normalize() {
	case Rectangle:
		return utils.CalculateSquareArea(s.Size[0].Width, s.Size[0].Height)
	case Triangle:
		return utils.CalculateTriangleArea(s.Size[0].Width, s.Size[0].Height)
	case Overlap:
		return utils.CalculateSquareArea(s.Size[0].Width, s.Size[0].Height) +
			utils.CalculateSquareArea(s.Size[1].Width, s.Size[1].Height) - s.OverlapArea()
	default:
		return 0
	}
}

func (s *Roof) OverlapArea() float64 {
	return utils.CalculateSquareArea(s.Size[0].Width, s.Size[0].Height) - utils.CalculateSquareArea(s.Overlap.Width, s.Overlap.Height)
}

// Normalize converts the ShapeType to lowercase.
func (s ShapeType) Normalize() ShapeType {
	return ShapeType(strings.ToLower(string(s)))
}
