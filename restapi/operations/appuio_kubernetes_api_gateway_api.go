// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/vshn/cdays-webapi-poc/restapi/operations/namespace"
)

// NewAppuioKubernetesAPIGatewayAPI creates a new AppuioKubernetesAPIGateway instance
func NewAppuioKubernetesAPIGatewayAPI(spec *loads.Document) *AppuioKubernetesAPIGatewayAPI {
	return &AppuioKubernetesAPIGatewayAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		customConsumers:     make(map[string]runtime.Consumer),
		customProducers:     make(map[string]runtime.Producer),
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		NamespaceCreateManagedNamespaceHandler: namespace.CreateManagedNamespaceHandlerFunc(func(params namespace.CreateManagedNamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation NamespaceCreateManagedNamespace has not yet been implemented")
		}),
		NamespaceDeleteManagedNamespaceHandler: namespace.DeleteManagedNamespaceHandlerFunc(func(params namespace.DeleteManagedNamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation NamespaceDeleteManagedNamespace has not yet been implemented")
		}),
		NamespaceGetManagedNamespaceHandler: namespace.GetManagedNamespaceHandlerFunc(func(params namespace.GetManagedNamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation NamespaceGetManagedNamespace has not yet been implemented")
		}),
		NamespaceGetManagedNamespacesHandler: namespace.GetManagedNamespacesHandlerFunc(func(params namespace.GetManagedNamespacesParams) middleware.Responder {
			return middleware.NotImplemented("operation NamespaceGetManagedNamespaces has not yet been implemented")
		}),
		NamespaceUpdateManagedNamespaceHandler: namespace.UpdateManagedNamespaceHandlerFunc(func(params namespace.UpdateManagedNamespaceParams) middleware.Responder {
			return middleware.NotImplemented("operation NamespaceUpdateManagedNamespace has not yet been implemented")
		}),
	}
}

/*AppuioKubernetesAPIGatewayAPI the appuio kubernetes API gateway API */
type AppuioKubernetesAPIGatewayAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	customConsumers map[string]runtime.Consumer
	customProducers map[string]runtime.Producer
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implementation in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// NamespaceCreateManagedNamespaceHandler sets the operation handler for the create managed namespace operation
	NamespaceCreateManagedNamespaceHandler namespace.CreateManagedNamespaceHandler
	// NamespaceDeleteManagedNamespaceHandler sets the operation handler for the delete managed namespace operation
	NamespaceDeleteManagedNamespaceHandler namespace.DeleteManagedNamespaceHandler
	// NamespaceGetManagedNamespaceHandler sets the operation handler for the get managed namespace operation
	NamespaceGetManagedNamespaceHandler namespace.GetManagedNamespaceHandler
	// NamespaceGetManagedNamespacesHandler sets the operation handler for the get managed namespaces operation
	NamespaceGetManagedNamespacesHandler namespace.GetManagedNamespacesHandler
	// NamespaceUpdateManagedNamespaceHandler sets the operation handler for the update managed namespace operation
	NamespaceUpdateManagedNamespaceHandler namespace.UpdateManagedNamespaceHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AppuioKubernetesAPIGatewayAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AppuioKubernetesAPIGatewayAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AppuioKubernetesAPIGatewayAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AppuioKubernetesAPIGatewayAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AppuioKubernetesAPIGatewayAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AppuioKubernetesAPIGatewayAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AppuioKubernetesAPIGatewayAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AppuioKubernetesAPIGatewayAPI
func (o *AppuioKubernetesAPIGatewayAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.NamespaceCreateManagedNamespaceHandler == nil {
		unregistered = append(unregistered, "namespace.CreateManagedNamespaceHandler")
	}

	if o.NamespaceDeleteManagedNamespaceHandler == nil {
		unregistered = append(unregistered, "namespace.DeleteManagedNamespaceHandler")
	}

	if o.NamespaceGetManagedNamespaceHandler == nil {
		unregistered = append(unregistered, "namespace.GetManagedNamespaceHandler")
	}

	if o.NamespaceGetManagedNamespacesHandler == nil {
		unregistered = append(unregistered, "namespace.GetManagedNamespacesHandler")
	}

	if o.NamespaceUpdateManagedNamespaceHandler == nil {
		unregistered = append(unregistered, "namespace.UpdateManagedNamespaceHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AppuioKubernetesAPIGatewayAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AppuioKubernetesAPIGatewayAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *AppuioKubernetesAPIGatewayAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *AppuioKubernetesAPIGatewayAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}

		if c, ok := o.customConsumers[mt]; ok {
			result[mt] = c
		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *AppuioKubernetesAPIGatewayAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}

		if p, ok := o.customProducers[mt]; ok {
			result[mt] = p
		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AppuioKubernetesAPIGatewayAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the appuio kubernetes API gateway API
func (o *AppuioKubernetesAPIGatewayAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AppuioKubernetesAPIGatewayAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/namespace/{customer}"] = namespace.NewCreateManagedNamespace(o.context, o.NamespaceCreateManagedNamespaceHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/namespace/{customer}/{name}"] = namespace.NewDeleteManagedNamespace(o.context, o.NamespaceDeleteManagedNamespaceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/namespace/{customer}/{name}"] = namespace.NewGetManagedNamespace(o.context, o.NamespaceGetManagedNamespaceHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/namespaces"] = namespace.NewGetManagedNamespaces(o.context, o.NamespaceGetManagedNamespacesHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/namespace/{customer}/{name}"] = namespace.NewUpdateManagedNamespace(o.context, o.NamespaceUpdateManagedNamespaceHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AppuioKubernetesAPIGatewayAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middleware as you see fit
func (o *AppuioKubernetesAPIGatewayAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}

// RegisterConsumer allows you to add (or override) a consumer for a media type.
func (o *AppuioKubernetesAPIGatewayAPI) RegisterConsumer(mediaType string, consumer runtime.Consumer) {
	o.customConsumers[mediaType] = consumer
}

// RegisterProducer allows you to add (or override) a producer for a media type.
func (o *AppuioKubernetesAPIGatewayAPI) RegisterProducer(mediaType string, producer runtime.Producer) {
	o.customProducers[mediaType] = producer
}
