//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//
package kube

import (
	"github.com/go-openapi/swag"
	namespacev1 "github.com/vshn/cdays-namespace-poc/pkg/apis"
	"github.com/vshn/cdays-webapi-poc/models"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClientManager holds the k8sclient
type ClientManager struct {
	clusters  map[string]ClusterClients
	K8sClient kubernetes.Interface
	CRDClient client.Client
}

type ClusterClients struct {
	K8sClient kubernetes.Interface
	CRDClient client.Client
	meta      *models.Cluster
}

type ClientProvider interface {
	GetClients(clusterName string) *ClusterClients
}

func NewClientManager(dev bool) (*ClientManager, error) {
	var cfg *rest.Config

	// If devel mode then use configuration flag path.
	if dev {
		// cfg, err = clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
		// if err != nil {
		// 	return nil, fmt.Errorf("could not load configuration: %s", err)
		// }
		cfg = &rest.Config{
			Host:        "https://rancher.vshn.net:443/k8s/clusters/c-gzznj",
			BearerToken: "kubeconfig-u-kuhg44bsif.c-gzznj:z7jkfwxdcr9j9lqg9c2sh64d9rbzrbpvkm6n64vr68tl9hmm4fhd9k",
		}
	}

	clients := clientsFromConfig(cfg)

	return &ClientManager{
		K8sClient: clients.K8sClient,
		CRDClient: clients.CRDClient,
		clusters:  make(map[string]ClusterClients, 0),
	}, nil
}

func newCluster(newCluster *models.Cluster) (*ClusterClients, error) {
	cfg := &rest.Config{
		Host:        swag.StringValue(newCluster.URL),
		BearerToken: newCluster.Token,
	}

	client := clientsFromConfig(cfg)

	client.meta = newCluster

	return client, nil
}

func clientsFromConfig(cfg *rest.Config) *ClusterClients {
	s := runtime.NewScheme()

	namespacev1.AddToSchemes.AddToScheme(s)

	options := client.Options{
		Scheme: s,
	}

	CRDClient, err := client.New(cfg, options)

	k8sCli, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil
	}

	return &ClusterClients{
		CRDClient: CRDClient,
		K8sClient: k8sCli,
	}
}
