# Step 7 App on Kubernetes

In this step, Kubernetes is used to manage all containers.

2 cotainers are used in this step.

- web (Golang): This is front-end application. This application gets data from wordservice using gRPC client and make html data.
- wordservice (Golang): This is back-end application using gRPC.

## Environment

- Kubernetes : single node cluster (v1.18.12-gke.1210) on GKE
- Golang : 1.13.12

## Setup

Run this command on your client PC.

### Setup GKE

```bash
# Create Kuberentes cluster on GCP
gcloud container clusters create word-database --machine-type=n1-standard-1 --num-nodes=1
gcloud container clusters get-credentials word-database
```

### Create container images

```bash
# Build and push containers to Container Registry on GCP
git clone https://github.com/sleek629/application-modernization-example.git
cd application-modernization-example/step7/src/internal/v1/web
docker build -t gcr.io/$GCP_PROJECT/web:1.0 .
docker push gcr.io/$GCP_PROJECT/web:1.0
cd ../wordservice
docker build -t gcr.io/$GCP_PROJECT/wordservice:1.0 .
docker push gcr.io/$GCP_PROJECT/wordservice:1.0
cd ../../v2/wordservice
docker build -t gcr.io/$GCP_PROJECT/wordservice:2.0 .
docker push gcr.io/$GCP_PROJECT/wordservice:2.0
```

## Examle 1

Deploy only one application.

### Create Kubernetes resources

```bash
cd application-modernization-example/step7/src/kubernetes-manifests/ex1
# Create namespaces (DNS namespace)
kubectl apply -f namespace.yaml
kubectl config set-context --current --namespace=word-database
# Create deployment (Application)
kubectl apply -f deployment.yaml
# Create service (Load balancer)
kubectl apply -f service.yaml
# Get IP to the service
kubectl get service
```

Access to http://EXTERNAL-IP

## Examle 2

Deploy three application using ReplicaSet with load balancer.

```bash
cd application-modernization-example/step7/src/kubernetes-manifests/ex2
kubectl apply -f deployment.yaml
```

Access to http://EXTERNAL-IP

You can see different Hostname and data when you reload the page. (if not, delete the cache)

If you delete one pod, you can see a new pod is created.

```bash
# Get pods list
kubectl get pods
# Delete one pod
kubectl delete pods {{name of one pod from the list above}}
# Get new pods list
kubectl get pods
```

## Examle 3

Deploy new version for gRPC API (wordservice v2) which enables sort by the word count.

```bash
cd application-modernization-example/step7/src/kubernetes-manifests/ex3
kubectl apply -f deployment.yaml
```

Access to http://EXTERNAL-IP

Now, you can see the sorted word counts.
