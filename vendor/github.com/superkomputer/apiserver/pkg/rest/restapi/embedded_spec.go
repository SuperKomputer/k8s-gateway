// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for SuperKomputer",
    "title": "RESTful API for Superkomputer",
    "contact": {
      "name": "VMware Inc",
      "email": "fanz@vmware.com"
    },
    "version": "0.1.0"
  },
  "basePath": "/v1",
  "paths": {
    "/clusters": {
      "get": {
        "description": "get a list of all clusters managed by the SuperKomputer",
        "summary": "get a list of all clusters",
        "operationId": "listClusters",
        "responses": {
          "200": {
            "description": "200 response with the list of clusters",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "post": {
        "description": "creates a cluster",
        "summary": "creates a cluster",
        "operationId": "createCluster",
        "parameters": [
          {
            "description": "the information of the cluster to be created",
            "name": "cluster",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "create cluster task has been accepted"
          },
          "409": {
            "$ref": "#/responses/errorClusterNameConflict"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    },
    "/clusters/{clusterId}": {
      "get": {
        "description": "get a cluster details with clusterId",
        "summary": "get a cluster",
        "operationId": "getCluster",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with a cluster details",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "404": {
            "$ref": "#/responses/errorClusterNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "put": {
        "description": "updates a cluster with the given update config",
        "summary": "updates a cluster",
        "operationId": "updateCluster",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be updated",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "description": "the new config of the cluster to be updated",
            "name": "cluster",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "update cluster task has been accepted"
          },
          "304": {
            "description": "no update required"
          },
          "404": {
            "$ref": "#/responses/errorClusterNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "delete": {
        "description": "deletes a cluster with the given name",
        "summary": "deletes a cluster",
        "operationId": "deleteCluster",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be deleted",
            "name": "clusterId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "delete cluster task has been accepted"
          },
          "404": {
            "$ref": "#/responses/errorClusterNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    },
    "/clusters/{clusterId}/account/{username}": {
      "get": {
        "description": "get payable account of a user on cluster",
        "summary": "get payable account of a user on cluster",
        "operationId": "getAccount",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "the name for the user that gathered the account",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with account information",
            "schema": {
              "$ref": "#/definitions/account"
            }
          },
          "404": {
            "$ref": "#/responses/errorAccountNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "post": {
        "description": "creates payable account of a user on cluster",
        "summary": "creates payable account of a user on cluster",
        "operationId": "createAccount",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "the name for the user that gathered the account",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "the payable account of a user on cluster",
            "name": "account",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/account"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "create payable account of a user on cluster has been accepted"
          },
          "404": {
            "$ref": "#/responses/errorClusterNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "delete": {
        "description": "deletes payable account of a user on cluster",
        "summary": "deletes payable account of a user on cluster",
        "operationId": "deleteAccount",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "the name for the user that gathered the account",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "delete account task has been accepted"
          },
          "404": {
            "$ref": "#/responses/errorAccountNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "## create a token which is used to provide identity\n",
        "summary": "log in to the SuperKomputer Service",
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userCredentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful login",
            "schema": {
              "$ref": "#/definitions/jwttoken"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "summary": "get a list of known users",
        "operationId": "listUsers",
        "responses": {
          "200": {
            "description": "200 response with the list of known users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "post": {
        "description": "creates a user",
        "summary": "creates a user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "the username of user to be created",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "create user task has been accepted"
          },
          "409": {
            "$ref": "#/responses/errorUserExists"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "description": "get details information of a user with given username",
        "summary": "get details information of a user with given username",
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "description": "the name of a user",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with the user information",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "$ref": "#/responses/errorUserNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "put": {
        "description": "updates a user with the given update config",
        "summary": "updates a user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the username to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "the new config of the user to be updated",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "update user task has been accepted"
          },
          "304": {
            "description": "no update required"
          },
          "404": {
            "$ref": "#/responses/errorUserNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "delete": {
        "description": "deletes a user with the given name",
        "summary": "deletes a user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the username to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "delete user task has been accepted"
          },
          "404": {
            "$ref": "#/responses/errorUserNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    },
    "/users/{username}/clusters": {
      "get": {
        "summary": "fetch the list of using clusters of user with given username",
        "operationId": "fetchUserClusters",
        "parameters": [
          {
            "type": "string",
            "description": "the name of a user",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with the list of using clusters of user",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    }
  },
  "definitions": {
    "account": {
      "description": "the payable account information",
      "type": "object",
      "required": [
        "username",
        "clusterId",
        "id"
      ],
      "properties": {
        "balance": {
          "type": "string"
        },
        "clusterId": {
          "description": "the cluster name, should be valid for use in dns names",
          "type": "string",
          "maxLength": 63,
          "minLength": 3,
          "pattern": "^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$"
        },
        "id": {
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "cluster": {
      "type": "object",
      "required": [
        "clusterId",
        "status"
      ],
      "properties": {
        "clusterId": {
          "description": "the clusterId, should be valid for use in dns names",
          "type": "string",
          "maxLength": 63,
          "minLength": 3,
          "pattern": "^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$"
        },
        "numMasters": {
          "description": "the number of master created for this kubernetes cluster",
          "type": "integer",
          "format": "int32"
        },
        "numWorkers": {
          "description": "the number of workers created for this kubernetes cluster",
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "description": "the status of the cluster",
          "type": "string"
        },
        "url": {
          "description": "the url to access the kubernetes cluster",
          "type": "string",
          "format": "uri"
        }
      }
    },
    "error": {
      "description": "the default error model for all the error responses coming from the SuperKomputer\n",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "The error code",
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    },
    "jwttoken": {
      "description": "a JWT token to carry identity information",
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "expires_in": {
          "type": "integer",
          "format": "int64"
        },
        "refresh_token": {
          "type": "string"
        },
        "scope": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "token_type": {
          "type": "string",
          "default": "bearer"
        }
      },
      "additionalProperties": true
    },
    "user": {
      "description": "An object for user, username is unique",
      "type": "object",
      "required": [
        "username"
      ],
      "properties": {
        "email": {
          "description": "the email of user",
          "type": "string"
        },
        "firstname": {
          "description": "the first name of user",
          "type": "string"
        },
        "id": {
          "description": "the id of the user",
          "type": "integer",
          "format": "int32"
        },
        "lastname": {
          "description": "the last name of user",
          "type": "string"
        },
        "username": {
          "description": "the name of the user",
          "type": "string"
        }
      }
    },
    "userCredentials": {
      "description": "credentials for logging in to SuperKomputer",
      "type": "object",
      "properties": {
        "password": {
          "type": "string",
          "format": "password"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "errorAccountNotFound": {
      "description": "The account was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorClusterNameConflict": {
      "description": "The provided cluster name already exists",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorClusterNotFound": {
      "description": "The cluster was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorDefault": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorInvalidRequest": {
      "description": "The request is not valid",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorUserExists": {
      "description": "The user with that username already exists",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorUserNotFound": {
      "description": "The user was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for SuperKomputer",
    "title": "RESTful API for Superkomputer",
    "contact": {
      "name": "VMware Inc",
      "email": "fanz@vmware.com"
    },
    "version": "0.1.0"
  },
  "basePath": "/v1",
  "paths": {
    "/clusters": {
      "get": {
        "description": "get a list of all clusters managed by the SuperKomputer",
        "summary": "get a list of all clusters",
        "operationId": "listClusters",
        "responses": {
          "200": {
            "description": "200 response with the list of clusters",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "creates a cluster",
        "summary": "creates a cluster",
        "operationId": "createCluster",
        "parameters": [
          {
            "description": "the information of the cluster to be created",
            "name": "cluster",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "create cluster task has been accepted"
          },
          "409": {
            "description": "The provided cluster name already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/clusters/{clusterId}": {
      "get": {
        "description": "get a cluster details with clusterId",
        "summary": "get a cluster",
        "operationId": "getCluster",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with a cluster details",
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          },
          "404": {
            "description": "The cluster was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "description": "updates a cluster with the given update config",
        "summary": "updates a cluster",
        "operationId": "updateCluster",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be updated",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "description": "the new config of the cluster to be updated",
            "name": "cluster",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/cluster"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "update cluster task has been accepted"
          },
          "304": {
            "description": "no update required"
          },
          "404": {
            "description": "The cluster was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "deletes a cluster with the given name",
        "summary": "deletes a cluster",
        "operationId": "deleteCluster",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be deleted",
            "name": "clusterId",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "delete cluster task has been accepted"
          },
          "404": {
            "description": "The cluster was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/clusters/{clusterId}/account/{username}": {
      "get": {
        "description": "get payable account of a user on cluster",
        "summary": "get payable account of a user on cluster",
        "operationId": "getAccount",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "the name for the user that gathered the account",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with account information",
            "schema": {
              "$ref": "#/definitions/account"
            }
          },
          "404": {
            "description": "The account was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "creates payable account of a user on cluster",
        "summary": "creates payable account of a user on cluster",
        "operationId": "createAccount",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "the name for the user that gathered the account",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "the payable account of a user on cluster",
            "name": "account",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/account"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "create payable account of a user on cluster has been accepted"
          },
          "404": {
            "description": "The cluster was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "deletes payable account of a user on cluster",
        "summary": "deletes payable account of a user on cluster",
        "operationId": "deleteAccount",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the clusterId to be queried",
            "name": "clusterId",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "description": "the name for the user that gathered the account",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "delete account task has been accepted"
          },
          "404": {
            "description": "The account was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/login": {
      "post": {
        "description": "## create a token which is used to provide identity\n",
        "summary": "log in to the SuperKomputer Service",
        "operationId": "login",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/userCredentials"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "successful login",
            "schema": {
              "$ref": "#/definitions/jwttoken"
            }
          }
        }
      }
    },
    "/users": {
      "get": {
        "summary": "get a list of known users",
        "operationId": "listUsers",
        "responses": {
          "200": {
            "description": "200 response with the list of known users",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/user"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "post": {
        "description": "creates a user",
        "summary": "creates a user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "the username of user to be created",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "create user task has been accepted"
          },
          "409": {
            "description": "The user with that username already exists",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{username}": {
      "get": {
        "description": "get details information of a user with given username",
        "summary": "get details information of a user with given username",
        "operationId": "getUser",
        "parameters": [
          {
            "type": "string",
            "description": "the name of a user",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with the user information",
            "schema": {
              "$ref": "#/definitions/user"
            }
          },
          "404": {
            "description": "The user was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "put": {
        "description": "updates a user with the given update config",
        "summary": "updates a user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the username to be updated",
            "name": "username",
            "in": "path",
            "required": true
          },
          {
            "description": "the new config of the user to be updated",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/user"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "update user task has been accepted"
          },
          "304": {
            "description": "no update required"
          },
          "404": {
            "description": "The user was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "deletes a user with the given name",
        "summary": "deletes a user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "x-nullable": false,
            "description": "the username to be deleted",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "delete user task has been accepted"
          },
          "404": {
            "description": "The user was not found",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/users/{username}/clusters": {
      "get": {
        "summary": "fetch the list of using clusters of user with given username",
        "operationId": "fetchUserClusters",
        "parameters": [
          {
            "type": "string",
            "description": "the name of a user",
            "name": "username",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with the list of using clusters of user",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "description": "Error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "account": {
      "description": "the payable account information",
      "type": "object",
      "required": [
        "username",
        "clusterId",
        "id"
      ],
      "properties": {
        "balance": {
          "type": "string"
        },
        "clusterId": {
          "description": "the cluster name, should be valid for use in dns names",
          "type": "string",
          "maxLength": 63,
          "minLength": 3,
          "pattern": "^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$"
        },
        "id": {
          "type": "integer",
          "format": "int32"
        },
        "username": {
          "type": "string"
        }
      }
    },
    "cluster": {
      "type": "object",
      "required": [
        "clusterId",
        "status"
      ],
      "properties": {
        "clusterId": {
          "description": "the clusterId, should be valid for use in dns names",
          "type": "string",
          "maxLength": 63,
          "minLength": 3,
          "pattern": "^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$"
        },
        "numMasters": {
          "description": "the number of master created for this kubernetes cluster",
          "type": "integer",
          "format": "int32"
        },
        "numWorkers": {
          "description": "the number of workers created for this kubernetes cluster",
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "description": "the status of the cluster",
          "type": "string"
        },
        "url": {
          "description": "the url to access the kubernetes cluster",
          "type": "string",
          "format": "uri"
        }
      }
    },
    "error": {
      "description": "the default error model for all the error responses coming from the SuperKomputer\n",
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "description": "The error code",
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    },
    "jwttoken": {
      "description": "a JWT token to carry identity information",
      "type": "object",
      "properties": {
        "access_token": {
          "type": "string"
        },
        "expires_in": {
          "type": "integer",
          "format": "int64"
        },
        "refresh_token": {
          "type": "string"
        },
        "scope": {
          "type": "array",
          "items": {
            "type": "string"
          }
        },
        "token_type": {
          "type": "string",
          "default": "bearer"
        }
      },
      "additionalProperties": true
    },
    "user": {
      "description": "An object for user, username is unique",
      "type": "object",
      "required": [
        "username"
      ],
      "properties": {
        "email": {
          "description": "the email of user",
          "type": "string"
        },
        "firstname": {
          "description": "the first name of user",
          "type": "string"
        },
        "id": {
          "description": "the id of the user",
          "type": "integer",
          "format": "int32"
        },
        "lastname": {
          "description": "the last name of user",
          "type": "string"
        },
        "username": {
          "description": "the name of the user",
          "type": "string"
        }
      }
    },
    "userCredentials": {
      "description": "credentials for logging in to SuperKomputer",
      "type": "object",
      "properties": {
        "password": {
          "type": "string",
          "format": "password"
        },
        "username": {
          "type": "string"
        }
      }
    }
  },
  "responses": {
    "errorAccountNotFound": {
      "description": "The account was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorClusterNameConflict": {
      "description": "The provided cluster name already exists",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorClusterNotFound": {
      "description": "The cluster was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorDefault": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorInvalidRequest": {
      "description": "The request is not valid",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorUserExists": {
      "description": "The user with that username already exists",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorUserNotFound": {
      "description": "The user was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  }
}`))
}
