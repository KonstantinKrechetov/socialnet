package main

import (
	"context"
	"encoding/json"
	chimiddleware "github.com/deepmap/oapi-codegen/pkg/chi-middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"net/http"
	"socialnet/api"
)

func main() {
	s := api.NewServer()

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	// Add swagger UI endpoints
	router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(swagger)
	})
	router.Handle("/swagger/", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		Path:    "/swagger/",
		SpecURL: "/swagger/doc.json",
	}, nil))
	router.Get("/_healthcheck/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("successful response")
	})

	// Enable validation of incoming requests
	validator := chimiddleware.OapiRequestValidatorWithOptions(
		swagger,
		&chimiddleware.Options{
			Options: openapi3filter.Options{
				AuthenticationFunc: func(c context.Context, input *openapi3filter.AuthenticationInput) error {
					return nil
				},
			},
		},
	)

	apiServer := api.HandlerWithOptions(
		api.NewStrictHandler(s, nil),
		api.ChiServerOptions{
			BaseURL:    "",
			BaseRouter: router,
			Middlewares: []api.MiddlewareFunc{
				validator,
			},
		},
	)

	addr := ":8000"
	httpServer := http.Server{
		Addr:    addr,
		Handler: apiServer,
	}

	log.Println("Server listening on", addr)
	err = httpServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
		return
	}
}
