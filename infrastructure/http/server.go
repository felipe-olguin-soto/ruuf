package http

import (
	"log"
	"net/http"

	"ruuf/internal/usecase"

	"github.com/gorilla/mux"
)

// StartServer starts the HTTP server.
func StartServer() error {
	router := mux.NewRouter()
	uc := getNewSolarPanelUsecase()
	handler := NewSolarPanelHandler(uc)
	router.HandleFunc("/calculate", handler.ProcessCalculatePanelsHandler).Methods("POST")

	log.Println("Server running on port 8080")
	return http.ListenAndServe(":8080", router)
}

// getNewSolarPanelUsecase returns a new instance of the SolarPanelUsecase.
func getNewSolarPanelUsecase() usecase.SolarPanelUsecase {
	return usecase.NewSolarPanelUsecase()
}
