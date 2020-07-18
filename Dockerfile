FROM golang:latest as builder
WORKDIR /app
COPY . .

ENV MONGO_URL=mongodb://mongo/rex
ENV URL_DB=mongo
ENV MONGO_TIMEOUT=30
ENV MONGO_DB=rex

RUN CGO_ENABLED=0 GOOS=linux go build -v -o inventory-service

FROM alpine
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/inventory-service /inventory-service

CMD ["/inventory-service"]