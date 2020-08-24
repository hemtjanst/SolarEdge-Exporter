FROM golang:1.15-alpine AS build
WORKDIR /source
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o solaredge-exporter -v main.go

FROM scratch
ENV INVERTER_ADDRESS 192.168.1.189
ENV INVERTER_PORT 502
ENV EXPORTER_INTERVAL 5

COPY --from=build /source/solaredge-exporter /solaredge-exporter
ENTRYPOINT ["/solaredge-exporter"]