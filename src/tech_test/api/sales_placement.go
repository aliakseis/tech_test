package api

import (
	"log"
	"tech_test/models"
	"tech_test/state"
	"tech_test/util"
)

// A new tech_test in POST body
func SalesPlacement(salesPlacementRequest models.SalesPlacementRequest) (models.Status, error) {

	log.Print(util.JsonToString(salesPlacementRequest))

	status := models.Status{Status: "success"}

	theState := state.GetInstance()
	theState.Mux.Lock()
	defer theState.Mux.Unlock()
	defer state.DoHouseKeeping(theState)

	details := theState.Details[salesPlacementRequest.ProductType]

	for i := int64(0); i < salesPlacementRequest.Quantity; i++ {
		details = append(details, salesPlacementRequest.Price)
	}

	theState.Details[salesPlacementRequest.ProductType] = details

	return status, nil
}
