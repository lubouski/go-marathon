package main

import (
	"fmt"
	"context"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	kubeconfig := flag.String("kubeconfig","/Users/Aliaksandr_Lubouski/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)	
	if err != nil {
		fmt.Printf("error %s building config from flags\n", err.Error())
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s building inCluster config\n", err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, creating clientset\n", err.Error())
	}

	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Printf("error %s while listing pods object\n", err.Error()) 
	}

	for _, pod := range pods.Items {
		fmt.Println(pod.CreationTimestamp)
	}
}
