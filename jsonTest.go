package main

import (
	"fmt"
	"reflect"

	yaml "sigs.k8s.io/yaml"
)

func main(){
	testJson := []byte(`{"interfaces": ["name": "test1", "name": "123", "name": "3.14", "name": "0e7480436552138"]}`)
	testYaml := []byte(`interfaces:
- name: 01c6e4e33a6b25a
- name: 0e7480436552138
`)
	var state map[string]interface{}
	var stateJson map[string]interface{}
	err := yaml.Unmarshal(testJson, &stateJson)
		if err != nil {
		fmt.Println("ERROR: ", err)
	}
	err = yaml.Unmarshal(testYaml, &state)
	if err != nil {
		fmt.Println("ERROR: ", err)
	}

	// JSON
	fmt.Println("JSON")
	interfaces := stateJson["interfaces"]
	for _, iface := range interfaces.([]interface{}) {
		name := iface.(map[string]interface{})["name"]
		fmt.Println(reflect.TypeOf(name.(string)))
		fmt.Println(name.(string))
	}
	// YAML
	fmt.Println("\nYAML")
	interfaces = state["interfaces"]
	for _, iface := range interfaces.([]interface{}) {
		name := iface.(map[string]interface{})["name"]
		fmt.Println(reflect.TypeOf(name.(string)))
		fmt.Println(name.(string))
	}
}
