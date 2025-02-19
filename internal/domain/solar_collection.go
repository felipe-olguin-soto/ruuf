package domain

// SolarCollection represents a collection of solar panels on a roof.
type SolarCollection struct {
	Roof       Roof       `json:"roof"`
	SolarPanel SolarPanel `json:"solar_panel"`
}
