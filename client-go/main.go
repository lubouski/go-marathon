package main

import (
  "context"
  "fmt"

  metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
  "k8s.io/client-go/tools/clientcmd"
  "k8s.io/client-go/kubernetes"
  _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
)

func main() {
  rules := clientcmd.NewDefaultClientConfigLoadingRules()
  kubeconfig := clientcmd.NewNonInteractiveDeferredLoadingClientConfig(rules, &clientcmd.ConfigOverrides{})
  config, err := kubeconfig.ClientConfig()
  if err != nil {
    panic(err)
  }
  clientset := kubernetes.NewForConfigOrDie(config)

  nodeList, err := clientset.CoreV1().Nodes().List(context.Background(), metav1.ListOptions{})
  if err != nil {
    panic(err)
  }
  for _,n := range nodeList.Items {
    fmt.Println(n.Name)
  }
}
