## Client-go library example for outofcluster API calls
In this particular example would be examinated most simple code to interact with the Kuberntes cluster via `client-go` library. What we would do it's quering `pods` at `default` namespace to get creation timestamps for all of the pods, why this particular example is handy because to get the same result with `kubectl get po` it's kinds tricky.
### Pre-requisites and Versions:
* Kubernetes cluster (kind cluster) version: `v1.24.0`
* go.mod for libraries versions
### Use case:
Lets execute first two containers at our cluster, then run `kubectl` command and after run out `main.go` script.
```
$ kubectl run  nginx-1 --image=nginx --restart=Never
$ kubectl run  nginx-2 --image=nginx --restart=Never
$ kubectl get po -owide
NAME      READY   STATUS    RESTARTS   AGE   IP           NODE                 NOMINATED NODE   READINESS GATES
nginx     1/1     Running   0          53m   10.244.0.5   kind-control-plane   <none>           <none>
nginx-2   1/1     Running   0          47m   10.244.0.6   kind-control-plane   <none>           <none>
# it's possible to get exact creation timestamp but it a bit complicated with `kubectl`, `jq` and `bash loops` :) 
$ go run main.go
2022-08-02 13:52:31 +0200 CEST
2022-08-02 13:58:44 +0200 CEST
```
### Walk-through code:
The most interesting part is code for accesing kubernetes `core/v1` api's:
```
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
```
With a help of `clientset` we could access main REST endpoints for kubernetes `pod kind` [kubernetes-core-v1-pod](https://github.com/kubernetes/client-go/blob/master/kubernetes/typed/core/v1/pod.go), method which we are interested is `List` which return result result v1.PodList from [api/core/v1/types.go](https://github.com/kubernetes/api/blob/f5e1938afa507ee250e4a0a7a45cd2498e3750d3/core/v1/types.go#L3968) which has field Items []Pod type. Pod type has inherited fields as `metav1.ObjectMeta` from [apimachinery/pkg/apis/meta/v1/types.go](https://github.com/kubernetes/apimachinery/blob/899984fb2d224c4a86ca7906e38754a3a62f8c41/pkg/apis/meta/v1/types.go#L111), `ObjectMeta` has a lot of fields one of them is `CreationTimestamp` which we could use at for range.
