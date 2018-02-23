FROM golang:1.10-alpine as gobuilder

WORKDIR /opt
COPY main.go .
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh && \
    go get github.com/hashicorp/vault/api && \
    go get gopkg.in/yaml.v2 && \
    go build -o vkseeder main.go

FROM golang:1.10-alpine

ENV VAULT_TOKEN ""
ENV VAULT_ADDR "http://127.0.0.1:8200"
ENV YAML_ENTRY_FILE "/tmp/example.yml"

WORKDIR /usr/local/bin

COPY --from=gobuilder /opt/vkseeder .
COPY example.yml /tmp
CMD vkseeder