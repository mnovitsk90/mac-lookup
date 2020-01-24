# base golang with linux to build go module
FROM golang:alpine AS build

RUN mkdir /src

ADD macLookup.go /src

RUN cd /src && \
    go build -o getMac ./macLookup.go

FROM alpine

WORKDIR /bin

COPY --from=build /src/getMac /bin

ENTRYPOINT ["./getMac"]