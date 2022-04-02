package processor

import (
	"fmt"
	"github.com/margostino/lumos/datasource"
	"github.com/margostino/lumos/utils"
)

var sdgCommands = []string{
	"goals",
	"poverty",
	"hunger",
	"health",
	"education",
	"equality",
	"water",
	"energy",
	"growth",
	"innovation",
	"inequality",
	"community",
	"resources",
	"climate",
	"oceans",
	"land",
	"institutions",
	"partnership",
}

func GoalsReply(input string) (bool, string) {
	if utils.Contains(sdgCommands, input) {
		data := datasource.GetData("sdg")
		return true, prepareGoalsReply(input, data)
	}
	return false, ""
}

func prepareGoalsReply(input string, data map[string]interface{}) string {
	if input == "goals" {
		return fmt.Sprintf("🎯   %s\n", data["description"].(string))
	}

	goals := data["goals"].(map[string]interface{})
	reply := fmt.Sprintf("🌱   %s\n", goals[input])
	return reply
}
