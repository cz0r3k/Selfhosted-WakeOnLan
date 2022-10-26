FROM golang:bullseye as builder
WORKDIR /app
COPY wol.go ./
RUN go build wol.go

FROM debian:bullseye
RUN apt-get -y update && apt-get -y install wakeonlan
COPY --from=builder /app/wol .
ADD ./static/ ./static/
EXPOSE 8080
CMD ["./wol"]
