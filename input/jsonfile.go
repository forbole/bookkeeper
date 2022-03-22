package input

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ImportJsonInput(path string)(*Data,error){
	jsonFile, err := os.Open(path)
	if err!=nil{
		return nil,err
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err!=nil{
		return nil,err
	}

	var input Data
	json.Unmarshal(byteValue,&input)

	return &input,nil

}