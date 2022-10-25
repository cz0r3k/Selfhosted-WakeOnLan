FROM golang:1.19
RUN apt-get -y update && apt-get -y install wakeonlan
COPY wol.go .
RUN go build wol.go
CMD ["./wol"]
