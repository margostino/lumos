package processor

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Replier func(input string) (bool, string)

var CountryMapping = loadIndex()
var Repliers = LoadRepliers()

func loadIndex() map[string]string {
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

func LoadRepliers() []Replier {
	return []Replier{
		CountryReply,
		IndicatorReply,
		GoalsReply,
		StartReply,
		FallbackReply, // always fallback last in slice
	}
}
