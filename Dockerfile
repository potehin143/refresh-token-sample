FROM golang:latest as builder

ENV GO111MODULE=auto
RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get go.mongodb.org/mongo-driver/mongo
RUN go get github.com/google/uuid
RUN go get golang.org/x/crypto/bcrypt

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main main.go

#Second stage

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/* .
CMD ["/app/main"]

EXPOSE 8080