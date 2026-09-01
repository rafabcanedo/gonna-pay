FROM golang:1.24-alpine

WORKDIR /app

# Copy the files that list your libs
COPY go.mod go.sum ./
RUN go mod download

# Copy all code for inside the container
COPY . .

# Container use this port:3333
EXPOSE 3333

# Comand for ru the server
CMD ["go", "run", "cmd/api/main.go"]