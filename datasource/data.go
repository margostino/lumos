package datasource

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func GetData(filename string) string {
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
	return string(jsonResponse)
}
