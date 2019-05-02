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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
			Name:        swag.String(space.GetName()),
			Description: space.Spec.Description,
			Customer:    swag.String(space.GetNamespace()),
		})
	}

	return namespaces
}

func (k *ClientManager) GetManagedNamespaceByName(name, customer string) *models.Namespace {

	key := client.ObjectKey{
		Namespace: customer,
		Name:      name,
	}

	managedNamespaces := &controlv1.ManagedNamespace{}

	k.CRDClient.Get(context.Background(), key, managedNamespaces)

	return &models.Namespace{
		Name:        swag.String(managedNamespaces.GetName()),
		Description: managedNamespaces.Spec.Description,
		Customer:    swag.String(managedNamespaces.GetNamespace()),
	}
}

func (k *ClientManager) CreateManagedNamespace(customer string, newNamespace *models.Namespace) (*models.Namespace, error) {
	toCreate := &controlv1.ManagedNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      swag.StringValue(newNamespace.Name),
			Namespace: customer,
		},
		Spec: controlv1.ManagedNamespaceSpec{
			Description: newNamespace.Description,
		},
	}

	err := k.CRDClient.Create(context.Background(), toCreate)
	if err != nil {
		return nil, err
	}

	return newNamespace, nil
}

func (k *ClientManager) DeleteManagedNamespace(customer, name string) (*models.Namespace, error) {
	managed := controlv1.ManagedNamespace{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: customer,
		},
	}

	err := k.CRDClient.Delete(context.Background(), &managed, client.PropagationPolicy(metav1.DeletePropagationBackground))

	return &models.Namespace{Name: swag.String(name), Customer: swag.String(customer)}, err
}

func (k *ClientManager) UpdateManagedNamespace(customer, name string, managedNamespace *models.Namespace) (*models.Namespace, error) {
	key := client.ObjectKey{
		Namespace: customer,
		Name:      name,
	}

	get := &controlv1.ManagedNamespace{}

	err := k.CRDClient.Get(context.Background(), key, get)

	if err != nil {
		return nil, err
	}

	get.Spec.Description = managedNamespace.Description

	err = k.CRDClient.Update(context.Background(), get)

	return managedNamespace, err
}
