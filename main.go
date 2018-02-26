package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/hashicorp/vault/api"
	yaml "gopkg.in/yaml.v2"
)

// Secrets to match YAML file
type Secrets struct {
	Keys []struct {
		Key    string                 `json:"key"`
		Values map[string]interface{} `json:"values"`
	} `json:"keys"`
}

func main() {

	AuthToken := os.Getenv("VAULT_TOKEN")
	VaultADDR := os.Getenv("VAULT_ADDR")
	YamlEntryFile := os.Getenv("YAML_ENTRY_FILE")

	client := (vaultClient(AuthToken, VaultADDR))
	clientLogical := client.Logical()
	clientSys := client.Sys()
	resultKeys := parseYAML(YamlEntryFile)

	for _, data := range resultKeys.Keys {
		clientSys.Mount(data.Key, &api.MountInput{
			Type:        "kv",
			Description: fmt.Sprintf("%v mount path", data.Key),
		})
		_, err := clientLogical.Write(data.Key, data.Values)
		if err != nil {
			fmt.Printf("%v failed to be written\n", data.Key)
			log.Fatal(err)
		}
		fmt.Printf("%v was written to vault successfully\n", data.Key)
	}
}

func vaultClient(vaultToken, host string) *api.Client {
	client, err := api.NewClient(&api.Config{
		Address: host,
	})

	if err != nil {
		log.Fatal(err)
	}

	client.SetToken(vaultToken)
	return client
}

func parseYAML(filename string) Secrets {
	var secret Secrets
	reader, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	buf, _ := ioutil.ReadAll(reader)
	yaml.Unmarshal(buf, &secret)

	return secret
}
