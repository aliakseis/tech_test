// Code generated by swagger. DO NOT EDIT.
// Source: swagger.json
package static

const SwaggerJson = `
{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "The purpose of this application is to provide\ntech_test manipulation REST API",
    "title": "tech_test manipulation REST API",
    "version": "0.0.1"
  },
  "basePath": "/",
  "paths": {
    "/sale_placement": {
      "post": {
        "description": "Returns sale_placement status.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "service"
        ],
        "summary": "Generates sale_placement record.",
        "operationId": "sale_placement",
        "parameters": [
          {
            "description": "Sale Placement Request",
            "name": "SalePlacementRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SalePlacementRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/status"
          },
          "500": {
            "$ref": "#/responses/status"
          }
        }
      }
    },
    "/sales_adjustment": {
      "post": {
        "description": "Returns sales_adjustment status.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "service"
        ],
        "summary": "Generates sales_adjustment record.",
        "operationId": "sales_adjustment",
        "parameters": [
          {
            "description": "Sales Adjustment Request",
            "name": "SalesAdjustmentRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SalesAdjustmentRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/status"
          },
          "500": {
            "$ref": "#/responses/status"
          }
        }
      }
    },
    "/sales_placement": {
      "post": {
        "description": "Returns sales_placement status.",
        "consumes": [
          "application/json"
        ],
        "tags": [
          "service"
        ],
        "summary": "Generates sales_placement record.",
        "operationId": "sales_placement",
        "parameters": [
          {
            "description": "Sales Placement Request",
            "name": "SalesPlacementRequest",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/SalesPlacementRequest"
            }
          }
        ],
        "responses": {
          "200": {
            "$ref": "#/responses/status"
          },
          "500": {
            "$ref": "#/responses/status"
          }
        }
      }
    }
  },
  "definitions": {
    "SalePlacementRequest": {
      "type": "object",
      "properties": {
        "price": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Price"
        },
        "product_type": {
          "type": "string",
          "x-go-name": "ProductType"
        }
      },
      "x-go-package": "tech_test/models"
    },
    "SalesAdjustmentRequest": {
      "type": "object",
      "properties": {
        "addend": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Addend"
        },
        "divider": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Divider"
        },
        "multiplier": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Multiplier"
        },
        "product_type": {
          "type": "string",
          "x-go-name": "ProductType"
        }
      },
      "x-go-package": "tech_test/models"
    },
    "SalesPlacementRequest": {
      "type": "object",
      "properties": {
        "price": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Price"
        },
        "product_type": {
          "type": "string",
          "x-go-name": "ProductType"
        },
        "quantity": {
          "type": "integer",
          "format": "int64",
          "x-go-name": "Quantity"
        }
      },
      "x-go-package": "tech_test/models"
    }
  },
  "responses": {
    "status": {
      "description": "Represents JSON status object returned by Rest Service",
      "headers": {
        "status": {
          "type": "string"
        },
        "status_message": {
          "type": "string"
        }
      }
    }
  }
}`

