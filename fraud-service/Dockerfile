FROM golang:1.24-alpine
WORKDIR /project
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o fraudsvc main.go
EXPOSE 8081
CMD ["./fraudsvc"]
