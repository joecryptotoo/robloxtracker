# syntax=docker/dockerfile:1

FROM golang:1.19-alpine AS build

WORKDIR /src

COPY . ./

RUN go build -o /src/robloxtracker

FROM alpine:latest AS final

COPY --from=build /src/robloxtracker /robloxtracker

ENTRYPOINT [ "/robloxtracker" ]
