###############################
# Dockerfile for netutils image
###############################

FROM alpine:latest

# Update package index and install necessary tools
RUN apk update && \
    apk add --no-cache \
    jq \
    grep \
    coreutils \
    bat \
    neovim \
    tcpdump \
    curl \
    bash \
    iputils \
    bind-tools

# Keep the container running
CMD ["tail", "-f", "/dev/null"]
