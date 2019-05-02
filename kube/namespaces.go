//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//
package kube

import (
	"fmt"

	"context"

	"github.com/go-openapi/swag"
	controlv1 "github.com/vshn/cdays-namespace-poc/pkg/apis/control/v1alpha1"
	"github.com/vshn/cdays-webapi-poc/models"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

func (k *ClientManager) AllManagedNamespaces() []*models.Namespace {
	namespaces := make([]*models.Namespace, 0)

	managedNamespaces := &controlv1.ManagedNamespaceList{}

	err := k.CRDClient.List(context.Background(), &client.ListOptions{}, managedNamespaces)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, space := range managedNamespaces.Items {
		namespaces = append(namespaces, &models.Namespace{
			Name: swag.String(space.GetName()),
		})
	}

	return namespaces
}

func (k *ClientManager) GetNamespaceByName(name, namespace string) *models.Namespace {

	key := client.ObjectKey{
		Namespace: namespace,
		Name:      name,
	}

	managedNamespaces := &controlv1.ManagedNamespace{}

	k.CRDClient.Get(context.Background(), key, managedNamespaces)

	return &models.Namespace{
		Name: swag.String(managedNamespaces.GetName()),
	}
}
