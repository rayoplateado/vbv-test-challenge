FROM golang:1.22.3

WORKDIR /app

ADD . /app
RUN go mod download

ENV APP_ENV=docker

RUN CGO_ENABLED=0 GOOS=linux go build -o /factorialfrontend

EXPOSE 8080

CMD [ "/factorialfrontend" ]