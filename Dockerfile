# syntax=docker/dockerfile:1

FROM golang:1.19.3-bullseye

WORKDIR /app

ADD /app ./app
ADD /constants ./constants
ADD /models ./models
ADD /queries ./queries
ADD /services ./services
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-heroes

EXPOSE 8080

CMD [ "/docker-heroes" ]