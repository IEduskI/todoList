package kit

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func RetrieveOptions() (status, priority map[interface{}]interface{}) {
	//Leer el archivo
	data, err := ioutil.ReadFile("options.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

	//Se crea el struct
	var dataMap map[string]interface{}

	//Unmarshall the yaml
	err = yaml.Unmarshal(data, &dataMap)
	if err != nil {
		fmt.Println(err)
		return
	}

	//Imprimir la informacion
	status = dataMap["status"].(map[interface{}]interface{})
	priority = dataMap["priority"].(map[interface{}]interface{})
	//id := 1

	return
}
