FROM golang:1.17-alpine as builder

RUN apk update && apk add --no-cache git ca-certificates tzdata 

COPY ./bin/companies /companies
COPY ./cmd/companies/sql /sql

FROM scratch AS final

# Import the time zone files
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# Import the CA certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Import the compiled go executable

COPY --from=builder /companies /companies
COPY --from=builder /sql /sql

WORKDIR /

ENTRYPOINT ["/companies"]

EXPOSE 8080
EXPOSE 8081