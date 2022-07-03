FROM golang:1.18.3-stretch as build
WORKDIR /build
COPY . /build
RUN make

FROM golang:1.18.3-stretch as run
COPY --from=build /build/app /app
ENTRYPOINT ["/app"]