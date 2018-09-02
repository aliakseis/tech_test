package api

import (
	"log"
	"tech_test/models"
	"tech_test/state"
	"tech_test/util"
)

// A new tech_test in POST body
func SalePlacement(salePlacementRequest models.SalePlacementRequest) (models.Status, error) {

	log.Print(util.JsonToString(salePlacementRequest))

	status := models.Status{Status: "success"}

	theState := state.GetInstance()
	theState.Mux.Lock()
	defer theState.Mux.Unlock()
	defer state.DoHouseKeeping(theState)

	details := theState.Details[salePlacementRequest.ProductType]
	details = append(details, salePlacementRequest.Price)
	theState.Details[salePlacementRequest.ProductType] = details

	return status, nil
}
