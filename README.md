# Kubernetes

## Introduction

Kubernetes is an open-source container orchestration platform designed to automate deploying, scaling, and operating application containers. It provides a framework to run distributed systems resiliently, with features like load balancing, self-healing, and automated rollouts and rollbacks.

# Pods
A Pod is the smallest deployable unit in Kubernetes, which can contain one or more containers. Pods share the same network namespace and can communicate with each other using localhost. They also share storage volumes, allowing them to persist data.
Pods are typically used to run a single instance of a service or application, and they can be managed using higher-level abstractions like Deployments, ReplicaSets, and StatefulSets.
Pods can be created and managed using YAML configuration files, which define the desired state of the Pod, including the container image, resource requests and limits, environment variables, and other settings.
Pods can be created using the `kubectl` command-line tool or through the Kubernetes API. They can also be managed using higher-level abstractions like Deployments, ReplicaSets, and StatefulSets.

# Deployments
A Deployment is a higher-level abstraction in Kubernetes that manages the lifecycle of Pods. It provides declarative updates to Pods and ReplicaSets, allowing users to define the desired state of the application and let Kubernetes handle the details of maintaining that state.
Deployments are typically used to manage stateless applications, where the Pods can be easily replaced or scaled without losing data. They provide features like rolling updates, rollbacks, and scaling, making it easy to manage the deployment of applications in a Kubernetes cluster.
Deployments can be created and managed using YAML configuration files, which define the desired state of the Deployment, including the container image, resource requests and limits, environment variables, and other settings.

# ReplicaSets
A ReplicaSet is a Kubernetes resource that ensures a specified number of Pod replicas are running at any given time. It is responsible for maintaining the desired state of the Pods, ensuring that the correct number of replicas are created and running.
ReplicaSets are typically used in conjunction with Deployments, which manage the lifecycle of ReplicaSets and provide declarative updates to Pods. When a Deployment is created, it automatically creates a ReplicaSet to manage the Pods.
ReplicaSets can also be used independently to manage the lifecycle of Pods, providing features like scaling and self-healing.
ReplicaSets can be created and managed using YAML configuration files, which define the desired state of the ReplicaSet, including the number of replicas, container image, resource requests and limits, environment variables, and other settings.

