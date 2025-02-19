package domain

import (
	"ruuf/pkg/utils"
)

// SolarPanel represents a panel with a width and height.
type SolarPanel struct {
	Size Size `json:"size"`
}

// Area returns the area of the solar panel.
func (s *SolarPanel) Area() float64 {
	return utils.CalculateSquareArea(s.Size.Width, s.Size.Height)
}
