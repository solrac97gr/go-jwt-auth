FROM golang:1.18-alpine as build
RUN mkdir /app
ADD .. /app/
WORKDIR /app
COPY ../go.mod .
COPY ../go.sum .
RUN go mod download
RUN go build -o main ./cmd/http/main.go

FROM busybox
COPY --from=build /app/main /app/
EXPOSE 5000
CMD ["./main"]