# Kubernetes Architecture
Kubernetes architecture consists of a master node and multiple worker nodes. The master node is responsible for managing the cluster, while the worker nodes run the applications. The key components of Kubernetes architecture include:
- **Master Node**: The control plane that manages the Kubernetes cluster.
- **API Server**: The front-end for the Kubernetes control plane, which exposes the Kubernetes API.
- **Controller Manager**: Manages controllers that regulate the state of the cluster.
- **Scheduler**: Assigns work to worker nodes based on resource availability and requirements.
- **etcd**: A distributed key-value store used for storing cluster data.
- **Kubelet**: An agent that runs on each worker node, ensuring that containers are running in a Pod.
- **Kube Proxy**: Manages network communication between Pods and services.
- **Pods**: The smallest deployable units in Kubernetes, which can contain one or more containers.
- **Services**: An abstraction that defines a logical set of Pods and a policy to access them.
- **Namespaces**: Virtual clusters within a Kubernetes cluster, used for resource isolation.
- **Volumes**: Storage resources that can be shared among containers in a Pod.
- **ConfigMaps**: Used to store non-confidential data in key-value pairs.
- **Secrets**: Used to store sensitive data, such as passwords and tokens.
- **Deployments**: A higher-level abstraction for managing Pods and ReplicaSets.
- **ReplicaSets**: Ensures that a specified number of Pod replicas are running at any given time.
- **StatefulSets**: Manages the deployment and scaling of a set of Pods, providing guarantees about the ordering and uniqueness of these Pods.
- **DaemonSets**: Ensures that all (or some) nodes run a copy of a Pod.
- **Jobs**: A controller that creates one or more Pods and ensures that a specified number of them successfully terminate.
- **CronJobs**: A controller that creates Jobs on a scheduled basis.
- **Ingress**: Manages external access to services, typically HTTP.
- **Network Policies**: Defines rules for how Pods communicate with each other and with other network endpoints.
- **Resource Quotas**: Limits the amount of resources that can be consumed by a namespace.
- **Limit Ranges**: Sets constraints on the resources that can be requested or limited for containers in a namespace.
- **Horizontal Pod Autoscaler**: Automatically scales the number of Pods in a deployment based on observed CPU utilization or other select metrics.
- **Vertical Pod Autoscaler**: Automatically adjusts the resource requests and limits for containers in a Pod based on usage.
- **Cluster Autoscaler**: Automatically adjusts the size of a Kubernetes cluster based on the resource requirements of the Pods.
- **Custom Resource Definitions (CRDs)**: Extends Kubernetes capabilities by allowing users to define their own resource types.
- **Operators**: A method of packaging, deploying, and managing a Kubernetes application.
- **Helm**: A package manager for Kubernetes that simplifies the deployment and management of applications.
- **Kustomize**: A tool for customizing Kubernetes YAML configurations.
- **kubectl**: The command-line tool for interacting with the Kubernetes API server.
- **Kubeadm**: A tool for bootstrapping Kubernetes clusters.
- **Kube-scheduler**: A component of the Kubernetes control plane that assigns Pods to nodes based on resource availability and constraints.
- **Kube-controller-manager**: A component of the Kubernetes control plane that runs controller processes.
- **Kubelet**: An agent that runs on each node in the cluster, ensuring that containers are running in Pods.
- **Kube-proxy**: A network proxy that runs on each node in the cluster, managing network communication between Pods and services.
- **CNI (Container Network Interface)**: A specification for configuring network interfaces in Linux containers.
- **CSI (Container Storage Interface)**: A specification for exposing arbitrary block and file storage systems to containerized workloads on Kubernetes.
- **RBAC (Role-Based Access Control)**: A method for regulating access to resources in Kubernetes based on the roles of individual users or groups.
- **Network Policies**: A specification for controlling the communication between Pods and other network endpoints.
- **Pod Security Policies**: A specification for controlling the security context of Pods.
- **Admission Controllers**: A piece of code that intercepts requests to the Kubernetes API server before persistence of the object, but after the request is authenticated and authorized.
- **Service Accounts**: A special type of account used by Pods to interact with the Kubernetes API.
- **Node Affinity**: A set of rules used by the scheduler to determine which nodes a Pod can be scheduled on based on node labels.
- **Taints and Tolerations**: A mechanism for controlling which Pods can be scheduled on which nodes.
- **Pod Disruption Budgets**: A policy that limits the number of disruptions that can occur to a set of Pods.
- **Resource Requests and Limits**: A way to specify the minimum and maximum amount of resources that a container can use.
- **Affinity and Anti-Affinity**: A set of rules that influence the scheduling of Pods based on labels and other criteria.
- **Service Mesh**: A dedicated infrastructure layer that manages service-to-service communication, providing features like traffic management, security, and observability.
- **Istio**: An open-source service mesh that provides a way to control how microservices share data with one another.
- **Linkerd**: A lightweight service mesh designed for simplicity and performance.
- **Kube-state-metrics**: A service that listens to the Kubernetes API and generates metrics about the state of the objects.
- **Prometheus**: An open-source monitoring and alerting toolkit designed for reliability and scalability.
- **Grafana**: An open-source platform for monitoring and observability, often used with Prometheus for visualizing metrics.
- **Fluentd**: An open-source data collector for unified logging, often used for aggregating logs from multiple sources.
- **ELK Stack (Elasticsearch, Logstash, Kibana)**: A set of tools for searching, analyzing, and visualizing log data in real time.
- **OpenShift**: A Kubernetes-based platform for developing, deploying, and managing applications.
- **Rancher**: An open-source platform for managing Kubernetes clusters, providing a user-friendly interface and additional features.
- **K3s**: A lightweight Kubernetes distribution designed for resource-constrained environments.
- **Minikube**: A tool that creates a local Kubernetes cluster for development and testing purposes.
- **Kind (Kubernetes IN Docker)**: A tool for running local Kubernetes clusters using Docker container "nodes".
- **KubeVirt**: An extension for Kubernetes that allows users to run and manage virtual machines alongside container workloads.
- **Knative**: A Kubernetes-based platform for building, deploying, and managing serverless applications.
- **KEDA (Kubernetes Event-Driven Autoscaling)**: An open-source project that provides event-driven autoscaling for Kubernetes workloads.
- **Kubeflow**: An open-source platform for machine learning on Kubernetes, providing tools for building, training, and deploying ML models.
- **Argo CD**: A declarative, GitOps continuous delivery tool for Kubernetes.
- **Flux**: A set of continuous and progressive delivery solutions for Kubernetes, focusing on GitOps principles.
- **Tekton**: An open-source framework for creating CI/CD systems on Kubernetes.
- **Kustomize**: A tool for customizing Kubernetes YAML configurations, allowing users to create reusable and maintainable configurations.
- **Kubeless**: A serverless framework for Kubernetes, allowing users to deploy functions as a service.
- **OpenFaaS**: An open-source framework for building serverless functions on Kubernetes, providing a simple way to deploy and manage functions.
- **KubeEdge**: An open-source system for extending native containerized application orchestration capabilities to hosts at the edge.
- **K3OS**: A lightweight operating system designed for running Kubernetes clusters, optimized for resource-constrained environments.
- **KubeArmor**: An open-source runtime security enforcement engine for containers and Kubernetes, providing security policies and monitoring.
- **Falco**: An open-source runtime security tool for detecting anomalous activity in applications running on containers.
- **Sysdig**: A cloud-native visibility and security platform for monitoring and securing Kubernetes environments.
- **Aqua Security**: A security platform for protecting containerized applications, providing vulnerability scanning, runtime protection, and compliance.
- **Twistlock**: A cloud-native security platform for protecting containerized applications, providing vulnerability management, compliance, and runtime protection.
- **StackRox**: A Kubernetes-native security platform that provides visibility and control over the security posture of Kubernetes environments.
- **Calico**: A networking and network security solution for containers, providing high-performance networking and policy enforcement.
- **Flannel**: A simple and easy way to configure a layer 3 network fabric designed for Kubernetes.
- **Weave Net**: A networking solution for Kubernetes that provides a simple way to connect and manage containers.
- **Cilium**: A networking and security solution for containers and microservices, providing visibility and control over network traffic.

