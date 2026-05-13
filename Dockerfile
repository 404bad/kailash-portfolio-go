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


