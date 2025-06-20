FROM golang:1.24-alpine
WORKDIR /project
COPY go.mod go.sum ./
RUN go mod download 
COPY . .
RUN go build -o txnsvc main.go
EXPOSE 8080
CMD ["./txnsvc"]
