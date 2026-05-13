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

##  Running Tests

```bash
go test ./...
```

##  Routes

| Route        | Page                        |
|--------------|-----------------------------|
| `/`          | Home — landing & projects   |
| `/about`     | About, experience, education|
| `/courses`   | Certifications & training   |
| `/contact`   | Contact information         |

## Custom Port

```bash
PORT=3000 go run main.go
```

==================================== DevOps part ==========================================
# DevOps

## Step 1: Containarize the application and it dependendcies

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

## step 2: kubernets

when we are writing kubernetes deployment out of the box behaviour is the kubernets will try to pull the images from image registry.OFC we can also use  the image form local but it is pratice to pull the image form registry
so we need to push it to registry either ecr or dockerhub

```bash
mkdir -p  k8s/manifests

```
 in kubernetes service discovery happend using selector and labels

deployment.yml

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




service.yml

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


ingress.yml

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


##  Tech Stack

- **Language:** Go (stdlib only — `net/http`)
- **Frontend:** HTML, CSS, vanilla JS
- **Fonts:** Space Mono + Syne (Google Fonts)

##  License

MIT — see [LICENSE](LICENSE)
