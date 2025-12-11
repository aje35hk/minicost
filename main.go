package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

// func BuildKubeConfigPath() (string, error) {
// 	home, err := os.UserHomeDir()

// 	if err != nil {
// 		return "", err
// 	}

// 	kubeconfig := filepath.Join(home, ".kube", "config")
// 	return kubeconfig, nil
// }

func main() {

	home, err := os.UserHomeDir()

	if err != nil {
		log.Fatalf("error getting user home directory: %v", err)
	}

	kubeconfig := filepath.Join(home, ".kube", "config")

	if err != nil {
		log.Fatalf("error building config %v", err)
	}

	// log.Printf("kubeconfig path: %s", kubeconfig)

	// Load the kubeconfig file or use in-cluster config

	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)

	if err != nil {
		log.Fatalf("error loading kubeconfig: %v", err)
	}

	// log.Printf("Successfully loaded kubeconfig: %+v", config)

	// Create the clientset

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		log.Fatalf("error creating Kubernetes client: %v", err)
	}

	// log.Printf("Successfully created Kubernetes client: %+v", clientset)

	// List Pods in the "default" namespace

	listpods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		log.Fatalf("error listing pods: %v", err)
	}

	for _, d := range listpods.Items {
		fmt.Println("Pod Name: " + d.Name)
	}
}
