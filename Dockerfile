# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

WORKDIR /src

COPY . ./

RUN go build -o /robloxtracker

CMD [ "/robloxtracker" ]
