//
// Copyright (c) 2019, VSHN AG, info@vshn.ch
// Licensed under "BSD 3-Clause". See LICENSE file.
//
//
package kube

import (
	"fmt"
	"path/filepath"

	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// ClientManager holds the k8sclient
type ClientManager struct {
	// currently only a single client supported
	K8sClient kubernetes.Interface
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

	k8sCli, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		return nil, err
	}

	return &ClientManager{
		K8sClient: k8sCli,
	}, nil
}
