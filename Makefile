build:
	eval $(minikube docker-env)
	docker build . -t go-auth:latest

initialize_minikube:
	minikube addons enable ingress
	minikube start --memory=4g --cpus=4 --driver=docker --cni=calico --force
	minikube tunnel &

start_local:
	helm upgrade -i go-auth ./.helm -n go-auth
	helm upgrade -i -n go-auth prometheus prometheus-community/prometheus  -f ./.helm/charts/monitoring/values.yaml

uninstall:
	helm uninstall -n go-auth go-auth
	helm uninstall -n go-auth go-auth

local_testrun:	build initialize_minikube start_local