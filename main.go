/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Note: the example only works with the code within the same release/branch.
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	apiclient "github.com/superkomputer/apiclient/client"
	"github.com/superkomputer/apiclient/client/operations"
	"github.com/superkomputer/apiclient/models"

	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

const namespacePrefix = "sk-ns"
const envExternalIps = "K8S_EXTERNALIPS"
const envK8sAPI = "K8S_API"
const envK8sUsername = "K8S_USERNAME"
const envK8sPassword = "K8S_PASSWORD"
const envK8sClusterID = "K8S_CLUSTERID"
const envSkAPI = "SK_API"
const envSkUserid = "SK_USERID"

type skConfig struct {
	ExternalIps  []string
	API          string
	UserID       string
	Namespaces   []string
	K8sAPI       string
	K8sUsername  string
	K8sPassword  string
	K8sClusterID string
}

func main() {
	var kubeconfig *string
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	}
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		panic(err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	c := &skConfig{}
	ns, err := validateNamespaces(clientset)
	if err != nil {
		panic(err)
	}
	c.Namespaces = ns
	fmt.Println(fmt.Sprintf("%d namespaces have been validated", len(ns)))

	c.ExternalIps = getExternalIps()
	c.API = getAPI()
	c.UserID = getUserID()
	c.K8sAPI = getK8sAPI()
	c.K8sUsername = getK8sUsername()
	c.K8sPassword = getK8sPassword()
	c.K8sClusterID = getK8sClusterID()
	//namespaces
	//username
	//password
	fmt.Println(fmt.Sprintf("configuration is set as %+v", c))

	transport := httptransport.New(getAPI(), "/v1", []string{"http"})
	client := apiclient.New(transport, strfmt.Default)
	// resp, err := client.Operations.GetCluster(operations.NewGetClusterParams())
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Printf("%#v\n", resp.Payload)

	status := "OK"
	cp := operations.NewCreateClusterParams()
	cp.Cluster = &models.Cluster{
		NumMasters: 1,
		NumWorkers: 3,
		Status:     &status,
		URL:        strfmt.URI(c.K8sAPI),
		ClusterID:  &c.K8sClusterID,
	}

	client.Operations.CreateCluster(cp)
}

func getExternalIps() []string {
	ips := os.Getenv(envExternalIps)
	return strings.Split(ips, ",")
}

func getAPI() string {
	return os.Getenv(envSkAPI)
}

func getK8sAPI() string {
	return os.Getenv(envK8sAPI)
}

func getK8sUsername() string {
	return os.Getenv(envK8sUsername)
}

func getK8sPassword() string {
	return os.Getenv(envK8sPassword)
}

func getK8sClusterID() string {
	return os.Getenv(envK8sClusterID)
}

func getUserID() string {
	return os.Getenv(envSkUserid)
}

func getNamespaces(clientset *kubernetes.Clientset) ([]v1.Namespace, error) {
	ns, err := clientset.Core().Namespaces().List(metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	return ns.Items, err
}

func validateNamespaces(clientset *kubernetes.Clientset) ([]string, error) {
	ns, err := getNamespaces(clientset)
	if err != nil {
		return nil, err
	}

	mapNs := make(map[string]bool)
	for _, n := range ns {
		mapNs[n.Name] = true
	}

	arrSkns := make([]string, 0)
	for i := 1; i <= 10; i++ {
		ia := strconv.Itoa(i)
		name := namespacePrefix + ia
		if ok, _ := mapNs[name]; !ok {
			if _, err := createNamespace(clientset, name); err != nil {
				return nil, err
			}
		}
		arrSkns = append(arrSkns, name)
	}
	return arrSkns, nil
}

func createNamespace(clientset *kubernetes.Clientset, name string) (*v1.Namespace, error) {
	fmt.Println(fmt.Sprintf("creating namespace %s", name))
	return clientset.Core().Namespaces().Create(
		&v1.Namespace{
			ObjectMeta: metav1.ObjectMeta{
				Name: name,
			},
		},
	)
}
