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

func SalesAdjustment(w http.ResponseWriter, r *http.Request) {
	// swagger:operation POST /sales_adjustment service sales_adjustment
	//
	// Generates sales_adjustment record.
	//
	// Returns sales_adjustment status.
	//
	// ---
	// consumes:
	//  - application/json
	// parameters:
	//  - name: SalesAdjustmentRequest
	//    in: body
	//    description: Sales Adjustment Request
	//    required: true
	//    schema:
	//      "$ref": "#/definitions/SalesAdjustmentRequest"
	// responses:
	//   '200':
	//     $ref: "#/responses/status"
	//   '500':
	//     $ref: "#/responses/status"
	var salesAdjustmentRequest models.SalesAdjustmentRequest
	response := models.Status{Status: "success"}

	err := util.ParseJsonObj(r, &salesAdjustmentRequest)
	if err != nil {
		response.Status = "error"
		response.StatusMessage = fmt.Sprintf("Failed to get parse JSON object: %v", err)
		log.Print(response.StatusMessage)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}
	response, err = api.SalesAdjustment(salesAdjustmentRequest)
	if err != nil {
		log.Print(response)
		util.RespondWithJson(w, http.StatusInternalServerError, response)
		return
	}

	util.RespondWithJson(w, http.StatusOK, response)
}
