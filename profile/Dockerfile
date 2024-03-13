# Stage 1: Build the application and install Gin
FROM golang:latest AS builder


# Set the current work directory inside the container
WORKDIR /app

# Install air reload
# RUN go install github.com/cosmtrek/air@latest

# COPY the mod and sum files to the current dir
COPY go.mod go.sum ./

# Download all the dependencies
RUN go mod download



# Install gin
# RUN go install github.com/codegangsta/gin@latest


# COPY the rest of the App code inside your dir

COPY . .


# BUILD the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o profile .

# Stage 2: Create the production container
FROM golang:latest AS production

# Set the current working directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/profile .
# COPY --from=builder /go/bin/gin /usr/local/bin/gin

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./profile"]
