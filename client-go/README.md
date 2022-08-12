## Kubernetes client-go library examples
Client-go is a golang client to talk to a kubernetes cluster API. An example at the root of the folder could be used to explain how client-go could works with GCP kubernetes cluster authentification.
### Next steps:
More precise examples would be stored at the particualr folders. 
### How to get it:
```
$ go get k8s.io/client-go@latest
```
### Kubernetes API concepts
Why Kubernetes controller has chosen to use events (i.e state changes) to drive its logic. There are two principled options to detect state change (the event itself):
* Edge-driven triggers (At the point in time the state change occurs, a handler is triggered - for example, from no pod to pod running).
* Level-driven triggers (The state is checked at regular intervals and if certain conditions are met (for example, pod running), then a handler is triggered).
The latter is a form of polling. It does not scale well with the number of objects, and the latency of controllers noticing changes depends on the interval of poling and how fast the API server can answer.
#### Optimistic Concurrency
In a nutshell that if and when the API server detects concurrent write attempts, it rejects the `latter` of the two write operations. It is then up to the client (controller, scheduler, kubectl, etc) to handle a conflict and potentially retry the write operation.
#### Kubernetes API - GroupVersionResource and GroupVersionKind
Resources are always part of an API group and a version, collectively referred to as `GroupVersionResource` (or GVR). A GVR uniquely defines an HTTP path. A concrete path, for example, in the default namespace would be `/api/batch/v1/namespaces/default/jobs`.
Similarly to GVRs, each `kind` lives in an API group, is versioned, and is identified via a `GroupVersionKind` (GVK).
GVKs and GVRs are related. GVKs are served under HTTp paths identified by GVRs. The process of mapping a GVK to a GVR is called `REST mapping`.
