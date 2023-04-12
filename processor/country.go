package processor

import (
	"fmt"
	"github.com/margostino/lumos/db"
)

func CountryReply(input string) (bool, string) {
	if input == "sweden" {
		data := db.GetData("sweden")
		return true, prepareCountryReply(data)
	}
	return false, ""
}

func prepareCountryReply(data map[string]interface{}) string {
	var reply string
	population := data["population"]
	reply += fmt.Sprintf("🇸🇪 Population:   %.0f\n", population)
	for topic, indicators := range data {
		if topic != "population" {
			reply += fmt.Sprintf("📌   %s\n", topic)
			values := indicators.(map[string]interface{})
			for key, value := range values {
				reply += fmt.Sprintf("  ►   %s: %v\n", key, value)
			}
		}
	}
	return reply
}
