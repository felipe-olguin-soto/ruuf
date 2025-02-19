package usecase

import (
	"errors"
	"fmt"
	"math"
	"ruuf/internal/domain"
	"ruuf/pkg/utils"
)

type solarPanelUseCase struct{}

// Calculate calculates the number of solar panels that can fit on a given shape.
type SolarPanelUsecase interface {
	Calculate(data domain.SolarCollection) (*SolarPanelResponse, error)
}

// SolarPanelResponse represents the response of the Calculate method.
type SolarPanelResponse struct {
	PanelsCount PanelsCount `json:"panels_count"`
	Metadata    Metadata    `json:"metadata"`
}

// Metadata represents the metadata of the Calculate method.
type Metadata struct {
	SolarPanel domain.SolarPanel `json:"solar_panel"`
	Roof       domain.Roof       `json:"roof"`
}

type PanelsCount struct {
	Standard int `json:"standard"`
	Rotated  int `json:"rotated"`
}

func (uc *solarPanelUseCase) Calculate(data domain.SolarCollection) (*SolarPanelResponse, error) {
	// Validate the dimensions of the solar panel and shape.
	if err := validateDimensions(data); err != nil {
		return nil, err
	}

	switch data.Roof.Type.Normalize() {
	case domain.Rectangle:
		return calculatePanelsForRectangle(data), nil
	case domain.Triangle:
		return calculatePanelsForTriangle(data), nil
	case domain.Overlap:
		return calculatePanelsForOverlappingRectangles(data), nil
	default:
		return nil, errors.New("invalid roof type")
	}

}

// calculateRectangle calculates the number of solar panels that can fit on a rectangle.
func calculatePanelsForRectangle(data domain.SolarCollection) *SolarPanelResponse {
	roofDimension := data.Roof.Size[0]

	countStandard := (int(roofDimension.Width/data.SolarPanel.Size.Width) * int(roofDimension.Height/data.SolarPanel.Size.Height))
	countRotated := (int(roofDimension.Width/data.SolarPanel.Size.Height) * int(roofDimension.Height/data.SolarPanel.Size.Width))

	// Create the response.
	res := SolarPanelResponse{
		PanelsCount: PanelsCount{
			Standard: countStandard,
			Rotated:  countRotated,
		},
		Metadata: Metadata{
			SolarPanel: data.SolarPanel,
			Roof:       data.Roof,
		},
	}
	return &res
}

// calculateTriangle calculates the number of solar panels that can fit on a triangle.
func calculatePanelsForTriangle(data domain.SolarCollection) *SolarPanelResponse {
	// Calculate the number of solar panels that can fit on a triangle based on the area.
	approximatePanels := int(math.Floor(data.Roof.Area() / data.SolarPanel.Area()))

	roofDimension := data.Roof.Size[0]
	totalPanelsStandard := 0
	totalPanelsRotated := 0
	rows := int(roofDimension.Height / data.SolarPanel.Size.Height)

	for i := 1; i <= rows; i++ {
		// Calculate the available width for each row.
		availableWidthStandard := roofDimension.Width - (float64(i) * (data.SolarPanel.Size.Width / float64(rows)))
		if availableWidthStandard > 0 {
			panelsInRowStandard := int(math.Floor(availableWidthStandard / data.SolarPanel.Size.Width))
			totalPanelsStandard += panelsInRowStandard
		}

		// Calculate the available width for each row when the solar panel is rotated.
		availableWidthRotated := roofDimension.Width - (float64(i) * (data.SolarPanel.Size.Height / float64(rows)))
		if availableWidthRotated > 0 {
			panelsInRowRotated := int(math.Floor(availableWidthRotated / data.SolarPanel.Size.Height))
			totalPanelsRotated += panelsInRowRotated
		}
	}

	resStandard := utils.Min(approximatePanels, totalPanelsStandard)
	resRotated := utils.Min(approximatePanels, totalPanelsRotated)

	return &SolarPanelResponse{
		PanelsCount: PanelsCount{
			Standard: resStandard,
			Rotated:  resRotated,
		},
		Metadata: Metadata{
			SolarPanel: data.SolarPanel,
			Roof:       data.Roof,
		},
	}
}

func calculatePanelsForOverlappingRectangles(data domain.SolarCollection) *SolarPanelResponse {
	// Calculate the number of panels in standard orientation
	countStandard := 0
	for _, dim := range data.Roof.Size {
		countStandard += int(dim.Width/data.SolarPanel.Size.Width) * int(dim.Height/data.SolarPanel.Size.Height)
	}

	// Calculate the number of panels in rotated orientation
	countRotated := 0
	for _, dim := range data.Roof.Size {
		countRotated += int(dim.Width/data.SolarPanel.Size.Height) * int(dim.Height/data.SolarPanel.Size.Width)
	}

	return &SolarPanelResponse{
		PanelsCount: PanelsCount{
			Standard: countStandard,
			Rotated:  countRotated,
		},
		Metadata: Metadata{
			SolarPanel: data.SolarPanel,
			Roof:       data.Roof,
		},
	}
}

// validateDimensions validates the dimensions of the solar panel and shape.
func validateDimensions(data domain.SolarCollection) error {
	// Check if the dimensions of the solar panel and shape are greater than zero.
	if data.SolarPanel.Size.Width <= 0 || data.SolarPanel.Size.Height <= 0 {
		return errors.New("the dimensions of the solar panel must be greater than zero")
	}

	// Validate the dimensions of the shape based on its type.
	switch data.Roof.Type.Normalize() {
	case domain.Triangle, domain.Rectangle:
		if len(data.Roof.Size) > 1 {
			return fmt.Errorf("the dimensions of the shape %s must be a single value", data.Roof.Type)
		}
	case domain.Overlap:
		if len(data.Roof.Size) != 2 {
			return fmt.Errorf("the dimensions of the shape %s must be a pair of values", data.Roof.Type)
		}
		if data.Roof.Overlap.Width <= 0 || data.Roof.Overlap.Height <= 0 {
			return errors.New("the dimensions of the overlap must be greater than zero")
		}
	}

	// Validate all dimensions in the slice
	for _, dim := range data.Roof.Size {
		if dim.Width <= 0 || dim.Height <= 0 {
			return errors.New("the dimensions of the roof must be greater than zero")
		}
		if data.SolarPanel.Size.Width > dim.Width || data.SolarPanel.Size.Height > dim.Height {
			return errors.New("the solar panel cannot be larger than the roof")
		}
	}

	return nil
}

// NewSolarPanelUsecase creates a new instance of the SolarPanelUsecase interface.
func NewSolarPanelUsecase() SolarPanelUsecase {
	return &solarPanelUseCase{}
}
