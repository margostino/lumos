package processor

import (
	"fmt"
	"github.com/margostino/lumos/datasource"
)

func (r CountryReplier) Apply(input string) bool {
	return input == "sweden"
}

func (r CountryReplier) Reply() string {
	data := datasource.GetData(r.Id)
	return prepare(data)
}

func prepare(data map[string]interface{}) string {
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
