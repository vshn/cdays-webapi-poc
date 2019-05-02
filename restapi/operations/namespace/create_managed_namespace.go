// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// CreateManagedNamespaceHandlerFunc turns a function with the right signature into a create managed namespace handler
type CreateManagedNamespaceHandlerFunc func(CreateManagedNamespaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn CreateManagedNamespaceHandlerFunc) Handle(params CreateManagedNamespaceParams) middleware.Responder {
	return fn(params)
}

// CreateManagedNamespaceHandler interface for that can handle valid create managed namespace params
type CreateManagedNamespaceHandler interface {
	Handle(CreateManagedNamespaceParams) middleware.Responder
}

// NewCreateManagedNamespace creates a new http.Handler for the create managed namespace operation
func NewCreateManagedNamespace(ctx *middleware.Context, handler CreateManagedNamespaceHandler) *CreateManagedNamespace {
	return &CreateManagedNamespace{Context: ctx, Handler: handler}
}

/*CreateManagedNamespace swagger:route PUT /namespace/{customer} namespace createManagedNamespace

CreateManagedNamespace create managed namespace API

*/
type CreateManagedNamespace struct {
	Context *middleware.Context
	Handler CreateManagedNamespaceHandler
}

func (o *CreateManagedNamespace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewCreateManagedNamespaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}