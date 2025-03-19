@echo off
echo Changing directory to k8s...
cd ../k8s

echo Applying Kubernetes configurations...
kubectl apply -f backend-config.yaml
kubectl apply -f backend-secrets.yaml
kubectl apply -f pvc.yaml
kubectl apply -f mysql-pods.yaml
kubectl apply -f backend-pods.yaml
kubectl apply -f service.yaml

echo Checking status of resources...
echo Pods 
kubectl get pods
echo Service
kubectl get services



echo Deployment completed!
pause