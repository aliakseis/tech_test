package api

import (
	"errors"
	"log"
	"tech_test/models"
	"tech_test/state"
	"tech_test/util"
)

// A new tech_test in POST body
func SalesAdjustment(salesAdjustmentRequest models.SalesAdjustmentRequest) (models.Status, error) {

	log.Print(util.JsonToString(salesAdjustmentRequest))

	status := models.Status{Status: "success"}

	if salesAdjustmentRequest.Divider == 0 {
		status.Status = "error"
		status.StatusMessage = "Zero divider"
		return status, errors.New(status.StatusMessage)
	}

	theState := state.GetInstance()
	theState.Mux.Lock()
	defer theState.Mux.Unlock()
	defer state.DoHouseKeeping(theState)

	details := theState.Details[salesAdjustmentRequest.ProductType]

	for i := 0; i < len(details); i++ {
		details[i] = details[i]*salesAdjustmentRequest.Multiplier/salesAdjustmentRequest.Divider + salesAdjustmentRequest.Addend
	}

	theState.Adjustments[salesAdjustmentRequest.ProductType] = append(
		theState.Adjustments[salesAdjustmentRequest.ProductType],
		state.Adjustment{
			Multiplier: salesAdjustmentRequest.Multiplier,
			Divider:    salesAdjustmentRequest.Divider,
			Addend:     salesAdjustmentRequest.Addend,
		})

	return status, nil
}
