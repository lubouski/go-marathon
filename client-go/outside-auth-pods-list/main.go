package main

import (
	"fmt"
	"context"
	"flag"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/kubernetes"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	// get kubeconfig from file
	kubeconfig := flag.String("kubeconfig","/Users/Aliaksandr_Lubouski/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)	
	if err != nil {
		panic(err.Error())
	}
	// inializing clientset to communicate with kubernetes objects
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// access corev1 pod object List method
	pods, err := clientset.CoreV1().Pods("default").List(context.Background(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// List method returns result *v1.PodList from api/core/v1/types.go which has field Items []Pod
	// Pod type has inherited fields as metav1.ObjectMeta from apimachinery/pkg/apis/meta/v1/types.go 
	for _, pod := range pods.Items {
		// ObjectMeta has a lot of fields one of them is `CreationTimestamp` which we could use
		fmt.Println(pod.CreationTimestamp)
	}
}
