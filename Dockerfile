# syntax=docker/dockerfile:1

FROM golang:1.24-bookworm

RUN mkdir -p /app/ui/dist

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

COPY auth /app/auth
COPY cors /app/cors
COPY hotswap /app/hotswap
COPY mac /app/mac
COPY models /app/models
COPY ui/dist /app/ui/dist

WORKDIR /app/mac
RUN git init
RUN git config user.name "AnyName" && git config user.email "any@email.com"
RUN git commit --allow-empty -n -m "Initial commit."

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 4000

# Run
CMD ["/app/mac/run.sh"]