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
	"os"
	"socialnet/api"
	"socialnet/db"
	"socialnet/handler/login"
)

func main() {
	ctx := context.Background()

	cfg := parseEnv()
	storage := db.NewPostgres(db.NewPool(ctx, cfg))
	defer storage.Close()

	err := db.Migrate(cfg)
	if err != nil {
		log.Fatal(err)
	}

	s := api.NewServer(login.NewHandler(storage))

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
		err := storage.Ping(ctx)
		if err != nil {
			log.Printf("failed connecting to db: %v", err)
		} else {
			log.Println("successfully connected to db!")
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
				loggerMW,
			},
		},
	)

	addr := ":8080"
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

func parseEnv() db.PoolConfig {
	cfg := db.PoolConfig{
		Username: "root",
		Password: "secret",
		Host:     "database",
		Port:     "5432",
		DbName:   "social_net",
	}

	if user, exists := os.LookupEnv("POSTGRES_USER"); exists {
		cfg.Username = user
	}
	if password, exists := os.LookupEnv("POSTGRES_PASSWORD"); exists {
		cfg.Password = password
	}
	if dbName, exists := os.LookupEnv("POSTGRES_DB"); exists {
		cfg.DbName = dbName
	}

	return cfg
}

func loggerMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}
		defer log.Println(fmt.Sprintf("Method: %s has been handled.", fmt.Sprintf("%v://%v%v", scheme, r.Host, r.RequestURI)))

		// serve
		next.ServeHTTP(w, r)
	})
}
