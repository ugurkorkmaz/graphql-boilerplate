package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"app/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var (
	_ctx          context.Context
	_chi          *chi.Mux
	_gql          *handler.Server
	appVersion    = "0.0.0"
	appName       string
	appPort       string
	appSecretKey  string
	appHeaderName string
	gqlComplexity int
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Set variable values from environment
	appName = os.Getenv("APP_NAME")
	appPort = os.Getenv("APP_PORT")
	appSecretKey = os.Getenv("APP_SECRET_KEY")
	appHeaderName = os.Getenv("APP_SECRET_HEADER_NAME")

	// Convert GQL_COMPLEXITY_LIMIT to integer
	gqlComplexity, err = strconv.Atoi(os.Getenv("GQL_COMPLEXITY_LIMIT"))
	if err != nil {
		log.Fatal("Invalid GQL_COMPLEXITY_LIMIT value")
	}

	_ctx = context.Background()
	_chi = chi.NewRouter()
}

func main() {
	// Add middleware
	_chi.Use(middleware.Logger, GlobalContext)

	// Add routes and configure the GraphQL server
	_gql = handler.NewDefaultServer(resolver.NewSchema())
	_gql.AddTransport(transport.POST{})
	_gql.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 5 * time.Second,
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	_gql.AddTransport(transport.Options{})
	_gql.AddTransport(transport.GET{})
	_gql.AddTransport(transport.POST{})
	_gql.AddTransport(transport.MultipartForm{
		MaxUploadSize: 10 * 1024 * 1024,
		MaxMemory:     32 << 20,
	})
	_gql.Use(extension.FixedComplexityLimit(gqlComplexity))

	// Add secured routes
	_chi.Group(func(r chi.Router) {
		_chi.Handle("/graphql", _gql)
		_chi.Handle("/websocket", _gql)
		_chi.Handle("/", playground.Handler(appName, "/graphql"))
	}).Use(Secret)

	// Start the HTTP server
	http.ListenAndServe(appPort, _chi)
}

// GlobalContext creates a new context for each HTTP request and adds the request data to the context.
// This middleware can be used to easily access the request data.
func GlobalContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_ctx = context.WithValue(r.Context(), "request", r)
		next.ServeHTTP(w, r.WithContext(_ctx))
	})
}

// Secret is a middleware that checks the secret key in the header.
func Secret(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := r.Header.Get(appHeaderName)
		if secretKey != appSecretKey {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"status": false, "message": "Unauthorized"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
