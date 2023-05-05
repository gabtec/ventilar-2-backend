# Build step
FROM golang:1.20 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server

# Final step
FROM scratch

COPY --from=builder /app/server /server

EXPOSE 4000

# TODO
# create user
# USER nonroot:nonroot

# CMD [ "/server", "--env", "production" ]
ENTRYPOINT [ "/server", "--env", "production" ]
# ENTRYPOINT [ "/server" ]