package processor

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/lumos/datasource"
	"github.com/margostino/lumos/loader"
	"strings"
)

var countryMapping = loader.Load()

func LookupReply(update tgbotapi.Update) string {
	input := normalize(update.Message.Text)

	if input != "sweden" { // strategy logic tbd
		return Greeting()
	}
	return datasource.GetData(countryMapping["sweden"])
}

func normalize(text string) string {
	input := strings.ToLower(text)
	input = strings.TrimSpace(input)
	return input
}
