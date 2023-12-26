# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /src

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download


# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY main.go ./
COPY apis ./apis
COPY cmd ./cmd
COPY config ./config
COPY entity ./entity
COPY repository ./repository
COPY service ./service
COPY usecase ./usecase

RUN go generate ./entity

RUN CGO_ENABLED=0 GOOS=linux go build -o app

# Run
ENTRYPOINT [ "/src/app" ]



