package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/go-chi/chi/v5"
	"github.com/go-openapi/runtime/middleware"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"net/http"
	"socialnet/api"
	"socialnet/handler/login"
	"socialnet/storage"
)

func main() {
	ctx := context.Background()

	dbConfig := storage.PoolConfig{
		Username: "postgres",
		Password: "postgres",
		Host:     "localhost",
		Port:     "5433",
		Dbname:   "postgres",
	}
	db := storage.NewPostgres(storage.NewPool(ctx, dbConfig))
	defer db.Close()

	s := api.NewServer(login.NewHandler())

	swagger, err := api.GetSwagger()
	if err != nil {
		log.Fatal(err)
	}

	router := chi.NewRouter()

	router.Get("/swagger/doc.json", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(swagger)
	})
	router.Handle("/swagger/", middleware.SwaggerUI(middleware.SwaggerUIOpts{
		Path:    "/swagger/",
		SpecURL: "/swagger/doc.json",
	}, nil))
	router.Get("/_healthcheck/", func(w http.ResponseWriter, r *http.Request) {
		err := db.Ping(ctx)
		if err != nil {
			fmt.Printf("failed connecting to db: %v", err)
		} else {
			fmt.Println("successfully connected to db!")
		}

		json.NewEncoder(w).Encode("successful response")
	})

	// Enable validation of incoming requests
	validator := nethttpmiddleware.OapiRequestValidatorWithOptions(
		swagger,
		&nethttpmiddleware.Options{
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
