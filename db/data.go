package db

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetData(filename string) map[string]interface{} {
	baseUrl := os.Getenv("APP_DATA_BASE_URL")
	url := fmt.Sprintf("%s/%s.json", baseUrl, filename)

	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	jsonResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	data := make(map[string]interface{})
	err = json.Unmarshal(jsonResponse, &data)

	if err != nil {
		log.Fatalln(err.Error())
	}

	return data
}
