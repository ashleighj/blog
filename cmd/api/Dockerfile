ARG service=blog_api

# ---- 1. BUILD ----
FROM golang:alpine AS builder
ARG service

LABEL stage=builder

# Basic libraries
RUN apk add --no-cache git alpine-sdk

WORKDIR /${project}

# Copy project list of dependencies
COPY ./go.mod .
COPY ./go.sum .

# Install pacakge dependencies
RUN go get -d -v ./...

# Copy project files
COPY . .

# Build application
RUN go build -o ./${project}

# ---- 2. FINALIZE ----
FROM alpine:3.12
ARG project
ENV entry=${project}

RUN apk update
RUN apk --no-cache add ca-certificates \
    curl

# Create non-root user for app
RUN adduser -D -g 'gouser' gouser && \
    mkdir -p /${project} && \
    chown -R gouser:gouser /${project}

USER gouser

# Copy files from builder
# Uncomment this line for builds on your local machine
COPY --from=builder /${project}/settings/settings.yml ${project}/config/config.yml
COPY --from=builder /${project}/${project} ${project}/
COPY --from=builder --chown=gouser:gouser /${project}/files/ /${project}/files/

EXPOSE 8080
WORKDIR /${project}
ENTRYPOINT ./${entry}