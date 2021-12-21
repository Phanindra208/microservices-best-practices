// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"path"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/gofrs/uuid"
	"github.com/justinas/alice"

	feedhandler "github.com/iAmPlus/microservice/restapi/handlers/feed"
	usershandlers "github.com/iAmPlus/microservice/restapi/handlers/users"
	"github.com/iAmPlus/microservice/restapi/operations"
	"github.com/iAmPlus/microservice/restapi/operations/feed"
	"github.com/iAmPlus/microservice/restapi/operations/health"
	"github.com/iAmPlus/microservice/restapi/operations/users"
	"github.com/iAmPlus/microservice/tracing"
)

//go:generate swagger generate server --target ../../microservice_best_practice --name Microservice --spec ../../../../../../../../var/folders/fw/kv072vln3ps_hcq809qryz7r0000gp/T/swagger.yaml462203815 --model-package models/apimodels --principal models.Principal

func configureFlags(api *operations.MicroserviceAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.MicroserviceAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.FeedCreateFeedHandler = feed.CreateFeedHandlerFunc(feedhandler.InserPost)

	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(usershandlers.InserUser)

	api.UsersFollowUserHandler = users.FollowUserHandlerFunc(usershandlers.FollowUser)

	api.FeedFriendsFeedHandler = feed.FriendsFeedHandlerFunc(feedhandler.GetUserFrendsFeed)

	api.FeedGetRelatedFeedHandler = feed.GetRelatedFeedHandlerFunc(feedhandler.GetRelatedFeed)

	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(usershandlers.GetAllUsers)

	if api.HealthLivenessHandler == nil {
		api.HealthLivenessHandler = health.LivenessHandlerFunc(func(params health.LivenessParams) middleware.Responder {
			return middleware.NotImplemented("operation health.Liveness has not yet been implemented")
		})
	}
	if api.HealthReadinessHandler == nil {
		api.HealthReadinessHandler = health.ReadinessHandlerFunc(func(params health.ReadinessParams) middleware.Responder {
			return middleware.NotImplemented("operation health.Readiness has not yet been implemented")
		})
	}

	api.UsersUnFollowUserHandler = users.UnFollowUserHandlerFunc(usershandlers.UnFollowUser)

	api.FeedUserFeedHandler = feed.UserFeedHandlerFunc(feedhandler.GetUserFeed)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

func setupMiddlewares(handler http.Handler) http.Handler {
	return alice.New(tracing.GetMiddleware()).Then(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	mw := make([]alice.Constructor, 0)

	// mw = append(mw, cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"*"},
	// 	AllowedMethods:   []string{"HEAD", "GET", "POST", "PUT", "PATCH", "DELETE"},
	// 	AllowedHeaders:   []string{"*"},
	// 	AllowCredentials: true,
	// }).Handler)

	mw = append(mw, alice.Constructor(func(h http.Handler) http.Handler {
		return middleware.Redoc(middleware.RedocOpts{
			Title:   "microservice API documentation",
			Path:    path.Join("/", "docs"),
			SpecURL: path.Join("/", "swagger.json"),
		}, middleware.Spec("/", SwaggerJSON, h))
	}))
	//	mw = append(mw, LogControlMiddleware)
	mw = append(mw, GetCorrelationIDMiddleware)
	mw = append(mw, GetCacheControlMiddleware)
	return alice.New(mw...).Then(handler)
}

// GetCacheControlMiddleware disables response caching.
func GetCacheControlMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-store")
		h.ServeHTTP(w, r)
	})
}

// GetCorrelationIDMiddleware loggs .
func GetCorrelationIDMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cID := newCorrelationID()
		r.Header.Set("X-Correlation-ID", cID)
		w.Header().Set("X-Correlation-ID", cID)
		h.ServeHTTP(w, r)
	})
}

func newCorrelationID() string {
	return uuid.NewV5(
		uuid.NamespaceURL,
		fmt.Sprintf("%s:%d", "hotels.iamplus.services", time.Now().UnixNano()),
	).String()
}
