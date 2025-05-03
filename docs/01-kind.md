# Kind

## Introduction to Kind
Kind is a tool for running local Kubernetes clusters using Docker container "nodes". It is primarily designed for testing Kubernetes itself, but may be used for local development or CI. Kind runs a single-node cluster in a single container, or multi-node clusters in multiple containers. It can also be used to run Kubernetes clusters in CI/CD pipelines.

## Installation

Perfect — if your **Docker Engine is already running inside WSL2**, then installing and using **Kind** (Kubernetes in Docker) without Docker Desktop is fully supported and straightforward.

Here’s how to do it **step-by-step** in your WSL2 environment:

---

## ✅ Step 1: Install Kind in WSL2

In your WSL terminal (e.g., Ubuntu):

```bash
curl -Lo ./kind https://kind.sigs.k8s.io/dl/v0.22.0/kind-linux-amd64
chmod +x ./kind
sudo mv ./kind /usr/local/bin/kind
```

> 📌 Replace `v0.22.0` with the latest version if needed — check [https://kind.sigs.k8s.io](https://kind.sigs.k8s.io)

---

## ✅ Step 2: Verify Kind Installation

```bash
kind --version
```

---

## ✅ Step 3: Create a Kind Cluster

Now just run:

```bash
kind create cluster
```

> This uses the Docker daemon that's running **inside WSL**, so no interaction with Docker Desktop is needed.

---

## 🧪 Step 4: Test the Kubernetes Cluster

Check cluster status:

```bash
kubectl cluster-info --context kind-kind
kubectl get nodes
```

If you don’t have `kubectl` installed yet:

```bash
curl -LO "https://dl.k8s.io/release/$(curl -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin/
```

---

## 🛠 Tips

* If you want `kind` to use a specific Docker context, check that with `docker context ls`.
* No changes are needed in Docker config as long as `dockerd` is running and accessible from the same WSL environment.
* You can shut down the cluster with:

```bash
kind delete cluster
```

---

