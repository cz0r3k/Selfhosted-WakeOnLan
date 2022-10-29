FROM golang:alpine3.16 as builder
WORKDIR /app
COPY wol.go ./
RUN go build wol.go

FROM alpine:3.16
COPY --from=builder /app/wol .
ADD ./static/ ./static/
EXPOSE 8080
CMD ["./wol"]
