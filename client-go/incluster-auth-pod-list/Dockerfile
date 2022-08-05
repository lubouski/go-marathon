FROM golang:1.17 as dev

WORKDIR /work

FROM golang:1.17 as build

WORKDIR /app
COPY ./app/* /app/
RUN go get k8s.io/client-go/tools/clientcmd && go get k8s.io/client-go@latest && go build -o podlister

FROM alpine as runtime 
RUN apk add gcompat
COPY --from=build /app/podlister /
CMD ./podlister
