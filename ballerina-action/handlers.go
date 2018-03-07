package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
)

func InitHandler(writer http.ResponseWriter, request *http.Request) {
	// Decode request message body
	decoder := json.NewDecoder(request.Body)
	var init Init
	err := decoder.Decode(&init)
	if err != nil {
		panic(err)
	}
	defer request.Body.Close()

	// Write function to a ballerina file
	content := []byte(init.Value.Code)
	fileName := "function.bal"
	if init.Value.Binary {
		fileName = "function.balx"
	}
	
	err = ioutil.WriteFile(fileName, content, 0644)
	if err != nil {
		panic(err)
	}

	if !init.Value.Binary {
		// Compile ballerina function
		_, err = exec.Command("sh", "-c", "ballerina build function.bal").Output()
		if err != nil {
			panic(err)
		}
	}

	writer.WriteHeader(http.StatusOK)
}

func RunHandler(writer http.ResponseWriter, request *http.Request) {
	// Execute ballerina function
	out, err := exec.Command("sh", "-c", "ballerina run function.balx").Output()
	if err != nil {
		panic(err)
	}
	log.Printf("%s", out)
	writer.Write(out)
}
