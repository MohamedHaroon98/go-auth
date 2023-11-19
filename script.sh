eval $(minikube docker-env) && docker build . -t go-auth:latest
helm uninstall -n go-auth go-auth
helm install go-auth ./.helm -n go-auth
sleep 10
minikube service -n go-auth go-auth-go-auth --url

# need to use GitHub Actions instead
