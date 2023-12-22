**Information about task**

*This guide walks you through the process of setting up a Kubernetes deployment, including creating a pod, attaching a service, using ConfigMaps, and setting up a Network Attachment Definition (NAD) with macvlan.


### I. Kubernates
## Prerequisites
Before you begin, ensure you have a Kubernetes cluster running and `kubectl` installed and configured to interact with your cluster.
```
 git clone https://github.com/Rameshwar-Kanade/Niphio_study/tree/main/kubernates
```
# Applying Kubernetes YAML Files
## Steps to Create a YAML File

### 1. Create Your YAML File

```
kubectl create -f pod.yaml
```
```
kubectl create -f deployment.yaml
```
```
kubectl create -f services.yaml
```
```
kubectl create -f configmap.yaml
```
```
kubectl create -f nad.yaml
```
## Steps to Apply a YAML File

### 2. Apply Your YAML File
```
kubectl apply -f pod.yaml
```
```
kubectl apply -f deployment.yaml
```
```
kubectl apply -f services.yaml
```
```
kubectl apply -f configmap.yaml
```
```
kubectl apply -f nad.yaml
```
## Steps to Verify a YAML File

### 3. Verify Your YAML File
```
kubectl get pod
```
```
kubectl get deployments
```
```
kubectl get svc
```
```
kubectl get cm
```
```
kubectl get network-attachment-definitions
```

### II. Helm Charts
```
git clone https://github.com/Rameshwar-Kanade/Niphio_study/tree/main/Helmchart
```
## Prerequisites
Before proceeding, ensure you have the following prerequisites met:

1. **Kubernetes Cluster**: A running Kubernetes cluster. This can be a local cluster (like Minikube or Kind) or a cloud-based cluster.

2. **Helm Installed**: Helm, the package manager for Kubernetes, must be installed on your machine. Helm simplifies deploying and managing Kubernetes applications.

3. **Kubectl Installed**: The Kubernetes command-line tool, kubectl, should be installed and configured to communicate with your cluster.

4. **Network Plugin for NAD**: Ensure that a compatible CNI plugin (like Multus) that supports Network Attachment Definitions (NAD) is installed in your cluster.
   ```
   sudo wget -O /opt/cni/bin/macvlan https://github.com/containernetworking/plugins/releases/download/v0.9.1/cni-plugins-linux-amd64-v0.9.1.tgz
   ```
   ```
   sudo tar -zxvf /opt/cni/bin/macvlan -C /opt/cni/bin/
   ```
   ```
   sudo /opt/cni/bin/macvlan --version
   ```
    ```
   ls /opt/cni/bin

5. **Basic Understanding of Kubernetes Concepts**: Familiarity with basic Kubernetes concepts like Pods, Deployments, Services, and ConfigMaps.

This command creates a chart directory along with the common files and directories used in a chart.
For example, 'helm create mychart' will create a directory structure that looks something like this:

```
mychart/ 
├── Chart.yaml          # A YAML file containing information about the chart
├── values.yaml         # The default configuration values for this chart (set the values and configuration it)
├── charts/             # A directory containing any charts upon which this chart depends.
├── templates/          # A directory of templates that, when combined with values,│                       
    01)pod.yaml
	  02)deployment.yaml
	  03)config.yaml
	  04)service.yaml
	  05)nad.yaml
```
**Create helm**
```
helm create my-chart
```
**Package and Deploy the Chart**
```
helm package my-chart
```
```
helm install test ./my-chart
```
**Verify the Deployment**
```
kubectl get all
```
**Update**
```
 helm upgrade test ./my-chart
```
