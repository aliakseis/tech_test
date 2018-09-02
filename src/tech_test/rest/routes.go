package rest

import (
	"tech_test/static"
	"net/http"
	"strings"
	"github.com/go-openapi/runtime/middleware"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.Handler
}

func getHandler(handlerFunc http.HandlerFunc) http.Handler {
	return handlerFunc
}

func getRoute(name string,
	method string,
	pattern string,
	handlerFunc http.HandlerFunc) Route {
		return Route{name, method, pattern, handlerFunc}
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		referer := r.Referer()
		if strings.HasPrefix(referer, "http://petstore.swagger.io/?url=http%3A%2F%2F") {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		}

		h.ServeHTTP(w, r)
	})
}

type Routes []Route

func routes() Routes {
	return Routes{
		Route{
			"sales_adjustment",
			"POST",
			"/sales_adjustment",
			accessControl(getHandler(SalesAdjustment)),
		},
		Route{
			"sales_adjustment",
			"OPTIONS",
			"/sales_adjustment",
			accessControl(getHandler(PostOptionsHandler)),
		},

		Route{
			"sales_placement",
			"POST",
			"/sales_placement",
			accessControl(getHandler(SalesPlacement)),
		},
		Route{
			"sales_placement",
			"OPTIONS",
			"/sales_placement",
			accessControl(getHandler(PostOptionsHandler)),
		},

		Route{
			"tech_test",
			"POST",
			"/sale_placement",
			accessControl(getHandler(SalePlacement)),
		},
		Route{
			"hoppity",
			"OPTIONS",
			"/sale_placement",
			accessControl(getHandler(PostOptionsHandler)),
		},
		
		Route{
			"api-docs",
			"GET",
			"/api-docs",
			middleware.Redoc(middleware.RedocOpts{Path: "api-docs"}, nil),
		},
		Route{
			"swagger.json",
			"GET",
			"/swagger.json",
			accessControl(middleware.Spec("", []byte(static.SwaggerJson), nil)),
		},
		getRoute(
			"defaulthandler",
			"GET",
			"/",
			DefaultHandler,
		),
	}
}
