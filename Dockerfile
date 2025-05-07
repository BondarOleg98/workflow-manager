FROM golang:1.24 AS build
WORKDIR /app
COPY app/ main.go go.sum go.mod .
WORKDIR /app
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o workflow-manager

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=build /app/workflow-manager .
EXPOSE 8080
CMD ["./workflow-manager"]
