package handlers

import (
	"github.com/fanzhangio/superkomputer/pkg/rest/models"
	"github.com/go-openapi/errors"
)

var clusterLastID int32

// CachedClusterList is cached user-cluster mapping
var CachedClusterList = make(map[string][]*models.Cluster)

// CachedClusters is cached clusters list
var CachedClusters = make(map[string]*models.Cluster)

// ListClusters gets a list of all clusters
func ListClusters() (result []*models.Cluster) {
	result = make([]*models.Cluster, 0)
	for _, v := range CachedClusterList {
		result = append(result, v...)
	}
	return
}

// CreateCluster createss a cluster by given cluster information
func CreateCluster(c *models.Cluster) error {
	if c == nil {
		return errors.New(402, "The request is not valid. Cluster playload must be provided for create cluster operation.")
	}
	// TODO Auth middleware
	if _, ok := CachedClusters[*c.ClusterID]; ok {
		return errors.New(409, "The cluster with that clusterId already exists")
	}
	CachedClusters[*c.ClusterID] = &models.Cluster{
		ClusterID:  c.ClusterID,
		NumMasters: c.NumMasters,
		NumWorkers: c.NumWorkers,
		Status:     c.Status,
		URL:        c.URL,
	}
	return nil
}

// GetCluster gets a cluster details with clusterId
func GetCluster(clusterID string) (*models.Cluster, error) {
	// TODO Auth middleware
	if cluster, ok := CachedClusters[clusterID]; ok {
		return cluster, nil
	}
	return nil, errors.New(404, "The cluster was not found.")
}

// UpdateCluster updates a cluster with the given update config and clusterId
func UpdateCluster(clusterID string, c *models.Cluster) error {
	// TODO Auth middleware
	if _, ok := CachedClusters[clusterID]; ok {
		CachedClusters[clusterID] = c
		return nil
	}
	return errors.New(404, "The user was not found.")
}

// DeleteCluster deletes a cluster with the given name
func DeleteCluster(clusterID string) error {
	// TODO Auth middleware
	if _, ok := CachedClusters[clusterID]; ok {
		delete(CachedClusters, clusterID)
		return nil
	}
	return errors.New(402, "The cluster was not found.")
}
