// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	graceful "github.com/tylerb/graceful"

	"github.com/fanzhangio/superkomputer/pkg/rest/handlers"
	"github.com/fanzhangio/superkomputer/pkg/rest/models"
	"github.com/fanzhangio/superkomputer/pkg/rest/restapi/operations"
)

//go:generate swagger generate server --target .. --name api --spec ../../../swagger/swagger.yml

func configureFlags(api *operations.API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CreateAccountHandler = operations.CreateAccountHandlerFunc(func(params operations.CreateAccountParams) middleware.Responder {
		// Create Account
		if err := handlers.CreateAccount(params.ClusterID, params.Username, params.Account); err != nil {
			return operations.NewCreateAccountDefault(404).WithPayload(&models.Error{Code: 400, Message: swag.String(err.Error())})
		}
		return operations.NewCreateAccountAccepted()
	})
	api.CreateClusterHandler = operations.CreateClusterHandlerFunc(func(params operations.CreateClusterParams) middleware.Responder {
		// Create Cluster
		if err := handlers.CreateCluster(params.Cluster); err != nil {
			return operations.NewCreateClusterDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return operations.NewCreateClusterAccepted()
	})
	api.CreateUserHandler = operations.CreateUserHandlerFunc(func(params operations.CreateUserParams) middleware.Responder {
		// Create User
		if err := handlers.CreateUser(params.User); err != nil {
			return operations.NewCreateUserDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return operations.NewCreateUserAccepted()
	})
	api.DeleteAccountHandler = operations.DeleteAccountHandlerFunc(func(params operations.DeleteAccountParams) middleware.Responder {
		// Delete Account
		if err := handlers.DeleteAccount(params.ClusterID, params.Username); err != nil {
			return operations.NewDeleteAccountNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return operations.NewDeleteAccountAccepted()
	})
	api.DeleteClusterHandler = operations.DeleteClusterHandlerFunc(func(params operations.DeleteClusterParams) middleware.Responder {
		// Delete Cluster
		if err := handlers.DeleteCluster(params.ClusterID); err != nil {
			return operations.NewDeleteClusterNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return operations.NewDeleteClusterAccepted()
	})
	api.DeleteUserHandler = operations.DeleteUserHandlerFunc(func(params operations.DeleteUserParams) middleware.Responder {
		// Delete User
		if err := handlers.DeleteUser(params.Username); err != nil {
			return operations.NewDeleteUserNotFound().WithPayload(&models.Error{Code: 402, Message: swag.String(err.Error())})
		}
		return operations.NewDeleteUserAccepted()
	})
	api.FetchUserClustersHandler = operations.FetchUserClustersHandlerFunc(func(params operations.FetchUserClustersParams) middleware.Responder {
		// Fetch User Clusters
		clusters, err := handlers.FetchUserClusters(params.Username)
		if err != nil {
			return operations.NewFetchUserClustersDefault(400).WithPayload(&models.Error{Code: 400, Message: swag.String(err.Error())})
		}
		return operations.NewFetchUserClustersOK().WithPayload(clusters)
	})
	api.GetAccountHandler = operations.GetAccountHandlerFunc(func(params operations.GetAccountParams) middleware.Responder {
		// Get Account
		accounts, err := handlers.GetAccount(params.ClusterID, params.Username)
		if err != nil {
			return operations.NewGetAccountNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return operations.NewGetAccountOK().WithPayload(accounts)
	})
	api.GetClusterHandler = operations.GetClusterHandlerFunc(func(params operations.GetClusterParams) middleware.Responder {
		// Get Cluster
		cluster, err := handlers.GetCluster(params.ClusterID)
		if err != nil {
			return operations.NewGetClusterNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return operations.NewGetClusterOK().WithPayload(cluster)
	})
	api.GetUserHandler = operations.GetUserHandlerFunc(func(params operations.GetUserParams) middleware.Responder {
		// Get User
		user, err := handlers.GetUser(params.Username)
		if err != nil {
			return operations.NewGetUserNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return operations.NewGetUserOK().WithPayload(user)
	})
	api.ListClustersHandler = operations.ListClustersHandlerFunc(func(params operations.ListClustersParams) middleware.Responder {
		// List Clusters
		return operations.NewListClustersOK().WithPayload(handlers.ListClusters())
	})
	api.ListUsersHandler = operations.ListUsersHandlerFunc(func(params operations.ListUsersParams) middleware.Responder {
		// List Users
		return operations.NewListUsersOK().WithPayload(handlers.ListUsers())
	})
	api.LoginHandler = operations.LoginHandlerFunc(func(params operations.LoginParams) middleware.Responder {
		return middleware.NotImplemented("operation .Login has not yet been implemented")
	})
	api.UpdateClusterHandler = operations.UpdateClusterHandlerFunc(func(params operations.UpdateClusterParams) middleware.Responder {
		// Update Cluster
		if err := handlers.UpdateCluster(params.ClusterID, params.Cluster); err != nil {
			return operations.NewUpdateClusterNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return operations.NewUpdateClusterAccepted()
	})
	api.UpdateUserHandler = operations.UpdateUserHandlerFunc(func(params operations.UpdateUserParams) middleware.Responder {
		// Update User
		if err := handlers.UpdateUser(params.Username, params.User); err != nil {
			return operations.NewUpdateUserNotFound().WithPayload(&models.Error{Code: 404, Message: swag.String(err.Error())})
		}
		return operations.NewUpdateUserAccepted()
	})

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
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
