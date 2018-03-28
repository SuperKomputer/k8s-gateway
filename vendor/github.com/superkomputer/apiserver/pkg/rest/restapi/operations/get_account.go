// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetAccountHandlerFunc turns a function with the right signature into a get account handler
type GetAccountHandlerFunc func(GetAccountParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetAccountHandlerFunc) Handle(params GetAccountParams) middleware.Responder {
	return fn(params)
}

// GetAccountHandler interface for that can handle valid get account params
type GetAccountHandler interface {
	Handle(GetAccountParams) middleware.Responder
}

// NewGetAccount creates a new http.Handler for the get account operation
func NewGetAccount(ctx *middleware.Context, handler GetAccountHandler) *GetAccount {
	return &GetAccount{Context: ctx, Handler: handler}
}

/*GetAccount swagger:route GET /clusters/{clusterId}/account/{username} getAccount

get payable account of a user on cluster

get payable account of a user on cluster

*/
type GetAccount struct {
	Context *middleware.Context
	Handler GetAccountHandler
}

func (o *GetAccount) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetAccountParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}