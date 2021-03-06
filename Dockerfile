# base golang with linux to build go module
FROM golang:alpine AS build

# create directory for build files
RUN mkdir /src

# put the GoLang module in the build directory
ADD macLookup.go /src

# cd to build directory and build the go module
RUN cd /src && \
    go build -o getMac ./macLookup.go

# only need linux environment to execute
# the already-built go module
FROM alpine

# set work directory to execute code from
WORKDIR /bin

# copy the compiled go code to
# work directory '/bin'
COPY --from=build /src/getMac /bin

# set location for docker container to begin executing
ENTRYPOINT ["./getMac"]