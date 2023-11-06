package main

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	chimiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-openapi/runtime/middleware"
	nethttpmiddleware "github.com/oapi-codegen/nethttp-middleware"
	"log"
	"net/http"
	"os"
	"socialnet/api"
	"socialnet/db"
	"socialnet/handler/login"
	"socialnet/handler/user_get_id"
	"socialnet/handler/user_register"
	"socialnet/server"
	"socialnet/service/encryptor"
)

func main() {
	ctx := context.Background()

	cfg := parseEnv()
	storage := db.NewPostgres(db.NewPool(cfg))
	defer storage.Close()

	err := db.Migrate(cfg)
	if err != nil {
		log.Fatal(err)
	}

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

	encryptorInstance := encryptor.New()
	s := &server.Server{
		LoginHandler:        login.NewHandler(storage, encryptorInstance),
		UserGetIdHandler:    user_get_id.NewHandler(storage),
		UserRegisterHandler: user_register.NewHandler(storage, encryptorInstance),
	}

	apiServer := api.HandlerWithOptions(
		s,
		api.ChiServerOptions{
			BaseRouter: router,
			Middlewares: []api.MiddlewareFunc{
				nethttpmiddleware.OapiRequestValidator(swagger), // Enable validation of incoming requests
				chimiddleware.Logger,
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
	// uncomment for debug
	//host := "localhost"
	//port := "5433"
	host := "database"
	port := "5432"

	cfg := db.PoolConfig{
		Username: "root",
		Password: "secret",
		Host:     host,
		Port:     port,
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
