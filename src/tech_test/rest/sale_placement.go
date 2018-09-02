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

func SalePlacement(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /sale_placement service sale_placement
	//
	// Generates sale_placement record.
	//
	// Returns sale_placement status.
	//
	// ---
	// consumes:
	//  - application/json
	// parameters:
	//  - name: SalePlacementRequest
	//    in: body
	//    description: Sale Placement Request
	//    required: true
	//    schema:
	//      "$ref": "#/definitions/SalePlacementRequest"
	// responses:
	//   '200':
	//     $ref: "#/responses/status"
	//   '500':
	//     $ref: "#/responses/status"
	var salePlacementRequest models.SalePlacementRequest
	response := models.Status{Status: "success"}

	err := util.ParseJsonObj(r, &salePlacementRequest)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = fmt.Sprintf("Failed to get parse JSON object: %v", err)
		log.Print(response.StatusMessage)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}
	response, err = api.SalePlacement(salePlacementRequest)
	if err != nil {
		log.Print(response)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}

	util.RespondWithJson(w, http.StatusOK, response)
}
