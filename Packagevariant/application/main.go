package main

import (
	"context"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	targetKubeconfig := flag.String("target-cluster", "/home/ubuntu/.kube/sourceclusterkubeconfig.yaml", "location for kubeconfig file for source cluster")
	flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *targetKubeconfig)
	if err != nil {
		fmt.Printf("Error building target cluster kubeconfig: %v\n", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating Kubernetes client for target cluster: %v\n", err)
		os.Exit(1)
	}

	namespace := "default"
	secretName := "target"
	secretData, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		log.Fatalf("Error getting secret: %s", err.Error())
	}

	kubeconfigData, ok := secretData.Data["targetclusterkubeconfig.txt"]
	if !ok {
		log.Fatalf("kubeconfig data not found in the secret")
	}

	decodedKubeconfig, err := base64.StdEncoding.DecodeString(string(kubeconfigData))
	if err != nil {
		log.Fatalf("Error decoding kubeconfig data: %s", err.Error())
	}

	targetConfig, err := clientcmd.RESTConfigFromKubeConfig(decodedKubeconfig)
	if err != nil {
		fmt.Printf("Error creating REST config from decoded kubeconfig: %v\n", err)
		os.Exit(1)
	}

	targetClientset, err := kubernetes.NewForConfig(targetConfig)
	if err != nil {
		fmt.Printf("Error creating Kubernetes client for target cluster using secret: %v\n", err)
		os.Exit(1)
	}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "exmaple-pod",
			Namespace: "default",
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "nginx",
					Image: "nginx",
				},
			},
		},
	}

	_, err = targetClientset.CoreV1().Pods("default").Create(context.TODO(), pod, metav1.CreateOptions{})
	if err != nil {
		fmt.Printf("Error deploying pod: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Pod deployed successfully on the target cluster")
}
