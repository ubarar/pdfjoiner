FROM golang:latest AS builder

COPY . /app

RUN cd /app && go mod tidy && cd /app/cmd && go build .

# multi-stage build
FROM ubuntu:latest

RUN mkdir /app
RUN apt update && apt install -y ghostscript

COPY --from=builder /app/cmd/cmd /app
COPY --from=builder /app/cmd/index.html /app
COPY --from=builder /app/cmd/cert.pem /app
COPY --from=builder /app/cmd/key.pem /app

EXPOSE 8080

WORKDIR /app

CMD /app/cmd