//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//
package kube

import (
	"fmt"
	"path/filepath"

	namespace "github.com/vshn/cdays-namespace-poc/pkg/apis"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ClientManager holds the k8sclient
type ClientManager struct {
	// currently only a single client supported
	K8sClient kubernetes.Interface
	CRDClient client.Client
}

func NewClientManager(dev bool) (*ClientManager, error) {
	var err error
	var cfg *rest.Config

	// If devel mode then use configuration flag path.
	if dev {
		cfg, err = clientcmd.BuildConfigFromFlags("", filepath.Join(homedir.HomeDir(), ".kube", "config"))
		if err != nil {
			return nil, fmt.Errorf("could not load configuration: %s", err)
		}
	} else {
		// Some dynamic client generation
	}

	s := runtime.NewScheme()

	namespace.AddToSchemes.AddToScheme(s)

	options := client.Options{
		Scheme: s,
	}

	CRDClient, err := client.New(cfg, options)

	k8sCli, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &ClientManager{
		K8sClient: k8sCli,
		CRDClient: CRDClient,
	}, nil
}
