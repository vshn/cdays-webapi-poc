//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//
package kube

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/vshn/cdays-webapi-poc/models"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (k *ClientManager) AllManagedNamespaces() []*models.Namespace {
	namespaces := make([]*models.Namespace, 0)

	ns, err := k.K8sClient.CoreV1().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		fmt.Println(err)
		return nil
	}

	for _, space := range ns.Items {
		namespaces = append(namespaces, &models.Namespace{
			Name: swag.String(space.GetName()),
		})
	}

	return namespaces
}
