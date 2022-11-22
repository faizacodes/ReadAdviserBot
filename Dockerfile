FROM golang:1.16-alpine

RUN mkdir /app

ADD . /app/

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o tgbot

CMD [ "/app/tgbot" ]