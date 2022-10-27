FROM golang:bullseye as builder
WORKDIR /app
COPY wol.go ./
RUN go build wol.go

FROM debian:bullseye
COPY --from=builder /app/wol .
ADD ./static/ ./static/
EXPOSE 8080
EXPOSE 8089
CMD ["./wol"]
