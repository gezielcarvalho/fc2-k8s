# Creating a ReplicaSet

This document provides an introduction to ReplicaSets and their usage.

## Introduction
A ReplicaSet is a Kubernetes resource that ensures a specified number of pod replicas are running at any given time. It is primarily used to maintain the desired state of applications by automatically replacing failed or terminated pods. ReplicaSets are often used in conjunction with Deployments to provide declarative updates to applications.

## Creating a yaml file
Create a file called `replicaset.yaml` and define the desired state of your ReplicaSet. Below is an example configuration:

```yaml
apiVersion: apps/v1
kind: ReplicaSet
metadata:
    name: example-replicaset
spec:
    replicas: 3
    selector:
        matchLabels:
            app: example-app
    template:
        metadata:
            labels:
                app: example-app
        spec:
            containers:
            - name: example-container
                image: nginx:1.21
```

This configuration ensures that three replicas of the `example-container` pod are always running. 

## Running your ReplicaSet
You can apply this file using the `kubectl apply -f replicaset.yaml` command.

## Verify pods
You can use the command `kubectl get pods` to verify that the desired number of pods are running. You should see three pods with names prefixed by `example-replicaset`. If any pod fails, the ReplicaSet will automatically create a new one to maintain the desired state.

A possible outpup would be similar to the following:

```
NAME                          READY   STATUS    RESTARTS   AGE
example-replicaset-abc123     1/1     Running   0          2m
example-replicaset-def456     1/1     Running   0          2m
example-replicaset-ghi789     1/1     Running   0          2m
```

This output confirms that the ReplicaSet is successfully maintaining the desired number of replicas.

## Verify replicas

To verify the number of replicas managed by the ReplicaSet, you can use the following command:

```bash
kubectl get replicaset
```

This will display information about the ReplicaSet, including the desired, current, and ready replicas. A sample output might look like this:

```
NAME                DESIRED   CURRENT   READY   AGE
example-replicaset  3         3         3       2m
```

In this example, the `DESIRED` column indicates the number of replicas specified in the ReplicaSet configuration, while the `CURRENT` and `READY` columns show the actual number of running and ready replicas, respectively. If all values match, it confirms that the ReplicaSet is functioning as expected.

## Deleting a pod
To test the behavior of the ReplicaSet, you can delete a pod managed by it using the following command:

```bash
kubectl delete pod <pod-name>
```

Replace `<pod-name>` with the name of one of the pods listed in the `kubectl get pods` output. For example:

```bash
kubectl delete pod example-replicaset-abc123
```

After deleting the pod, the ReplicaSet will automatically create a new pod to maintain the desired number of replicas. You can verify this by running:

```bash
kubectl get pods
```

You should see a new pod with a different name created by the ReplicaSet to replace the deleted one. This demonstrates the self-healing capability of ReplicaSets in Kubernetes.
