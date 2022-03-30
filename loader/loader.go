package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func Load() map[string]string {
	var countryMapping map[string]string
	baseUrl := os.Getenv("APP_DATA_BASE_URL")
	filename := "country-mapping.json"
	url := fmt.Sprintf("%s/%s", baseUrl, filename)

	response, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	jsonResponse, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(jsonResponse, &countryMapping)

	return countryMapping
}
