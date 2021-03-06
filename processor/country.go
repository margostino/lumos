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
	population := data["population"]
	reply += fmt.Sprintf("πΈπͺ Population:   %.0f\n", population)
	for topic, indicators := range data {
		if topic != "population" {
			reply += fmt.Sprintf("π   %s\n", topic)
			values := indicators.(map[string]interface{})
			for key, value := range values {
				reply += fmt.Sprintf("  βΊ   %s: %v\n", key, value)
			}
		}
	}
	return reply
}
