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
	"github.com/vshn/cdays-webapi-poc/restapi/operations/cluster"
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
		return namespace.NewGetManagedNamespacesOK().WithPayload(kubeClient.AllManagedNamespaces(params.Clustername))
	})

	api.NamespaceGetManagedNamespaceHandler = namespace.GetManagedNamespaceHandlerFunc(func(params namespace.GetManagedNamespaceParams) middleware.Responder {
		return namespace.NewGetManagedNamespaceOK().WithPayload(kubeClient.GetManagedNamespaceByName(params.Clustername, params.Name, params.Customer))
	})

	api.NamespaceCreateManagedNamespaceHandler = namespace.CreateManagedNamespaceHandlerFunc(func(params namespace.CreateManagedNamespaceParams) middleware.Responder {
		newNamespace, err := kubeClient.CreateManagedNamespace(params.Clustername, params.Customer, params.Body)

		if err != nil {
			fmt.Printf("error while creating a new namespace: %v\n", err)
			return nil
		}

		return namespace.NewCreateManagedNamespaceOK().WithPayload(newNamespace)
	})

	api.NamespaceDeleteManagedNamespaceHandler = namespace.DeleteManagedNamespaceHandlerFunc(func(params namespace.DeleteManagedNamespaceParams) middleware.Responder {
		deleted, err := kubeClient.DeleteManagedNamespace(params.Clustername, params.Customer, params.Name)
		if err != nil {
			fmt.Printf("can't delete managed namespace: %v", err)
		}

		return namespace.NewDeleteManagedNamespaceOK().WithPayload(deleted)
	})

	api.NamespaceUpdateManagedNamespaceHandler = namespace.UpdateManagedNamespaceHandlerFunc(func(params namespace.UpdateManagedNamespaceParams) middleware.Responder {
		updated, err := kubeClient.UpdateManagedNamespace(params.Clustername, params.Customer, params.Name, params.Body)

		if err != nil {
			fmt.Printf("error updating managed namespace: %v", err)
		}

		return namespace.NewUpdateManagedNamespaceOK().WithPayload(updated)
	})

	api.ClusterGetAllClustersHandler = cluster.GetAllClustersHandlerFunc(func(params cluster.GetAllClustersParams) middleware.Responder {
		return cluster.NewGetAllClustersOK().WithPayload(kubeClient.ListAllClusters())
	})

	api.ClusterRegisterClusterHandler = cluster.RegisterClusterHandlerFunc(func(params cluster.RegisterClusterParams) middleware.Responder {
		return cluster.NewRegisterClusterOK().WithPayload(kubeClient.AddCluster(params.Body))
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
