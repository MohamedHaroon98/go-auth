# go-auth


## Stack

1. go
2. mongodb
3. docker
4. helm
5. kubernetes

## TODO

Use GitHub actions, use container registry

### Minikube Installation
> Using Minikube's ingress on WSL2 will timeout your request, thus I follow this instruction
1. `minikube start --memory=4g --cpus=4 --driver=docker --cni=calico` 
2. `minikube tunnel`
3. `curl -X POST -H "Content-Type: application/json" -d '{"username":"your_usernam1e","password":"your_password"}' http://0.0.0.0/register`


`eval $(minikube docker-env) && docker build . -t go-auth:latest`
`helm uninstall -n go-auth go-auth && helm install go-auth ./.helm -n go-auth`