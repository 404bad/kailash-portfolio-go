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

```bash
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






##  Tech Stack

- **Language:** Go (stdlib only — `net/http`)
- **Frontend:** HTML, CSS, vanilla JS
- **Fonts:** Space Mono + Syne (Google Fonts)

##  License

MIT — see [LICENSE](LICENSE)
