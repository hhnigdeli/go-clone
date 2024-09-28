package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func Parse(doc string) {

	// Open our jsonFile
	jsonFile, err := os.Open(doc)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened : %s\n", doc)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	for key, value := range result["dependencies"].(map[string]interface{}) {
		fmt.Println("::: Dependency >>> ", key)
		fmt.Println("::::::::::::::::::::")
		fmt.Printf(" %v\n", value)
		fmt.Println("::::::::::::::::::::")
	}

}
