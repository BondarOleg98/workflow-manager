FROM docker.io/library/golang:1.24
ENV GO111MODULE=on
WORKDIR /app
COPY app ./app
COPY *.go go.* ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o workflow-manager
EXPOSE 8080
CMD ["./workflow-manager"]
