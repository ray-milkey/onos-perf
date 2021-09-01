FROM onosproject/golang-build:v0.6.3 as build

ENV GO111MODULE=on
COPY . /go/src/github.com/onosproject/onos-perf
RUN cd /go/src/github.com/onosproject/onos-perf && GOFLAGS=-mod=vendor make build

FROM alpine:3.11
RUN apk add libc6-compat

USER nobody

COPY --from=build /go/src/github.com/onosproject/onos-perf/build/_output/onos-perf /usr/local/bin/onos-perf

ENTRYPOINT ["onos-perf"]
