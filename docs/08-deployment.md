# Creating a Deployment

A Deployment is a higher-level abstraction that manages ReplicaSets and provides declarative updates to applications. It allows you to define the desired state of your application and automatically manages the underlying ReplicaSets and Pods to achieve that state.

## Writing a Deployment YAML File
# Create a file called `deployment.yaml` and define the desired state of your Deployment. Below is an example configuration:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: example-deployment
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
#
This configuration ensures that three replicas of the `example-container` pod are always running. The Deployment will manage the underlying ReplicaSet and ensure that the desired state is maintained.
## Running your Deployment
You can apply this file using the `kubectl apply -f deployment.yaml` command.
## Verify pods
You can use the command `kubectl get pods` to verify that the desired number of pods are running. You should see three pods with names prefixed by `example-deployment`. If any pod fails, the Deployment will automatically create a new one to maintain the desired state.
A possible output would be similar to the following:
```
NAME                          READY   STATUS    RESTARTS   AGE
example-deployment-abc123     1/1     Running   0          2m
example-deployment-def456     1/1     Running   0          2m
example-deployment-ghi789     1/1     Running   0          2m
```
This output confirms that the Deployment is successfully maintaining the desired number of replicas.