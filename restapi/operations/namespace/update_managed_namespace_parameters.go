// Code generated by go-swagger; DO NOT EDIT.

package namespace

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vshn/cdays-webapi-poc/models"
)

// NewUpdateManagedNamespaceParams creates a new UpdateManagedNamespaceParams object
// no default values defined in spec.
func NewUpdateManagedNamespaceParams() UpdateManagedNamespaceParams {

	return UpdateManagedNamespaceParams{}
}

// UpdateManagedNamespaceParams contains all the bound params for the update managed namespace operation
// typically these are obtained from a http.Request
//
// swagger:parameters updateManagedNamespace
type UpdateManagedNamespaceParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: body
	*/
	Body *models.Namespace
	/*
	  Required: true
	  In: path
	*/
	Clustername string
	/*
	  Required: true
	  In: path
	*/
	Customer string
	/*
	  Required: true
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewUpdateManagedNamespaceParams() beforehand.
func (o *UpdateManagedNamespaceParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.Namespace
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			res = append(res, errors.NewParseError("body", "body", "", err))
		} else {
			// validate body object
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Body = &body
			}
		}
	}
	rClustername, rhkClustername, _ := route.Params.GetOK("clustername")
	if err := o.bindClustername(rClustername, rhkClustername, route.Formats); err != nil {
		res = append(res, err)
	}

	rCustomer, rhkCustomer, _ := route.Params.GetOK("customer")
	if err := o.bindCustomer(rCustomer, rhkCustomer, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindClustername binds and validates parameter Clustername from path.
func (o *UpdateManagedNamespaceParams) bindClustername(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Clustername = raw

	return nil
}

// bindCustomer binds and validates parameter Customer from path.
func (o *UpdateManagedNamespaceParams) bindCustomer(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Customer = raw

	return nil
}

// bindName binds and validates parameter Name from path.
func (o *UpdateManagedNamespaceParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	o.Name = raw

	return nil
}
