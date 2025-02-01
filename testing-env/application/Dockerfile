############################################################

# this is the struture for the go application
#
# /app                              # working directory
# ├── cmd
# │   └── application
# │       └── main                  # The main program executable file
# ├── go.mod                        # Go module definition (dependencies and versioning)
# └── src                           # Source code directory
#     └── main.go                   # Main application source file
#
# 3 directories, 3 files


############################################################

FROM golang:alpine


# Set the working directory in the container
WORKDIR /app


# Copy go.mod and go.sum for dependencies
COPY ./go.mod  ./


# Download dependencies
RUN go mod download


# Copy the Go source files into the src directory
COPY main.go /app/src/main.go


# make executable directory cmd server
RUN mkdir -p /app/cmd/application


# Build the Go executable
RUN go build -o /app/cmd/application/main /app/src/main.go


# Set the CMD to run the executable
CMD ["/app/cmd/application/main"]


# Expose the application port
EXPOSE 5000

############################################################
