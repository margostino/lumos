package processor

import (
	"encoding/json"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/margostino/lumos/common"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	Msg    string `json:"text"`
	ChatID int64  `json:"chat_id"`
	Method string `json:"method"`
}

var bot, _ = newBot()

func Reply(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var update tgbotapi.Update
	if err := json.Unmarshal(body, &update); err != nil {
		log.Fatal("Error updating →", err)
	}

	log.Printf("[%s@%d] %s", update.Message.From.UserName, update.Message.Chat.ID, update.Message.Text)

	if update.Message.Location != nil {
		geoData := fmt.Sprintf("Latitude: %f - Latitude: %f\n", update.Message.Location.Latitude, update.Message.Location.Longitude)
		reply := geoData
		log.Print(geoData)
		data := Response{
			Msg:    reply,
			Method: "sendMessage",
			ChatID: update.Message.Chat.ID,
		}

		message, _ := json.Marshal(data)
		log.Printf("Response %s", string(message))
		fmt.Fprintf(w, string(message))
	} else {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, " ")
		//msg.ReplyToMessageID = update.Message.MessageID
		btn := tgbotapi.KeyboardButton{
			RequestLocation: true,
			Text:            "Send location",
		}
		msg.ReplyMarkup = tgbotapi.NewReplyKeyboard([]tgbotapi.KeyboardButton{btn})
		bot.Send(msg)
	}

	//reply := processor.LookupReply(update)
	//
	//data := Response{
	//	Msg:    reply,
	//	Method: "sendMessage",
	//	ChatID: update.Message.Chat.ID,
	//}
	//
	//message, _ := json.Marshal(data)
	//log.Printf("Response %s", string(message))
	//fmt.Fprintf(w, string(message))

}

func newBot() (*tgbotapi.BotAPI, error) {
	client, error := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	//bot.Debug = true
	common.SilentCheck(error, "when creating a new BotAPI instance")
	//log.Printf("Authorized on account %s\n", client.Self.UserName)
	return client, error
}
