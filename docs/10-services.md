# Kubernetes Services
Kubernetes Services are an abstraction that defines a logical set of Pods and a policy by which to access them. This is useful because Pods are ephemeral and can be created or destroyed at any time, so a Service provides a stable endpoint for accessing the Pods.

## Types of Services
Kubernetes supports several types of Services, including ClusterIP, NodePort, LoadBalancer, and ExternalName, each serving different use cases and access methods.

### ClusterIP
- **Description**: The default type of Service. It exposes the Service on a cluster-internal IP. This means that the Service is only reachable from within the cluster.
- **Use Case**: Used for internal communication between Pods.
- **Example**:
```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-clusterip-service
spec:
  type: ClusterIP
  selector:
    app: my-app
  ports:
    - protocol: TCP
      port: 80
      targetPort: 8000
```