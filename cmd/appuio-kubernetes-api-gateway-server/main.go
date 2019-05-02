//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//

package main

import (
	"fmt"
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	flags "github.com/jessevdk/go-flags"
	"github.com/vshn/cdays-webapi-poc/kube"
	"github.com/vshn/cdays-webapi-poc/restapi"
	"github.com/vshn/cdays-webapi-poc/restapi/operations"
	"github.com/vshn/cdays-webapi-poc/restapi/operations/namespace"
)

func main() {

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewAppuioKubernetesAPIGatewayAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	kubeClient, err := kube.NewClientManager(true)
	if err != nil {
		fmt.Println(err)
	}

	api.NamespaceGetManagedNamespacesHandler = namespace.GetManagedNamespacesHandlerFunc(func(params namespace.GetManagedNamespacesParams) middleware.Responder {
		return namespace.NewGetManagedNamespacesOK().WithPayload(kubeClient.AllManagedNamespaces())
	})

	api.NamespaceGetManagedNamespaceHandler = namespace.GetManagedNamespaceHandlerFunc(func(params namespace.GetManagedNamespaceParams) middleware.Responder {
		return namespace.NewGetManagedNamespaceOK().WithPayload(kubeClient.GetNamespaceByName(params.Name, params.Namespace))
	})

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Appuio Kubernetes API gateway"
	parser.LongDescription = swaggerSpec.Spec().Info.Description

	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			log.Fatalln(err)
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.Port = 8080

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
