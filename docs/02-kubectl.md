# Kubectl

# Introduction
Kubectl is the command-line tool for interacting with Kubernetes clusters. It allows you to deploy applications, inspect and manage cluster resources, and view logs.

# Installation
## 🐳 Installing kubectl on WSL

---
To install `kubectl` in your **WSL2 environment (e.g., Ubuntu)**, follow the official and safe method using the Kubernetes release binaries.

---

## ✅ Step-by-Step: Install `kubectl` in WSL2

### 1. Download the Latest Stable Version


```bash
curl -LO "https://dl.k8s.io/release/v1.29.3/bin/linux/amd64/kubectl"
```


---

### 2. Make It Executable

```bash
chmod +x kubectl
```

---

### 3. Move It to a Directory in Your PATH

```bash
sudo mv kubectl /usr/local/bin/
```

---

### 4. Verify Installation

```bash
kubectl version --client
```

You should see something like:

```
Client Version: version.Info{Major:"1", Minor:"29", GitVersion:"v1.29.3", ...}
```


---

You're now ready to use `kubectl` to manage your Kubernetes clusters.
