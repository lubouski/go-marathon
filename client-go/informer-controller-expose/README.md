## Kubernetes client-go Controller build with Informer pattern
Number one job of Kubernetes Controller is to `watch` objects for the desired state and the actual state, then send instructions to make the actual state be more like the desired state. In order to retrieve an object's information, the controller sends a request to Kubernetes API server.
However, repeatedly retrieving information from the API server can become expensive. Thus, in order to get and list objects multiple times in code, Kubernetes developers end up using cache which has already been provided by the client-go library. Additionally, the controller doesn't really want to send requests continuously. It only cares about events when the object has been created, modified or deleted. The client-go library provides the Listwatcher interface that performs an initial list and starts a watch on a particular resource.
### SharedInformer
The informer creates a local cache of a set of resources only used by itself. But, in Kubernetes, there is a bundle of controllers running and caring about multiple kinds of resources. This means that there will be an overlap - one resource is being cared by more than one controller.
In this case, the SharedInformer helps to create a single shared cache among controllers. This means cached resources won't be duplicated and by doing that, the memory overhead of the system is reduced. Besides, each SharedInformer only creates a single watch on the upstream server, regardless of how many downstream consumers are reading events from the informer. This also reduces the load on the upstream server. This is common for the kube-controller-manager which has so many internal controllers.
The SharedInformer has already provided hooks to receive notifications of adding, updating and deleting a particular resource. It also provides convenience functions for accessing shared caches and determining when a cache is primed. This saves us connections against the API server, duplicate serialization costs server-side, duplicate deserialization costs controller-side, and duplicate caching costs controller-side.
```
lw := cache.NewListWatchFromClient(…)
sharedInformer := cache.NewSharedInformer(lw, &api.Pod{}, resyncPeriod)
```
### Workqueue
The SharedInformer can't track where each controller is up to (because it's shared), so the controller must provide its own queuing and retrying mechanism (if required). Hence, most Resource Event Handlers simply place items onto a per-consumer workqueue.
Whenever a resource changes, the Resource Event Handler puts a key to the Workqueue.
