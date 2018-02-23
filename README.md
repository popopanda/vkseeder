[![CircleCI](https://circleci.com/gh/popopanda/vkseeder/tree/master.svg?style=svg)](https://circleci.com/gh/popopanda/vkseeder/tree/master)

# Reads values from a yml file and imports it to Vault

# Prerequisites
The following Environment Variables need to be set:

- VAULT_TOKEN: Auth token to write to vault
- VAULT_ADDR: Address of the Vault host
- YAML_ENTRY_FILE: Path of the yml file

Your Vault token would require the following policy:

```
path "secret/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
```


#### Sample yml file

```
keys:
    - key: secret/shiba
      values:
        username: shiba inu
        password: muchwow
    - key: secret/spaceship
      values:
        destination: mars
    - key: secret/car
      values:
        color: Red
    - key: secret/moon
      values:
        Distance: 238900 miles
```

This would result in adding the keys in the yaml file as k/v in vault. Thus reading from vault using the keys `secret/shiba` and `secret/moon` would get:

```
vault read secret/shiba

Key                 Value
---                 -----
refresh_interval    768h
password            muchwow
username            shiba inu

vault read secret/moon

Key                 Value
---                 -----
refresh_interval    768h
Distance            238900 miles
```

# Docker run
`docker run -e VAULT_TOKEN=$VAULT_TOKEN -e VAULT_ADDR=$VAULT_ADDR -e YAML_ENTRY_FILE=$YAML_ENTRY_FILE popopanda/vkseeder`

By Default the docker image sets the following environment variables in the container

```
ENV VAULT_TOKEN ""
ENV VAULT_ADDR "http://127.0.0.1:8200"
ENV YAML_ENTRY_FILE "/tmp/example.yml"
```
