FROM golang:1.26-alpine AS build
WORKDIR /src
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
RUN go build -o /out/park ./...
FROM alpine:3.21
COPY --from=build /out/park /park
EXPOSE 8080
ENTRYPOINT ["/park"]
