run ```kubectl apply -f namespace```
run ```kubectl apply -f deployment```
run ```kubectl apply -f service```


kubectl create namespace argocd
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml


```shell
kubectl apply -f namespace
kubectl apply -f deployment
kubectl apply -f service
```