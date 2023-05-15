FROM golang:1.20-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -ldflags "-s -w" -o ./snyp

ENV ADDR :80
EXPOSE 80

CMD [ "./snyp" ]

FROM alpine:latest

COPY --from=builder /app/snyp /app/snyp

ENTRYPOINT [ "/app/snyp" ]
