FROM golang:alpine3.16 as builder
WORKDIR /app
COPY wol.go ./
RUN go build wol.go

FROM alpine:3.16
ARG WEB_PORT
COPY --from=builder /app/wol .
ADD ./static/ ./static/
EXPOSE ${WEB_PORT}
CMD ["./wol"]
