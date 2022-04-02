package processor

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"strings"
)

func LookupReply(update tgbotapi.Update) string {
	input := normalize(update.Message.Text)
	for _, replier := range Repliers {
		if match, reply := replier(input); match {
			return reply
		}
	}
	return "invalid command"
}

func normalize(text string) string {
	input := strings.ToLower(text)
	input = strings.TrimSpace(input)
	return input
}
