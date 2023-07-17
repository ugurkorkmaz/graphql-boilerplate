package main

import (
	"context"
	"net/http"
	"time"

	"app/resolver"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
)

var (
	_ctx context.Context
	_chi *chi.Mux
	_gql *handler.Server
)

var (
	// Version is the version of the application.
	version string = "0.0.0"
	// Name of the application.
	name string = "app"
	// Port the application will listen on.
	port string = ":8080"
	// SecretKey is the secret key used to sign JWT tokens.
	key string = "secret"
	// HeaderName is the name of the header that contains the secret key.
	headerName string = "X-App-Key"
	// GraphQLComplexityLimit is the maximum complexity of a GraphQL query.
	gcl int = 35
)

func init() {
	// Init global variables
	_ctx = context.Background()
	// Init chi router
	_chi = chi.NewRouter()
}

func main() {
	// Add middleware
	_chi.Use(middleware.Logger, GlobalContext)

	// Add routes and configure the GraphQL server
	_gql = handler.NewDefaultServer(resolver.NewSchema())
	// Add POST transport
	_gql.AddTransport(transport.POST{})
	// Add websocket transport
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
	// Add OPTIONS transport
	_gql.AddTransport(transport.Options{})
	// Add GET transport
	_gql.AddTransport(transport.GET{})
	// Add POST transport
	_gql.AddTransport(transport.POST{})
	// Add multipart form transport
	_gql.AddTransport(transport.MultipartForm{
		MaxUploadSize: 10 * 1024 * 1024,
		MaxMemory:     32 << 20,
	})
	// Complexity limit
	_gql.Use(extension.FixedComplexityLimit(gcl))

	// Add routes Secured
	_chi.Group(func(r chi.Router) {
		_chi.Handle("/graphql", _gql)
		_chi.Handle("/websocket", _gql)
		_chi.Handle("/", playground.Handler(name, "/graphql"))
	}).Use(Secret)

	http.ListenAndServe(port, _chi)
}

// GlobalContext creates a new context for each HTTP request and adds the request data to the context.
// This middleware can be used to easily access the request data.
func GlobalContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Create a new context and add the request data to it using the custom struct type as key.
		//lint:ignore SA1029 We need to use the deprecated function here.
		_ctx = context.WithValue(r.Context(), "request", r)
		// Call the http.Handler with the new context.
		next.ServeHTTP(w, r.WithContext(_ctx))
	})
}

// Secret is a middleware that checks the secret key in the header.
func Secret(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		secretKey := r.Header.Get(headerName)
		if secretKey != key {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`{"status":false,"message":"Unauthorized"}`))
			return
		}
		next.ServeHTTP(w, r)
	})
}
