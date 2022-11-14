# Golang base image
FROM golang:alpine as builder

# Adds maintainer information
LABEL maintainer="Gill Erick <ogayogill95@gmail.com>"

# Installs git as it is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Sets the current working directory inside the container
WORKDIR /app

# Copies go mod and sum files
COPY go.mod go.sum ./

# Downloads all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

# Copies the source from the current directory to the working directory inside the container
COPY . .

# Builds the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Starts a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copies the pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Expose port 8080 to the outside world
EXPOSE 8080

#Command to run the executable
CMD ["./main"]
