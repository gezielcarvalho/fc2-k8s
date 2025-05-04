# Upgrading Apps in a ReplicaSet

Upgrading applications in a ReplicaSet is a common task in Kubernetes. This process typically involves updating the container image version or modifying the configuration of the pods managed by the ReplicaSet. Below are the steps to upgrade an application in a ReplicaSet.

## Step 1: Update the Hello World Application 
In this example, we will be using a simple Go application that serves a "Hello E4Devs" message. The application is running in a container managed by the ReplicaSet.

## Build and Push the New Image
1. Update the `main.go` file to change the message from "Hello E4Devs" to "Hello E4Devs - Updated Version".
2. Build the new Docker image and push it to your container registry. For example:
   ```bash
   docker build -t your-registry/hello-world:2.0 .
   docker push your-registry/hello-world:2.0
   ```
3. Update the ReplicaSet configuration file (e.g., `replicaset.yaml`) to use the new image version. Change the image line from:
   ```yaml
   image: your-registry/hello-world:1.0
   ```
   to:
   ```yaml
   image: your-registry/hello-world:2.0
   ```
4. Apply the updated ReplicaSet configuration:
   ```bash
   kubectl apply -f replicaset.yaml
   ```
RMK. A ReplicaSet is not designed to trigger updates when the image changes. It only manages the number of pods with a matching template, and does not support rolling updates or image redeploys on changes.

5. Force the ReplicaSet to update by removing the old pods:
   ```bash
   kubectl delete pod -l app=hello-world
   ```
   This will cause the ReplicaSet to create new pods with the updated image.


## Conclusion
Upgrading applications in a ReplicaSet is a straightforward process in Kubernetes. By following the steps outlined above, you can ensure that your applications are running the latest versions with minimal downtime. Always remember to test the new version thoroughly before removing the old ReplicaSet to avoid any disruptions in service.
