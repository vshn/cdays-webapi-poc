// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DeleteManagedNamespaceHandlerFunc turns a function with the right signature into a delete managed namespace handler
type DeleteManagedNamespaceHandlerFunc func(DeleteManagedNamespaceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteManagedNamespaceHandlerFunc) Handle(params DeleteManagedNamespaceParams) middleware.Responder {
	return fn(params)
}

// DeleteManagedNamespaceHandler interface for that can handle valid delete managed namespace params
type DeleteManagedNamespaceHandler interface {
	Handle(DeleteManagedNamespaceParams) middleware.Responder
}

// NewDeleteManagedNamespace creates a new http.Handler for the delete managed namespace operation
func NewDeleteManagedNamespace(ctx *middleware.Context, handler DeleteManagedNamespaceHandler) *DeleteManagedNamespace {
	return &DeleteManagedNamespace{Context: ctx, Handler: handler}
}

/*DeleteManagedNamespace swagger:route DELETE /namespace/{customer}/{name} namespace deleteManagedNamespace

DeleteManagedNamespace delete managed namespace API

*/
type DeleteManagedNamespace struct {
	Context *middleware.Context
	Handler DeleteManagedNamespaceHandler
}

func (o *DeleteManagedNamespace) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDeleteManagedNamespaceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}