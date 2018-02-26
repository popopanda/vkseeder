[![CircleCI](https://circleci.com/gh/popopanda/vkseeder/tree/master.svg?style=svg)](https://circleci.com/gh/popopanda/vkseeder/tree/master)

[![Go Report Card](https://goreportcard.com/badge/github.com/popopanda/vkseeder)](https://goreportcard.com/report/github.com/popopanda/vkseeder)

# Reads values from a yml file and imports it to Vault

# Prerequisites
The following Environment Variables need to be set:

- VAULT_TOKEN: Auth token to write to vault
- VAULT_ADDR: Address of the Vault host
- YAML_ENTRY_FILE: Path of the yml file

Your Vault token would require specified policies, for example:

```
path "secret/*" {
  capabilities = ["create", "read", "update", "delete", "list"]
}
```


#### Sample yml file

```
keys:
    - key: secret/foo
      values:
        value: bar
    - key: foo/secret
      values:
        value: bar
    - key: baz
      values:
        value: foobar
```

This would result in adding the keys in the yaml file as k/v in vault. Thus reading from vault using the keys `secret/shiba` and `secret/moon` would get:

```
vault read secret/foo

Key                 Value
---                 -----
refresh_interval    768h
value               bar

vault read foo/secret

Key                 Value
---                 -----
refresh_interval    768h
value               bar

vault read baz

Key                 Value
---                 -----
refresh_interval    768h
value               foobar
```

# Docker run
`docker run -e VAULT_TOKEN=$VAULT_TOKEN -e VAULT_ADDR=$VAULT_ADDR -e YAML_ENTRY_FILE=$YAML_ENTRY_FILE popopanda/vkseeder`

By Default the docker image sets the following environment variables in the container

```
ENV VAULT_TOKEN ""
ENV VAULT_ADDR "http://127.0.0.1:8200"
ENV YAML_ENTRY_FILE "/tmp/example.yml"
```

#Limitations
- Entries removed from the yaml file are not deleted in vault