package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/home/ubuntu/.kube/config", "location to your kubeconfig file")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		log.Fatalf("Error building config from flags: %v", err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatalf("Error creating clientset: %v", err)
	}

	namespace := "default"
	secretName := "kindc3-secret"
	secret, err := clientset.CoreV1().Secrets(namespace).Get(context.Background(), secretName, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			log.Fatalf("Secret %s not found in namespace %s", secretName, namespace)
		} else {
			log.Fatalf("Error getting secret: %v", err)
		}
	}

	certificateData := secret.Data["certificate-authority-data"]
	serverData := secret.Data["server"]
	certificateAuthority := string(certificateData)
	server := string(serverData)

	fmt.Printf("Certificate Authority Data: %s\n", certificateAuthority)
	fmt.Printf("Server: %s\n", server)
}
