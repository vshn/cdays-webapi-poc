// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetManagedNamespaceHandlerFunc turns a function with the right signature into a get managed namespace handler
type GetManagedNamespaceHandlerFunc func(GetManagedNamespaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetManagedNamespaceHandlerFunc) Handle(params GetManagedNamespaceParams) middleware.Responder {
	return fn(params)
}

// GetManagedNamespaceHandler interface for that can handle valid get managed namespace params
type GetManagedNamespaceHandler interface {
	Handle(GetManagedNamespaceParams) middleware.Responder
}

// NewGetManagedNamespace creates a new http.Handler for the get managed namespace operation
func NewGetManagedNamespace(ctx *middleware.Context, handler GetManagedNamespaceHandler) *GetManagedNamespace {
	return &GetManagedNamespace{Context: ctx, Handler: handler}
}

/*GetManagedNamespace swagger:route GET /{clustername}/namespace/{customer}/{name} namespace getManagedNamespace

GetManagedNamespace get managed namespace API

*/
type GetManagedNamespace struct {
	Context *middleware.Context
	Handler GetManagedNamespaceHandler
}

func (o *GetManagedNamespace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetManagedNamespaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
