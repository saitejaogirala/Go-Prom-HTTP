FROM golang:1.15.0

WORKDIR /app

RUN export GO111MODULE=on

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . .

# Build the application
RUN go build -o main .

# Expose port 9000 to the outside world
EXPOSE 9000


CMD ["./main"]