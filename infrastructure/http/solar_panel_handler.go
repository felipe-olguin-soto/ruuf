package http

import (
	"encoding/json"
	"net/http"
	"ruuf/internal/domain"
	"ruuf/internal/usecase"
)

type SolarPanelHandler struct {
	uc usecase.SolarPanelUsecase
}

func NewSolarPanelHandler(uc usecase.SolarPanelUsecase) *SolarPanelHandler {
	return &SolarPanelHandler{uc: uc}
}

// ProcessCalculatePanelsHandler handles the request to calculate the number of solar panels that can fit on a given shape.
func (h *SolarPanelHandler) ProcessCalculatePanelsHandler(w http.ResponseWriter, r *http.Request) {
	var data domain.SolarCollection

	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	res, err := h.uc.Calculate(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
