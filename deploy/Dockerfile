FROM golang:1.18-alpine as build
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY ./go.mod .
COPY ./go.sum .
ENV GOPROXY https://proxy.golang.org,direct
RUN go mod download
ENV CGO_ENABLED=0
RUN GOOS=linux go build -o main ./cmd/http/main.go

FROM scratch as serve
WORKDIR /app
COPY --from=build /app/config/env.json ./config/env.json
COPY --from=build /app/main .
CMD ["/app/main"]