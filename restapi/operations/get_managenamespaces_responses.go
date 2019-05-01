// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/vshn/cdays-webapi-poc/models"
)

// GetManagenamespacesOKCode is the HTTP code returned for type GetManagenamespacesOK
const GetManagenamespacesOKCode int = 200

/*GetManagenamespacesOK returns a list of namespaces

swagger:response getManagenamespacesOK
*/
type GetManagenamespacesOK struct {

	/*list of namespaces
	  In: Body
	*/
	Payload []*models.Namespace `json:"body,omitempty"`
}

// NewGetManagenamespacesOK creates GetManagenamespacesOK with default headers values
func NewGetManagenamespacesOK() *GetManagenamespacesOK {

	return &GetManagenamespacesOK{}
}

// WithPayload adds the payload to the get managenamespaces o k response
func (o *GetManagenamespacesOK) WithPayload(payload []*models.Namespace) *GetManagenamespacesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get managenamespaces o k response
func (o *GetManagenamespacesOK) SetPayload(payload []*models.Namespace) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetManagenamespacesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.Namespace, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}