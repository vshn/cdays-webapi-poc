package kube

import (
	"fmt"

	"github.com/go-openapi/swag"
	"github.com/vshn/cdays-webapi-poc/models"
)

func (k *ClientManager) AddCluster(addCluster *models.Cluster) *models.Cluster {

	cluster, err := newCluster(addCluster)

	if err != nil {
		fmt.Printf("error adding cluster %v", err)
	}

	version, _ := cluster.K8sClient.Discovery().ServerVersion()

	cluster.meta.Version = version.String()

	k.clusters[swag.StringValue(addCluster.Name)] = *cluster

	return addCluster
}

func (k *ClientManager) ListAllClusters() []*models.Cluster {

	returnClusters := make([]*models.Cluster, 0)

	for _, value := range k.clusters {

		meta := value.meta

		meta.Token = ""

		returnClusters = append(returnClusters, meta)
	}

	return returnClusters
}
