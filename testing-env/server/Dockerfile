############################################################

# this is the struture for the go server
#
# /app                              # working directory
# ├── cmd
# │   └── server
# │       └── main                  # The main program executable file
# ├── go.mod                        # Go module definition (dependencies and versioning)
# ├── go.sum
# └── src                           # Source code directory
#     └── main.go                   # Main application source file
#
# 3 directories, 4 files


############################################################

FROM golang:alpine


# Set the working directory in the container
WORKDIR /app


# Copy go.mod and go.sum for dependencies
COPY ./go.mod ./go.sum ./


# Download dependencies
RUN go mod download


# Copy the Go source files into the src directory
COPY main.go /app/src/main.go


# make executable directory cmd server
RUN mkdir -p /app/cmd/server


# Build the Go executable
RUN go build -o /app/cmd/server/main /app/src/main.go


# Set the CMD to run the executable
CMD ["/app/cmd/server/main"]


# Expose the application port
EXPOSE 8090

############################################################
