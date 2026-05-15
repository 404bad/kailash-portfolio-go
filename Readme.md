# Kailash Badu — Portfolio

A personal portfolio website built with **Go** (`net/http`) — no frameworks, no dependencies. Serves static HTML pages for a fast, minimal footprint.

## 📁 Project Structure

```
kailash-portfolio/
├── static/
│   ├── images/         # Static image assets
│   ├── home.html       # Landing page
│   ├── about.html      # About & experience
│   ├── courses.html    # Certifications & training
│   └── contact.html    # Contact information
├── main.go             # Go HTTP server
├── main_test.go        # Unit tests
├── go.mod              # Go module definition
├── .gitignore
├── LICENSE
└── README.md
```

## 🚀 Running Locally

**Prerequisites:** Go 1.21+

```bash
# Clone the repo
git clone https://github.com/404bad/kailash-portfolio-go.git
cd portfolio

# Run the server
go run main.go

or

go build -o dist

./dist

# Server starts at:
# 🚀 Kailash Badu Portfolio running at http://localhost:8080
```

## Running Tests

```bash
go test ./...
```

## Routes

| Route      | Page                         |
| ---------- | ---------------------------- |
| `/`        | Home — landing & projects    |
| `/about`   | About, experience, education |
| `/courses` | Certifications & training    |
| `/contact` | Contact information          |

## Custom Port

```bash
PORT=3000 go run main.go
```

==================================== DevOps part ==========================================

# DevOps

## Step 1: Containerize the application and it dependency

Docker is a containerization platform that packages an application and its dependencies into a portable unit called a container.

we use docker here because consistency across environments, isolation, build once run anywhere (portability), scalability(docker's lightweight nature makes it easy to spin up many copies of an application to handle high traffic)

```Dockerfile
FROM golang:1.21-alpine  as base
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o dist

FROM gcr.io/distroless/base
COPY --from=base /app/dist .
COPY --from=base /app/static ./static

EXPOSE 8080
CMD ["./dist"]
```

```bash
docker build -t kailashbadu/kailash-portfolio-go:v1 .

docker run -itd -p 8080:8080 kailashbadu/kailash-portfolio-go:v1

check http://localhost:8080

docker login

docker push kailashbadu/kailash-portfolio-go:v1


```

## step 2: kubernetes

Kubernetes is an open source container orchestration platform that automates deployment, scaling, and management of containerized applications.

- self-healing
- auto-scaling
- rolling updates
- cluster environment
- load balancing

when we are writing kubernetes deployment out of the box behavior is the kubernetes will try to pull the images from image registry.OFC we can also use the image form local but it is practice to pull the image form registry
so we need to push it to registry either ecr or dockerhub

```bash
mkdir -p  k8s/manifests

```

in kubernetes service discovery happens using selector and labels

deployment.yml

deployment is a kubernetes resource that manages the number of identical pods (replicaset), it handles rolling updates (zero downtime deployment), roll back if something goes wrong, manages replicaset underneath

```yml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: kailash-portfolio
  namespace: portfolio
  labels:
    app: kailash-portfolio
spec:
  replicas: 2

  selector:
    matchLabels:
      app: kailash-portfolio

  template:
    metadata:
      labels:
        app: kailash-portfolio
    spec:
      containers:
        - name: kailash-portfolio
          image: kailashbadu/kailash-portfolio-go:v1
          imagePullPolicy: Always

          ports:
            - containerPort: 8080
              protocol: TCP

          env:
            - name: PORT
              value: "8080"

          resources:
            requests:
              cpu: "50m"
              memory: "32Mi" # 32MB RAM guaranteed
            limits:
              cpu: "200m" # 0.2 CPU cores max
              memory: "64Mi" # 64MB RAM max

          # Liveness probe — if this fails, Kubernetes restarts the container
          livenessProbe:
            httpGet:
              path: /
              port: 8080
            initialDelaySeconds: 5
            periodSeconds: 15
            failureThreshold: 3

          # Readiness probe — if this fails, pod is removed from the load balancer
          # (no traffic sent until the app is truly ready)
          readinessProbe:
            httpGet:
              path: /
              port: 8080
            initialDelaySeconds: 3
            periodSeconds: 10
            failureThreshold: 3
```

service.yml

kubernetes object/resource that gives pods stable network endpoint (pods come and go ip changes),
it act as a gateway between network and the pods

load balances traffic across all matching pods

acts as DNS entry inside the cluster

there are three types of services in kubernetes: clusterIP, nodePort, LoadBalancer

btw pod is a smallest deployable unit in kubernetes. It wraps one or more containers that share: the same network namespace ( same ip ), the same storage volumes , the same lifecycle

and a cluster is a set of machines (nodes) that run kubernetes

```yml
apiVersion: v1
kind: Service
metadata:
  name: kailash-portfolio-service
  namespace: portfolio
  labels:
    app: kailash-portfolio-service
spec:
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
  selector:
    app: kailash-portfolio
  type: ClusterIP
```

ingress.yml

Ingress is a kubernetes resource that manages external HTTP/HTTPS traffic into the cluster providing: path based routing, host based routing

without ingress every service needs its own loadbalancer = expensive since each LB costs money in cloud

with ingress: one load balancer -> ingress controller -> routes to multiple services

```yml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: kailash-portfolio-ingress

  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  ingressClassName: nginx
  rules:
    - host: kailash-portfolio.local
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: kailash-portfolio-service
                port:
                  number: 80
```

Namespace: A kubernetes namespace is a logical partition within a single physical cluster that creates a "virtual cluster". It is primarily to isolate groups of resources , such as pods and services, when multiple teams or projects share the same

```bash
kubectl create namespace portfolio
kubectl apply -f deployment.yml
kubectl apply -f service.yml
kubectl apply -f ingress.yml

kubectl get ns

minikube addons enable ingress

kubectl get pods -n ingress-nginx
kubectl get deployment -n portfolio
kubectl get svc -n portfolio
kubectl get ing -n portfolio


minikube ip

sudo vim /etc/hosts

192.168.49.2 kailash-portfolio.local

http://kailash-portfolio.local

```

The below command is to stop minikube and free its resources.

```bash
minikube stop # This just pauses the cluster.
minikube delete # Delete Minikube cluster (recommended full reset)

# Full cleanup (Docker + Minikube))
minikube delete --all --purge
docker system prune -af
docker volume prune -f
docker network prune -f


```

## Tech Stack

- **Language:** Go (stdlib only — `net/http`)
- **Frontend:** HTML, CSS, vanilla JS
- **DevOps:** Docker,kubernetes
