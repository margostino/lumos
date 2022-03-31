package datasource

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetData(filename string) string {
	var data string
	baseUrl := os.Getenv("APP_DATA_BASE_URL")
	url := fmt.Sprintf("%s/%s", baseUrl, filename)

	response, err := http.Get(url)

	if err != nil {
		log.Fatalln(err)
	}

	jsonResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	dataMap := make(map[string]interface{})
	err = json.Unmarshal(jsonResponse, &dataMap)

	if err != nil {
		log.Fatalln(err.Error())
	}

	for topic, indicators := range dataMap {
		data += fmt.Sprintf("📌   %s\n", topic)
		values := indicators.(map[string]interface{})
		for key, value := range values {
			data += fmt.Sprintf("  ►   %s: %v\n", key, value)
		}
	}

	return data
}
