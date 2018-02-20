package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Secrets struct {
	Dog struct {
		Name    string `json:"name"`
		Secrets []struct {
			Username string `json:"username,omitempty"`
			Password string `json:"password,omitempty"`
			Color    string `json:"color,omitempty"`
		} `json:"secrets"`
	} `json:"Dog"`
	Cat struct {
		Name    string `json:"name"`
		Secrets []struct {
			Words string `json:"words,omitempty"`
			Color string `json:"color,omitempty"`
		} `json:"secrets"`
	} `json:"Cat"`
}

func main() {

	filename, err := filepath.Abs("./demo.yml")
	if err != nil {
		log.Fatal(err)
	}
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	var secrets Secrets

	err = yaml.Unmarshal(yamlFile, &secrets)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Value: %#v\n", secrets.Dog.Secrets["Shiba"])

}
