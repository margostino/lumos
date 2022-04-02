package processor

import (
	"fmt"
	"github.com/margostino/lumos/datasource"
)

func IndicatorReply(input string) (bool, string) {
	if input == "indicators" {
		data := datasource.GetData("indicators")
		return true, prepareIndicatorReply(data)
	}
	return false, ""
}

func prepareIndicatorReply(data map[string]interface{}) string {
	var reply string
	for indicator, value := range data {
		metadata := value.(map[string]interface{})
		reply += fmt.Sprintf("📌 Indicator:  %s\n", indicator)
		reply += fmt.Sprintf("📚 Source:  %s\n", metadata["source"])
		reply += fmt.Sprintf("📝 Description  %s\n", metadata["description"])
	}
	return reply
}
