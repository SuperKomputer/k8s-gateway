package handlers

import (
	"github.com/fanzhangio/superkomputer/pkg/rest/models"
	"github.com/go-openapi/errors"
)

var accountLastID int32

// CachedClusterAccounts is cached cluster-accounts mapping
var CachedClusterAccounts = make(map[string][]*models.Account)

// GetAccount gets payable account of a user on cluster
func GetAccount(clusterID, username string) (*models.Account, error) {
	// TODO Auth middleware
	cluster, ok := CachedClusters[clusterID]
	if !ok {
		return nil, errors.New(404, "The cluster was not found.")
	}
	for _, account := range CachedClusterAccounts[*cluster.ClusterID] {
		if *account.Username == username && *account.ClusterID == clusterID {
			return account, nil
		}
	}
	return nil, errors.New(404, "The account was not found.")
}

// DeleteAccount deletes payable account of a user on cluster
func DeleteAccount(clusterID, username string) error {
	// TODO Auth middleware
	cluster, ok := CachedClusters[clusterID]
	if !ok {
		return errors.New(404, "The cluster was not found.")
	}
	accounts := CachedClusterAccounts[*cluster.ClusterID]
	for i, ac := range accounts {
		if *ac.Username == username && *ac.ClusterID == clusterID {
			accounts = append(accounts[:i], accounts[i+1:]...)
			return nil
		}
	}
	return errors.New(404, "The account was not found.")
}

// CreateAccount creates payable account of a user on cluster
func CreateAccount(clusterID, username string, a *models.Account) error {
	// TODO Auth middleware
	cluster, ok := CachedClusters[clusterID]
	if !ok {
		return errors.New(404, "The cluster was not found.")
	}
	accounts := CachedClusterAccounts[*cluster.ClusterID]
	for _, ac := range accounts {
		if *ac.Username == username && *ac.ClusterID == clusterID {
			return errors.New(409, "The account with clusterId for user already exists")
		}
	}
	newID := NewID(&accountLastID)
	accounts = append(accounts, &models.Account{
		Balance:   a.Balance,
		ClusterID: &clusterID,
		ID:        &newID,
		Username:  a.Username,
	})
	return nil
}
