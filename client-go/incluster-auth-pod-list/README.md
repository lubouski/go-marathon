## Tutorial on creating, deploying and providing permissions to client-go library inside cluster
This particular example is based on `outofcluster client-go example` in terms of code base, what is really crucial in this tutorial are steps for building container binary, deploying it to kubernetes cluster and providing sufficient authorization capabilities to run API calls to kubernetes API server.
### Pre-requisites and Versions:
* Kubernetes cluster (kind cluster) version: `v1.24.0`
* go.mod for libraries versions
### Work env and Build container image:
* key note this tutorial uses `podman`, if you have `docker` start the same commands with docker
You could have clean and easy to use `dev` env using `Dockerfile`, to start coding just build a golang container with `--target dev`.
```
$ podman build --target dev . -t go
$ podman run -it -v ${PWD}:/work go sh
$ go version
```
All created files inside dev container (`podman run` and execute shell `sh`), will appear at your current directiry `${PWD}` by means of mounting you current directory to `/work` inside a dev container.
Let's build container image for Kubernetes cluster and upload it to DockerHub registry for example.
```
$ podman build . -t podlister:0.1.1
[1/3] STEP 1/2: FROM golang:1.17 AS dev
[1/3] STEP 2/2: WORKDIR /work
--> Using cache db37b20eb673e4cdbf39069ec22eea470791cbfaf2e236a67fc47c3d70478ccb
--> db37b20eb67
[2/3] STEP 1/4: FROM golang:1.17 AS build
[2/3] STEP 2/4: WORKDIR /app
--> Using cache 431a5515639b229fe08a9f7d20e44c7b418c731e0224b65c846387bbc960decc
--> 431a5515639
[2/3] STEP 3/4: COPY ./app/* /app/
--> Using cache 32cc07822d08f2b6a01bc9e80a2482bc4e6f8b7f1fa22a48046e1b825336d4fa
--> 32cc07822d0
[2/3] STEP 4/4: RUN go get k8s.io/client-go/tools/clientcmd && go get k8s.io/client-go@latest && go build -o podlister
...
$ podman tag podlister:0.1.1 lubowsky/client-go:podlister-v2
$ podman push lubowsky/client-go:podlister-v2
``` 
After multistage build we should get small `alpine` based image with `podlister` binary as a comand to run.
## Creating Kubernetes Deployment, Role and RoleBinding
After we get our container hosted on container registry, we should create deployment for your container image to run it on top of kubernetes cluster.
```
# at repository we already have `deployment.yaml` file
$ kubectl create deployment podlister --image=lubowsky/client-go:podlister-v2 --dry-run=client -oyaml > deployment.yml
$ kubectl apply -f deployment.yml
$ kubectl logs podlister-6b6c6f7d8d-trc9w
error stat /Users/Aliaksandr_Lubouski/.kube/config: no such file or directory building config from flags
error pods is forbidden: User "system:serviceaccount:default:default" cannot list resource "pods" in API group "" in the namespace "default" while listing pods object
```
Every container inside kubernetes has mounted [default certificates](https://github.com/kubernetes/client-go/blob/f10f16e02953ca3d38154c1227efe6612a885267/rest/config.go#L512) for our Kubernetes default SA, which our function `InClusterConfig` is cheking. To overcome this issue we need to create `Role` and `RoleBinding` for our default SA.
```
$ kubectl delete -f deployment.yml
$ kubectl create role lister --resource pods --verb list
role.rbac.authorization.k8s.io/lister created
$ kubectl get sa -n default
NAME      SECRETS   AGE
default   0         47m
# create rolebinding for default SA at default namespace
$ kubectl create rolebinding lister --role lister --serviceaccount default:default
rolebinding.rbac.authorization.k8s.io/lister created
$ kubectl get po
NAME                         READY   STATUS             RESTARTS     AGE
podlister-6b6c6f7d8d-gckjr   0/1     CrashLoopBackOff   1 (6s ago)   10s
# it's ok for deployment to be at STATUS CrashLoopBackOff as we could se from the logs container did his job
$ kubectl logs podlister-6b6c6f7d8d-gckjr
error stat /Users/Aliaksandr_Lubouski/.kube/config: no such file or directory building config from flags
2022-08-05 10:07:01 +0000 UTC
```
From `logs` we could see that container did his job, outputed CreationTime (2022-08-05 10:07:01 +0000 UTC). 
