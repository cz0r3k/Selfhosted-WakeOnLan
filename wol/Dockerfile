FROM golang:1.19 as builder
WORKDIR /app
COPY wol.go ./
RUN go build wol.go


FROM debian:stable-slim
RUN apt-get -y update && apt-get -y install wakeonlan
COPY --from=builder /app/wol .
CMD ["./wol"]