package processor

import (
	"fmt"
	"github.com/margostino/lumos/datasource"
)

func CountryReply(input string) (bool, string) {
	if input == "sweden" {
		data := datasource.GetData("sweden")
		return true, prepareCountryReply(data)
	}
	return false, ""
}

func prepareCountryReply(data map[string]interface{}) string {
	var reply string
	for topic, indicators := range data {
		reply += fmt.Sprintf("📌   %s\n", topic)
		values := indicators.(map[string]interface{})
		for key, value := range values {
			reply += fmt.Sprintf("  ►   %s: %v\n", key, value)
		}
	}
	return reply
}
