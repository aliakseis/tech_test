package rest

import (
	"fmt"
	"log"
	"net/http"
	"tech_test/api"
	"tech_test/models"
	"tech_test/util"
)

// A new tech_test in POST body

func SalesPlacement(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /sales_placement service sales_placement
	//
	// Generates sales_placement record.
	//
	// Returns sales_placement status.
	//
	// ---
	// consumes:
	//  - application/json
	// parameters:
	//  - name: SalesPlacementRequest
	//    in: body
	//    description: Sales Placement Request
	//    required: true
	//    schema:
	//      "$ref": "#/definitions/SalesPlacementRequest"
	// responses:
	//   '200':
	//     $ref: "#/responses/status"
	//   '500':
	//     $ref: "#/responses/status"
	var salesPlacementRequest models.SalesPlacementRequest
	response := models.Status{Status: "success"}

	err := util.ParseJsonObj(r, &salesPlacementRequest)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = fmt.Sprintf("Failed to get parse JSON object: %v", err)
		log.Print(response.StatusMessage)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}
	response, err = api.SalesPlacement(salesPlacementRequest)
	if err != nil {
		log.Print(response)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}

	util.RespondWithJson(w, http.StatusOK, response)
}
