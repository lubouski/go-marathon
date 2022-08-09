package main

import (
	"flag"
	"fmt"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	kubeconfig := flag.String("kubeconfig", "/Users/Aliaksandr_Lubouski/.kube/config", "location to your kubeconfig file")
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		fmt.Printf("error %s building config from flags\n", err.Error())
		// https://github.com/kubernetes/client-go/blob/f10f16e02953ca3d38154c1227efe6612a885267/rest/config.go#L117
		config, err = rest.InClusterConfig()
		if err != nil {
			fmt.Printf("error %s building inCluster config\n", err.Error())
		}
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("error %s, creating clientset\n", err.Error())
	}

	ch := make(chan struct{})
	// creating informer for particular namespace (eksposer)
	// ListOptions could help with restrict the list of returned objects
	// https://github.com/kubernetes/apimachinery/blob/899984fb2d224c4a86ca7906e38754a3a62f8c41/pkg/apis/meta/v1/types.go#L322
	informers := informers.NewFilteredSharedInformerFactory(clientset, 10*time.Minute, "eksposer", func(opts *metav1.ListOptions) {})
	c := newController(clientset, informers.Apps().V1().Deployments())
	informers.Start(ch)
	c.run(ch)
	fmt.Println(informers)
}
