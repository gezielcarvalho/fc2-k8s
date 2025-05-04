# Working with Pods

Pods are the smallest deployable units in Kubernetes. They can contain one or more containers.
Pods are used to run applications and services in a Kubernetes cluster. They can be managed using `kubectl` commands.

## Creating a Pod

First create a YAML file named `pod-config.yaml` with the following content:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: my-container
      image: nginx
```
This YAML file defines a Pod named `my-pod` that runs an Nginx container.
To create the Pod, run the following command:

```bash

kubectl apply -f pod-config.yaml
```
This command will create the Pod in your Kubernetes cluster.
You can check the status of the Pod using:

```bash
kubectl get pods
```

The output will look like this:

```bash
NAME      READY   STATUS    RESTARTS   AGE
my-pod   1/1     Running   0          1m
```
This indicates that the Pod is running and ready to accept traffic.
## Accessing a Pod
To access the Pod, you can use the following command:

```bash 
kubectl exec -it my-pod -- /bin/bash
```
This command will open a shell inside the Pod, allowing you to interact with the container running inside it.

## Accessing the web application
To access the web application running in the Pod, you can use port forwarding. Run the following command:

```bash
kubectl port-forward my-pod 8080:80
```
This command will forward port 80 of the Pod to port 8080 on your local machine. You can then access the web application by opening your browser and going to `http://localhost:8080/`.

## Delete the pod
To delete a pod, you can use the following command

```bash
kubectl delete pod my-pod
```