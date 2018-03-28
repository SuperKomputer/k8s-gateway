package handlers

import (
	"github.com/go-openapi/errors"

	"github.com/fanzhangio/superkomputer/pkg/rest/models"
)

var userLastID int32

// CachedUserList is cached user list
var CachedUserList = make(map[string]*models.User)

// ListUsers impl
func ListUsers() (result []*models.User) {
	result = make([]*models.User, 0)
	for _, item := range CachedUserList {
		result = append(result, item)
	}
	return
}

// CreateUser impl
func CreateUser(u *models.User) error {
	if u == nil {
		return errors.New(402, "The request is not valid. User playload must be provided for create user operation.")
	}
	// TODO Auth middleware

	if _, ok := CachedUserList[*u.Username]; ok {
		return errors.New(409, "The user with that username already exists")
	}
	CachedUserList[*u.Username] = &models.User{
		Email:     u.Email,
		Firstname: u.Firstname,
		ID:        NewID(&userLastID),
		Lastname:  u.Lastname,
		Username:  u.Username,
	}
	return nil
}

// GetUser impl
func GetUser(username string) (*models.User, error) {
	// TODO Auth middleware
	if user, ok := CachedUserList[username]; ok {
		return user, nil
	}
	return nil, errors.New(402, "The user was not found.")
}

// DeleteUser impl
func DeleteUser(username string) error {
	// TODO Auth middleware
	if _, ok := CachedUserList[username]; ok {
		delete(CachedUserList, username)
		return nil
	}
	return errors.New(402, "The user was not found.")
}

// UpdateUser impl
func UpdateUser(username string, u *models.User) error {
	// TODO Auth middleware
	if _, ok := CachedUserList[username]; ok {
		CachedUserList[username] = u
		return nil
	}
	return errors.New(402, "The user was not found.")
}

// FetchUserClusters impl
func FetchUserClusters(username string) (result []*models.Cluster, err error) {
	result = make([]*models.Cluster, 0)
	user, ok := CachedUserList[username]
	if !ok {
		return nil, errors.New(402, "The user was not found.")
	}
	clusterlist, ok := CachedClusterList[*user.Username]
	if !ok {
		return nil, errors.New(404, "The cluster list was not found.")
	}
	return clusterlist, nil
}
