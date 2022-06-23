# Argo workloads with k8s kind

Repo for testing argocd

## Prequiste

Install the followings tool

- Docker
- K8s kind
- Helm v3

## Usage

### Use kind to create our cluster

```bash
kind create cluster --name argocd --config kind/cluster-config.yaml
```

### Install argocd by helm

```bash
helm install -f .argocd/helm-values.yaml argocd argo/argo-cd --namespace argocd --create-namespace
```

### Create bootstrap app and initital resources

- Github secret creds, can be replaced by any scm bitbucket, gitlab

```bash
k apply -f .argocd/git-creds.yaml
```

- Bootstrap app

```bash
k apply -f .argocd/bootstrap-app.yaml
```

- Create nginx ingress controller

```bash
k apply -f .argocd/nginx-ingress-controller.yaml
```

- Create argocd ingress object

```bash
k apply -f .argocd/argocd-ingress.yaml
```

---
